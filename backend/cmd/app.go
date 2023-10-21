package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/rs/cors"

	app "github.com/abdulsalam/pixel8labs/config"
	handler "github.com/abdulsalam/pixel8labs/internal/handler/github"
	service "github.com/abdulsalam/pixel8labs/internal/service/github"
	usecase "github.com/abdulsalam/pixel8labs/internal/usecase/github"

	_middleware "github.com/abdulsalam/pixel8labs/middleware"
)

var (
	routes = flag.Bool("routes", false, "Generate router documentation")
	path   = "config/"
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

	// Load services.
	log.Println("Setup service")
	svc := service.New(oauthConfig, config)

	// Load usecase.
	log.Println("Setup usecase")
	uc := usecase.New(svc)

	// Load handler.
	log.Println("Setup handler")
	handler := handler.New(uc)

	// Routes.
	log.Println("Setup routes")
	r := setupRoutes(chi.NewRouter(), handler)

	// Run apps.
	log.Printf("App start on port :%s", config.App.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", config.App.Port), r)
}

func setupRoutes(
	r *chi.Mux,
	handler *handler.Handler,
) *chi.Mux {
	// Create a new CORS middleware with default options.
	corsMiddleware := cors.Default()

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
