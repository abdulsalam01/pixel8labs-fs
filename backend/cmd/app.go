package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/rs/cors"

	app "github.com/abdulsalam/pixel8labs/config"
	handler "github.com/abdulsalam/pixel8labs/internal/handler/github"
	repo "github.com/abdulsalam/pixel8labs/internal/repo/github"
	service "github.com/abdulsalam/pixel8labs/internal/service/github"
	usecase "github.com/abdulsalam/pixel8labs/internal/usecase/github"

	_middleware "github.com/abdulsalam/pixel8labs/middleware"
)

var (
	routes = flag.Bool("routes", false, "Generate router documentation") //nolint:all
	path   = "config/"
	ctx    = context.Background()
)

func main() {
	flag.Parse()

	// Load appConfig.
	config, err := app.LoadAppConfig(path)
	if err != nil {
		return
	}

	// Load oauthConfig.
	oauthConfig := app.LoadOAuthConfig(config)

	// Load inMem database.
	bigCache, err := bigcache.New(ctx, bigcache.Config{
		// Time after which entry can be evicted
		LifeWindow: 10 * time.Minute, // Set per now, since no need to keep by time for a long time.
		// Number of shards (must be a power of 2)
		Shards: 1024,
		// RPS * lifeWindow, used only in initial memory allocation.
		MaxEntriesInWindow: 1000 * 10 * 60,
		// Max entry size in bytes, used only in initial memory allocation.
		MaxEntrySize: 500,
		// Skip spamming prints information about additional memory allocation.
		Verbose: false,
	})
	if err != nil {
		return
	}

	// Load repo.
	log.Println("Setup repo")
	repository := repo.New(bigCache)

	// Load services.
	log.Println("Setup service")
	svc := service.New(oauthConfig, config)

	// Load usecase.
	log.Println("Setup usecase")
	uc := usecase.New(svc, repository)

	// Load handler.
	log.Println("Setup handler")
	handler := handler.New(uc)

	// Routes.
	log.Println("Setup routes")
	r := setupRoutes(chi.NewRouter(), handler)

	// Run apps.
	log.Printf("App start on port :%s", config.App.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", config.App.Port), r) //nolint:all
}

func setupRoutes(
	r *chi.Mux,
	handler *handler.Handler,
) *chi.Mux {
	// Create a new CORS middleware with default options.
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	})

	r.Use(corsMiddleware.Handler)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/api", func(r chi.Router) {
		r.Get("/user", _middleware.GenericMiddleware(handler.GetUser))
		r.Get("/repo", _middleware.GenericMiddleware(handler.GetRepo))

		r.Post("/login", _middleware.GenericMiddleware(handler.GetCode))
		r.Post("/logout", _middleware.GenericMiddleware(handler.GetLogout))
		r.Get("/callback", _middleware.GenericMiddleware(handler.Callback))
	})

	return r
}
