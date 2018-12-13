package router

import "net/http"

//Route Struct to hold Route information
type Route struct {
	Name        string
	Pattern     string
	Method      string
	HandlerFunc http.HandlerFunc
}

//Routes is a list of Routes
type Routes []Route

//InitRoutes ... Initializing Routes to get the project started
func InitRoutes() Routes {
	var routes = Routes{
		Route{
			"Health",
			"GET",
			"/health",
			HealthCheck,
		},
		Route{
			"Callback",
			"GET",
			"/callback",
			Callback,
		},
	}

	return routes
}
