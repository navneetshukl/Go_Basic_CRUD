package main

import (
	"go_modules/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	mux := chi.NewRouter()

	mux.Post("/add",routes.AddBook)
	mux.Get("/getall",routes.GetAllBooks)
	mux.Get("/get",routes.GetBook)
	mux.Put("/update",routes.UpdateBook)
	mux.Delete("/delete",routes.DeleteBook)

	http.ListenAndServe(":8080",mux)
}