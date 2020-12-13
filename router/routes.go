package router

import (
	"net/http"

	"github.com/wyllisMonteiro/go-api-template/controllers"
)

// Route stores data for mux for an API route
type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes stores all API Routes
type Routes []Route

var routes = Routes{
	Route{
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: controllers.GetSample,
	},
}
