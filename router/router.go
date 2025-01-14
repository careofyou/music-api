package router

import (
	"net/http"

	"github.com/careofyou/music-api/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Routes() http.Handler {
    router := chi.NewRouter()
    router.Use(middleware.Recoverer)
    router.Use(cors.Handler(cors.Options {
        AllowedOrigins: []string{"http://*", "https://*"},
        AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders: []string{"Link"},
        AllowCredentials: true,
        MaxAge: 300,
    }))

    router.Get("/api/v1/songs", controllers.GetAllSongs)
    router.Get("/api/v1/songs/song/{id}", controllers.GetSongById)
    router.Post("/api/v1/songs/song", controllers.CreateSong)
    router.Put("/api/v1/songs/song/{id}", controllers.UpdateSong)
    router.Delete("/api/v1/songs/song/{id}", controllers.DeleteSong)

    return router
}
