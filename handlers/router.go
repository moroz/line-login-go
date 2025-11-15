package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/moroz/line-login-go/config"
)

func Router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)

	page := PageController()
	r.Get("/", page.Index)

	oauth2 := OAuth2Controller(config.LineClientId, config.LineClientSecret)
	r.Get("/oauth/line/redirect", oauth2.Init)

	return r
}
