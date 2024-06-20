package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"llm-manager/api/middleware"
	"llm-manager/api/router"
	"llm-manager/api/router/commit"
	"llm-manager/internal/config"
	"llm-manager/internal/log"
	"net/http"
	"time"
)

// HttpServer is used to react Server instance whenever what you want
var HttpServer *Server

type Server struct {
	privateKey string
	cert       string
	server     *http.Server
	router     *mux.Router
}

// Listen listens and serves
// If cert and key are given tls server will be started.
func (s *Server) Listen() error {
	var err error
	appConfig := config.AppConfig.Config

	if appConfig.Serve.Certificate != "" && appConfig.Serve.PrivateKey != "" {
		err = s.server.ListenAndServeTLS(appConfig.Serve.Certificate, appConfig.Serve.PrivateKey)
	} else {
		err = s.server.ListenAndServe()
	}
	return err
}

// Shutdown returns an error if there is an error.
func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		log.LoggerInstance.Logger.Fatalf("Server Shutdown Failed:%+v", err)
	}

	return nil
}

// configureHealthRoute sets a function for handling to "health" requests
func (s *Server) configureHealthRoute() {
	s.router.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})
}

// configureAPIRoutes sets a function for handling to "api" requests
func (s *Server) configureAPIRoutes() {
	//api routers
	routers := []router.Router{
		commit.NewRouter(),
	}

	for _, router := range routers {
		for _, route := range router.Routes() {
			s.router.HandleFunc(fmt.Sprintf("/api/%v", config.AppConfig.Config.Api.Version)+route.Path(), route.Handler()).
				Methods(route.Method())
		}
	}
}

// ConfigureMux configures routes and middlewares
func (s *Server) ConfigureMux() {
	s.router = mux.NewRouter()

	s.router.Use(middleware.Logger)

	s.configureAPIRoutes()
	s.configureHealthRoute()

	s.server.Handler = s.router
}

// New creates a HttpServer struct
func New() {

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.AppConfig.Config.Serve.Port),
		WriteTimeout: time.Second * 300,
		ReadTimeout:  time.Second * 300,
		IdleTimeout:  time.Second * 300,
	}

	HttpServer = &Server{server: server}
}
