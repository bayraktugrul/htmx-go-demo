package handlers

import (
	"github.com/bayraktugrul/htmx-go-demo/internal/templates"
	"net/http"
	"time"
)

type TimeHandLer struct{}

func NewTimeHandler() *TimeHandLer {
	return &TimeHandLer{}
}

func (h *TimeHandLer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := templates.Time(time.Now().Format("2006-01-02 15:04:05")).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
