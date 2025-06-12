package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	return &WebServer{
		Router:        chi.NewRouter(),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) Start() {
	http.ListenAndServe(s.WebServerPort, s.Router)
}
