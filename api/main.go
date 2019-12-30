package api

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/yulpa/yulmails/api/abuse"
	"github.com/yulpa/yulmails/api/entity"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Configuration is the global configuration of the API server
type Configuration struct {
	// Database configuration
	Database ConfDB `yaml:"database"`
	// Server configuration
	Server ConfSrv `yaml:"server"`
}

// ConfDB is the database configuration
type ConfDB struct {
	// Username of the database
	Username string `yaml:"username"`
	// Password of the database
	Password string `yaml:"password"`
	// Host of the database
	Host string `yaml:"host"`
	// Port of the database
	Port int `yaml:"port"`
	// Name of the database to use
	Name string `yaml:"name"`
}

type ConfSrv struct {
	// Port to listen on
	Port int `yaml:"port"`
}

// @title YulmailsAPI
// @version 0.1.0
// @description Manage Yulmails resources from this API
// @termsOfService https://yulpa.io

// @contact.name Mathieu Tortuyaux
// @contact.email mathieu.tortuyaux@yulpa.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host api-dev.yulmails.io
// @BasePath /
// @tag.name entity
// @tag.name abuse
func StartAPI(apiConfig string) error {
	content, err := ioutil.ReadFile(apiConfig)
	if err != nil {
		return errors.Wrap(err, "unable to open api config file")
	}
	var c Configuration
	if err := yaml.Unmarshal(content, &c); err != nil {
		return errors.Wrap(err, "unable to load configuration")
	}
	conn := fmt.Sprintf(
		"user=%s dbname=%s sslmode=disable password=%s host=%s port=%d",
		c.Database.Username, c.Database.Name, c.Database.Password, c.Database.Host, c.Database.Port,
	)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return errors.Wrap(err, "unable to create db connection")
	}
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
		// health check on DB connection
		if err := db.Ping(); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	r.Route("/entities", func(r chi.Router) {
		r.Mount("/", entity.NewRouter())
	})
	r.Route("/abuses", func(r chi.Router) {
		r.Mount("/", abuse.NewRouter(db))
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", c.Server.Port), r); err != nil {
		return err
	}
	return nil
}
