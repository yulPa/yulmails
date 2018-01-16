package api

import (
	"github.com/gorilla/mux"

	"net/http"

	"github.com/yulPa/yulmails/mongo"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(routes []Route) *mux.Router {
	/*
	   Create a custom router
	   parameter: <[]Route> An array of route to register
	   return: <router>A mux router
	*/
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

func GetRouterV1(session mongo.Session) *mux.Router {
	/*
		Retourn V1 API
	*/
	var routes = Routes{
		Route{
			Name:        "Create an entity",
			Method:      http.MethodPost,
			Pattern:     "/api/v1/entity",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) { CreateEntity(session, w, r) },
		},
		Route{
			Name:        "Read Entities",
			Method:      http.MethodGet,
			Pattern:     "/api/v1/entities",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) { ReadEntities(session, w, r) },
		},
		Route{
			Name:        "Create a environment for entity",
			Method:      http.MethodPost,
			Pattern:     "/api/v1/entity/{entity}/environment",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) { CreateEnvironment(session, w, r) },
		},
		Route{
			Name:        "Read one entity",
			Method:      http.MethodGet,
			Pattern:     "/api/v1/entity/{entity}",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) { ReadEntity(session, w, r) },
		},
		Route{
			Name:        "Read one environment",
			Method:      http.MethodGet,
			Pattern:     "/api/v1/entity/{entity}/environment/{environment}",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) { ReadEnvironment(session, w, r) },
		},
		Route{
			Name:        "Delete an entity",
			Method:      http.MethodDelete,
			Pattern:     "/api/v1/entity/{entity}",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) { DeleteEntity(session, w, r) },
		},
		Route{
			Name:        "Update an entity",
			Method:      http.MethodPost,
			Pattern:     "/api/v1/entity/{entity}",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) { UpdateEntity(session, w, r) },
		},
		Route{
			Name:        "Delete an environment",
			Method:      http.MethodDelete,
			Pattern:     "/api/v1/entity/{entity}/environment/{environment}",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) { DeleteEnvironment(session, w, r) },
		},
		Route{
			Name:        "Update an environment",
			Method:      http.MethodPost,
			Pattern:     "/api/v1/entity/{entity}/environment/{environment}",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) { UpdateEnvironment(session, w, r) },
		},
		Route{
			Name:        "Read environments associated to an entity",
			Method:      http.MethodGet,
			Pattern:     "/api/v1/entity/{entity}/environment",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) { ReadEnvironments(session, w, r) },
		},
		Route{
			Name:        "Read mails associated to an environment",
			Method:      http.MethodGet,
			Pattern:     "/api/v1/entity/{entity}/environment/{environment}/mails",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) { ReadMails(session, w, r) },
		},
	}

	var router = NewRouter(routes)

	return router
}
