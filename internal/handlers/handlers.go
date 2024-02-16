package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/utiiz/go/notion/internal/templates/pages"
)

func HomeGetHandler(w http.ResponseWriter, r *http.Request) {
	page := pages.Base()
	templ.Handler(page)
}
