package api

import (
	"github.com/gorilla/mux"

	"net/http"
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

func GetRouterV1() *mux.Router {
	/*
		Retourn V1 API
	*/
	var routes = Routes{
		Route{
			Name:        "Create an entity",
			Method:      http.MethodPost,
			Pattern:     "/api/v1/entity",
			HandlerFunc: CreateEntity,
		},
		Route{
			Name:        "Get entitys",
			Method:      http.MethodGet,
			Pattern:     "/api/v1/entity",
			HandlerFunc: GetEntity,
		},
	}

	var router = NewRouter(routes)

	return router
}
