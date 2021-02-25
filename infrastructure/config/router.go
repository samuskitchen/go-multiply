package config

import (
	"github.com/go-chi/chi"
	v1 "go-multiply/domain/multiply/application/handler/v1"
	"net/http"
)

// Routes returns the API V1 Handler with configuration.
func Routes() http.Handler {
	router := chi.NewRouter()

	mh := v1.NewMultiplyHandler()
	router.Mount("/multiply", routesMultiply(mh))

	return router
}

// routesMultiply returns router with each endpoint.
func routesMultiply(v1 *v1.MultiplyRouter) http.Handler {
	router := chi.NewRouter()

	router.Get("/", v1.GetMultiplyNumbers)

	return router
}