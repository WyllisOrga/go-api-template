package router

import (
	"github.com/gorilla/mux"
)

//InitRoutes Load handlers
func InitRoutes(router *mux.Router) *mux.Router {
	var api = router.PathPrefix("/api/v1").Subrouter()

	for _, route := range routes {
		api.
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandlerFunc)
	}

	return api
}
