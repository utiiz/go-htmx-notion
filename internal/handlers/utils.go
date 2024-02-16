package handlers

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/utiiz/go/notion/internal/templates/pages"
)

func onError(w http.ResponseWriter, err error, msg string, code int) {
	if err != nil {
		http.Error(w, msg, code)
		log.Println(msg, err)
	}
}

func RenderView(w http.ResponseWriter, r *http.Request, view templ.Component, layoutPath string) {
	if r.Header.Get("Hx-Request") == "true" {
		err := view.Render(r.Context(), w)
		onError(w, err, "Internal server error", http.StatusInternalServerError)
		return
	}
	err := pages.Base(layoutPath).Render(r.Context(), w)
	onError(w, err, "Internal server error", http.StatusInternalServerError)
}
