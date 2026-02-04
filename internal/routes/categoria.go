package routes

import (
	"myapi/internal/handlers"

	"github.com/gorilla/mux"
)

func CategoriaRoutes(r *mux.Router) {
	r.HandleFunc("/api/categorias", handlers.ListCategorias).Methods("GET")
	r.HandleFunc("/api/categorias/{id}", handlers.GetCategoria).Methods("GET")
	r.HandleFunc("/api/categorias", handlers.CreateCategoria).Methods("POST")
	r.HandleFunc("/api/categorias", handlers.UpdateCategoria).Methods("PUT")
	r.HandleFunc("/api/categorias/{id}", handlers.DeleteCategoria).Methods("DELETE")
}
