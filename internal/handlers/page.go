package handlers

import (
	"net/http"

	"go.temporal.io/sdk/client"

	"github.com/go-microfrontend/items-renderer/internal/domain"
	"github.com/go-microfrontend/items-renderer/internal/templates/page"
)

const (
	PageEndpoint = "GET /product/{id}"
)

type PageHandler struct {
	client  client.Client
	options *client.StartWorkflowOptions
}

func NewPage(
	client client.Client,
	options *client.StartWorkflowOptions,
) *PageHandler {
	return &PageHandler{
		client:  client,
		options: options,
	}
}

func (h *PageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	we1, _ := h.client.ExecuteWorkflow(
		r.Context(),
		*h.options,
		"GetProductByID",
		r.PathValue("id"),
	)

	we2, _ := h.client.ExecuteWorkflow(
		r.Context(),
		*h.options,
		"GetProductCharacteristicByID",
		r.PathValue("id"),
	)

	var p domain.Product
	we1.Get(r.Context(), &p)

	var pc domain.ProductCharacteristic
	we2.Get(r.Context(), &pc)

	page.Product(p, pc).Render(r.Context(), w)
}
