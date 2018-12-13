package router

import "github.com/gorilla/mux"

//NewRouter initializes a new router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	routes := InitRoutes()
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
