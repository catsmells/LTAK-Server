package api

import (
    "net/http"
    "path/filepath"

    "github.com/go-chi/chi/v5"
    "github.com/jmoiron/sqlx"
    "github.com/catsmells/ltak-server/internal/config"
    "github.com/catsmells/ltak-server/internal/tiles"
    "github.com/catsmells/ltak-server/internal/users"
    "github.com/catsmells/ltak-server/internal/markers"
    "github.com/catsmells/ltak-server/internal/websocket"
)

func NewRouter(cfg *config.Config, db *sqlx.DB, hub *websocket.Hub) http.Handler {
    r := chi.NewRouter()

    // Static assets
    r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

    // Tiles (GET)
    tileStore := tiles.NewFileStore(cfg.TileRoot)
    r.Get("/tiles/{z}/{x}/{y}.png", tiles.Handler(tileStore))

    // Users
    userStore := users.NewStore(db)
    r.Post("/users/position", users.UpdatePositionHandler(userStore, hub))
    r.Get("/users/positions", users.GetPositionsHandler(userStore))

    // Markers
    markerStore := markers.NewStore(db)
    r.Route("/markers", func(r chi.Router) {
        r.Get("/", markers.ListHandler(markerStore))
        r.Post("/", markers.CreateHandler(markerStore, hub))
        r.Put("/{id}", markers.UpdateHandler(markerStore, hub))
        r.Delete("/{id}", markers.DeleteHandler(markerStore, hub))
    })

    // WebSocket for realtime updates
    r.Get("/ws", websocket.ServeWs(hub))

    return r
}
