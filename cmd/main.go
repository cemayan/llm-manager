package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"git-observer/api/server"
	"git-observer/api/types"
	"git-observer/api/util"
	"git-observer/internal/backend"
	"git-observer/internal/config"
	"git-observer/internal/log"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var llmBackend string
var prompt string

func init() {

	config.New()
	log.New()

	cobra.OnInitialize()
	rootCmd.PersistentFlags().StringVarP(&llmBackend, "backend", "b", "", "backend")
	rootCmd.PersistentFlags().StringVarP(&prompt, "prompt", "p", "", "prompt")

	server.New()
	backend.Init()
}

var rootCmd = &cobra.Command{
	Use:   "git-observer",
	Short: "git-observer",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func main() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if prompt != "" {

		switch config.AppConfig.Config.Api.Backend {
		case "ollama":
			query, err := backend.BackendInstance.
				Query(util.PrepareOllamaRequest(config.AppConfig.Config.Ollama.Model, prompt).Marshall(), nil)
			if err != nil {
				fmt.Println(string(util.PrepareErrorResponse(err, 1003).Marshall()))
				return
			}

			var ollamaResp types.OllamaResponse
			err = json.Unmarshal(query, &ollamaResp)

			if err != nil {
				fmt.Println(string(util.PrepareErrorResponse(err, 1001).Marshall()))
				return
			}

			fmt.Println(string(util.PrepareResponse(ollamaResp.Response).Marshall()))
		case "lingoose":
			query, err := backend.BackendInstance.
				Query([]byte(prompt), nil)
			if err != nil {
				fmt.Println(string(util.PrepareErrorResponse(err, 1003).Marshall()))
				return
			}

			fmt.Println(string(util.PrepareResponse(query).Marshall()))
		case "langchaingo":
			query, err := backend.BackendInstance.
				Query([]byte(prompt), nil)
			if err != nil {
				fmt.Println(string(util.PrepareErrorResponse(err, 1003).Marshall()))
				return
			}

			fmt.Println(string(util.PrepareResponse(query).Marshall()))
		}

		return
	}

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
