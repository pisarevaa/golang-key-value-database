package router

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	"kvdatabase/internal/api"
)

// Создание роутера.
func NewRouter(handlers *api.Handler) http.Handler{
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/key/{key}", handlers.GetValue)
	mux.HandleFunc("POST /v1/key/{key}", handlers.SetValue)
	mux.HandleFunc("DELETE /v1/key/{key}", handlers.DeleteValue)
	mux.Handle("GET /swagger/", httpSwagger.WrapHandler)

	var handler http.Handler = mux
	handler = api.RecoveryMiddleware(handler)
	handler = api.NewLoggingMiddleware()(handler)

	return handler
}