package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"llm-manager/api/server"
	"llm-manager/internal/backend"
	"llm-manager/internal/cli"
	"llm-manager/internal/config"
	"llm-manager/internal/log"
	"llm-manager/internal/structs"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var prompt string
var output string

func init() {

	// config initializer
	config.New()
	// logrus initializer
	log.New()

	cobra.OnInitialize()
	rootCmd.PersistentFlags().StringVarP(&prompt, "prompt", "p", "", "prompt for llm backend")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "output to file(json,text)")

	// server initializer
	server.New()
	// backend initializer
	backend.Init()
}

var rootCmd = &cobra.Command{
	Use:   "llm-manager",
	Short: "llm-manager",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func main() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// If output is not empty output will be set to output type
	if output != "" {
		if v, _ok := structs.OutputMap[output]; _ok {
			config.AppConfig.Config.Api.Output = v
		}
	}

	// If prompt is not empty it means app runs by cli
	if prompt != "" {
		cli.Exec(prompt)
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
