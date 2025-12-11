package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "time"

    "github.com/catsmells/ltak-server/internal/api"
    "github.com/catsmells/ltak-server/internal/config"
    "github.com/catsmells/ltak-server/internal/db"
    "github.com/catsmells/ltak-server/internal/websocket"
)

func main() {
    cfg := config.LoadOrDefault()

    sqldb, err := db.Open(cfg.DBPath)
    if err != nil {
        log.Fatal(err)
    }
    defer sqldb.Close()

    hub := websocket.NewHub()
    go hub.Run()

    r := api.NewRouter(cfg, sqldb, hub)

    srv := &http.Server{
        Addr:         cfg.ListenAddr,
        Handler:      r,
        ReadTimeout:  15 * time.Second,
        WriteTimeout: 30 * time.Second,
        IdleTimeout:  60 * time.Second,
    }

    go func() {
        log.Printf("LTAK Server listening on %s", cfg.ListenAddr)
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Server error: %v", err)
        }
    }()

    // graceful shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt)
    <-quit
    log.Println("shutting down server...")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    _ = srv.Shutdown(ctx)
    log.Println("server stopped")
}
