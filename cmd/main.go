package main

import (
	"log/slog"
	"net/http"
	"os"

	"go.temporal.io/sdk/client"

	"github.com/go-microfrontend/host-page/internal/domain"
	"github.com/go-microfrontend/host-page/internal/templates"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	dialer, err := client.Dial(client.Options{HostPort: os.Getenv("TEMPORAL_ADDR")})
	if err != nil {
		slog.Error("failed to connect to temporal", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer dialer.Close()

	http.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		var item domain.Item
		f, err := dialer.ExecuteWorkflow(
			r.Context(),
			client.StartWorkflowOptions{TaskQueue: os.Getenv("TASK_QUEUE")},
			"GetItemByIDWF",
			id,
		)
		if err != nil {
			slog.Error("failed to execute workflow", slog.String("error", err.Error()))
			http.Error(w, "Internal", http.StatusInternalServerError)
			return
		}

		err = f.Get(r.Context(), &item)
		if err != nil {
			slog.Error("failed to get future", slog.String("error", err.Error()))
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		component := templates.Item(&item)
		err = component.Render(r.Context(), w)
		if err != nil {
			slog.Error("failed to render", slog.String("error", err.Error()))
			http.Error(w, "Internal", http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe("0.0.0.0:42069", nil)
}
