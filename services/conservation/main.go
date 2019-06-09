package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	gw "gitlab.com/tortuemat/yulmails/services/conservation/v1beta1"
)

func run(db *sql.DB) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterConservationServiceHandlerFromEndpoint(ctx, mux, "127.0.0.1:9090", opts)
	if err != nil {
		return err
	}

	pattern, err := runtime.NewPattern(1, []int{int(utilities.OpLitPush), 0}, []string{"healthz"}, "")
	if err != nil {
		log.Printf("unable to create pattern: %s", err)
	}
	mux.Handle(
		"GET",
		pattern,
		func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
			ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
			defer cancel()
			if err := db.PingContext(ctx); err != nil {
				http.Error(w, fmt.Sprintf("unable to connect to db: %v", err), http.StatusFailedDependency)
				return
			}
			w.WriteHeader(http.StatusOK)
			return
		},
	)
	log.Println("HTTP server started on localhost:8080")
	return http.ListenAndServe(":8080", mux)
}

func main() {
	databaseUser := os.Getenv("POSTGRES_USER")
	databasePassword := os.Getenv("POSTGRES_PASSWORD")
	databaseHost := os.Getenv("POSTGRES_HOST")
	databaseDB := os.Getenv("POSTGRES_DB")
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable",
		databaseUser,
		databaseDB,
		databasePassword,
		databaseHost,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("unable to access database: %v", err)
	}
	defer db.Close()
	log.Printf("connected to database: %s on %s", databaseDB, databaseHost)
	dao := &Dao{db: db}

	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	gw.RegisterConservationServiceServer(server, &ConservationService{DaoService: dao})

	go server.Serve(listen)
	log.Println("gRPC server started on localhost:9090")
	if err := run(db); err != nil {
		log.Fatal(err)
	}
}
