package todoapp

import (
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(routes *chi.Mux) {
	routes.HandleFunc("/", IndexView)
	routes.HandleFunc("/create", CreateEntry)
	routes.Post("/{id}/delete", DeleteEntry)
}
