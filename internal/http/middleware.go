package http

import (
	"github.com/gorilla/handlers"
	"net/http"
)

func CORS() func(http.Handler) http.Handler {
	headersOk := handlers.AllowedHeaders([]string{"Proto"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	return handlers.CORS(headersOk, originsOk, methodsOk, handlers.AllowCredentials())

}
