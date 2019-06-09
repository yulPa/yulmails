package main

import (
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	gw "gitlab.com/tortuemat/yulmails/services/conservation/v1beta1"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterConservationServiceHandlerFromEndpoint(ctx, mux, "127.0.0.1:9090", opts)
	if err != nil {
		return err
	}
	log.Println("HTTP server started on localhost:8080")
	return http.ListenAndServe(":8080", mux)
}

func main() {
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	gw.RegisterConservationServiceServer(server, new(ConservationService))

	go server.Serve(listen)
	log.Println("gRPC server started on localhost:9090")
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
