package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// SqlRequest is the request to send
// to the DB through Service
type SqlRequest struct {
	Request  string `json:"request"`
	Commit   bool   `json:"commit"`
	Fetchone bool   `json:"fetchone"`
}

// Service is a db wrapper
type Service struct {
	db *sql.DB
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	db := s.db
	switch r.URL.Path {
	default:
		http.Error(w, "not found", http.StatusNotFound)
		return
	case "/healthz":
		ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
		defer cancel()

		if err := db.PingContext(ctx); err != nil {
			http.Error(w, fmt.Sprintf("unable to connect to db: %v", err), http.StatusFailedDependency)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	case "/":
		if r.Method == http.MethodPost {
			var req SqlRequest
			var data []byte
			data, _ = ioutil.ReadAll(r.Body)
			if data != nil {
				var name string
				if err := json.Unmarshal(data, &req); err != nil {
					log.Printf("unable to unmarshal request: %v", err)
				}
				log.Printf("requesting: %s", req.Request)
				row := db.QueryRow(req.Request)
				if err := row.Scan(&name); err != nil {
					log.Printf("unable to scan row: %v", err)
				}
				log.Println(name)
				w.WriteHeader(http.StatusOK)
				return
			}
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
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
	service := &Service{db: db}
	if err != nil {
		log.Fatalf("unable to access database: %v", err)
	}
	defer db.Close()
	log.Printf("connected to %s on %s", databaseDB, databaseHost)
	log.Fatal(http.ListenAndServe(":8080", service))
}
