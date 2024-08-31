package handlers

import (
	"github.com/bayraktugrul/htmx-go-demo/internal/templates"
	"net/http"
)

type HomePageHandler struct{}

func NewHomePageHandler() *HomePageHandler {
	return &HomePageHandler{}
}

func (h *HomePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.Home()
	err := templates.Layout(c, "Home").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
