package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vansh845/go-microservices/handler"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Route("/orders", loadOrderRoutes)

	return router
}

func loadOrderRoutes(r chi.Router) {

	orderHandler := &handler.Order{}

	r.Post("/", orderHandler.Create)
	r.Get("/", orderHandler.List)
	r.Get("/{id}", orderHandler.ListById)
	r.Post("/{id}", orderHandler.UpdateById)
	r.Delete("/{id}", orderHandler.DeleteById)

}
