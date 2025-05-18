package handlers

import (
	"net/http"

	"go.temporal.io/sdk/client"

	"github.com/go-microfrontend/items-renderer/internal/domain"
	"github.com/go-microfrontend/items-renderer/internal/templates/categories"
)

const (
	CatalogueEndpoint = "GET /{$}"
)

type CatalogueHandler struct {
	client  client.Client
	options *client.StartWorkflowOptions
}

func NewCatalogue(
	client client.Client,
	options *client.StartWorkflowOptions,
) *CatalogueHandler {
	return &CatalogueHandler{
		client:  client,
		options: options,
	}
}

func (h *CatalogueHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	we, _ := h.client.ExecuteWorkflow(
		r.Context(),
		*h.options,
		"GetCategories",
	)

	var cs []domain.Category
	we.Get(r.Context(), &cs)

	categories.Grid(cs).Render(r.Context(), w)
}
