package api

import (
	"net/http"
	"time"

	"github.com/yulpa/yulmails/api/entity"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Configuration is the global configuration of the API server
type Configuration struct {
	// Database configuration
	Database ConfDB `json:"database"`
	// Server configuration
	Server ConfSrv `json:"server"`
}

// ConfDB is the database configuration
type ConfDB struct {
	// Username of the database
	Username string `json:"username"`
	// Password of the database
	Password string `json:"password"`
	// Host of the database
	Host string `json:"host"`
	// Port of the database
	Port int `json:"port"`
	// Name of the database to use
	Name string `json:"name"`
}

type ConfSrv struct {
	// Port to listen on
	Port int `json:"port"`
}

func StartAPI(apiConfig string) error {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	// RESTy routes for "articles" resource
	r.Route("/entities", func(r chi.Router) {
		r.Mount("/", entity.NewRouter())
	})

	if err := http.ListenAndServe(":12800", r); err != nil {
		return err
	}
	return nil
}
