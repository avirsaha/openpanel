package api

import (
    "net/http"

    "gopanel/api/handlers"
    "gopanel/config"
    "gopanel/middleware"
)

func SetupRouter(cfg *config.Config) *http.ServeMux {
    mux := http.NewServeMux()

    // Auth routes
    mux.HandleFunc("/api/v1/login", middleware.Apply(handlers.LoginHandler, middleware.Logging))
    mux.HandleFunc("/api/v1/logout", middleware.Apply(handlers.LogoutHandler, middleware.Logging))

    // User routes
    mux.HandleFunc("/api/v1/users", middleware.Apply(handlers.UsersHandler, middleware.Auth, middleware.Logging))

    // Web/hosting routes
    mux.HandleFunc("/api/v1/web", middleware.Apply(handlers.WebHandler, middleware.Auth, middleware.Logging))

    return mux
}

