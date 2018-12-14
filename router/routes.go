package router

import (
	"net/http"

	"github.com/srgupta5328/weather/api"
)

//Route Struct to hold Route information
type Route struct {
	Name        string
	Method      string
	Pattern     string
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
			api.HealthCheck,
		},
		Route{
			"Callback",
			"GET",
			"/callback",
			api.Callback,
		},
		Route{
			"Get Forecast",
			"GET",
			"/forecast",
			api.GetForecastHandler,
		},
	}

	return routes
}
