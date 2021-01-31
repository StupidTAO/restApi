package router

import (
	"net/http"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.handlerFunc
		//handler = Logger(handler, route.Name)

		router.
			Methods(route.method).
			Path(route.pattern).
			Name(route.name).
			Handler(handler)

	}
	return router
}