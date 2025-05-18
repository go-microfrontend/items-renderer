package main

import (
	"log/slog"
	"net/http"
	"os"

	"go.temporal.io/sdk/client"

	"github.com/go-microfrontend/items-renderer/internal/handlers"
)

const addr = ":8080"

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	slog.SetLogLoggerLevel(slog.LevelDebug)

	clt, err := client.Dial(client.Options{HostPort: os.Getenv("TEMPORAL_ADDR"), Logger: logger})
	if err != nil {
		slog.Error("failed to connect to temporal", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer clt.Close()

	catalogueHandler := handlers.NewCatalogue(
		clt,
		&client.StartWorkflowOptions{TaskQueue: os.Getenv("TASK_REPO_QUEUE")},
	)

	categoryHandler := handlers.NewCategory(
		clt,
		&client.StartWorkflowOptions{TaskQueue: os.Getenv("TASK_REPO_QUEUE")},
	)

	pageHandler := handlers.NewPage(
		clt,
		&client.StartWorkflowOptions{TaskQueue: os.Getenv("TASK_REPO_QUEUE")},
	)

	server := compose(catalogueHandler, categoryHandler, pageHandler)
	server.ListenAndServe()
}

func compose(catalogueHandler, categoryHandler, pageHandler http.Handler) *http.Server {
	mux := http.NewServeMux()
	mux.Handle(handlers.CategoryEndpoint, categoryHandler)
	mux.Handle(handlers.CatalogueEndpoint, catalogueHandler)
	mux.Handle(handlers.PageEndpoint, pageHandler)

	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}
