package handlers

import (
	"net/http"

	"go.temporal.io/sdk/client"

	"github.com/go-microfrontend/items-renderer/internal/domain"
	"github.com/go-microfrontend/items-renderer/internal/templates/products"
)

const (
	CategoryEndpoint = "GET /{category}"
)

type CategoryHandler struct {
	client  client.Client
	options *client.StartWorkflowOptions
}

func NewCategory(
	client client.Client,
	options *client.StartWorkflowOptions,
) *CategoryHandler {
	return &CategoryHandler{
		client:  client,
		options: options,
	}
}

func (h *CategoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	we, _ := h.client.ExecuteWorkflow(
		r.Context(),
		*h.options,
		"GetProductsByCategory",
		r.PathValue("category"),
	)

	var ps []domain.Product
	we.Get(r.Context(), &ps)

	products.Grid(ps).Render(r.Context(), w)
}
