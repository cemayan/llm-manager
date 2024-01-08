package main

import (
	"context"
	"errors"
	"git-observer/api/server"
	"git-observer/internal/backend"
	"git-observer/internal/config"
	"git-observer/internal/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	config.New()
	log.New()
	server.New()
	backend.Init()
}

func main() {

	server.HttpServer.ConfigureMux()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {

		if err := server.HttpServer.Listen(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.LoggerInstance.Logger.Fatalf("error starting server: %s\n", err)
		}

	}()

	log.LoggerInstance.Logger.Infof("Server started on port: %v\n", config.AppConfig.Config.Serve.Port)

	<-done
	log.LoggerInstance.Logger.Infoln("Server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.HttpServer.Shutdown(ctx); err != nil {
		log.LoggerInstance.Logger.Fatalf("Server Shutdown Failed:%+v", err)
	}
}
