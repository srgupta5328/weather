package api

import "github.com/gorilla/mux"

//NewRouter initializes a new router
func NewRouter() *mux.Router {
	routes := InitRoutes()
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
