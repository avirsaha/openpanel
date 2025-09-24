package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "gopanel/config"
    "gopanel/api"
    "gopanel/internal/db"
    "gopanel/internal/system"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig("config/config.yaml")
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    // Initialize services
    err = db.Init(cfg.Database)
    if err != nil {
        log.Fatalf("Failed to connect to DB: %v", err) }

    err = system.CheckDependencies()
    if err != nil {
        log.Fatalf("Missing system dependencies: %v", err)
    }

    // Set up API routes
    router := api.SetupRouter(cfg)

    // Optional: Serve static frontend
    fs := http.FileServer(http.Dir("./web/static"))
    router.Handle("/static/", http.StripPrefix("/static/", fs))

    // Start HTTP server
    addr := fmt.Sprintf(":%d", cfg.Server.Port)
    log.Printf("Starting OpenPanel server on %s", addr)

    if err := http.ListenAndServe(addr, router); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}

