package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pindarEng/stockSimulator/handler"
	"net/http"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Get("/dashboard", handler.DashboardHandler)
	router.Get("/stocks", handler.StockHandler)

	router.Route("/orders", loadOrderRoutes)
	return router
}

func loadOrderRoutes(router chi.Router) {
	orderHandler := &handler.Order{}

	router.Post("/", orderHandler.Create)
	router.Get("/", orderHandler.GetAll)
	router.Get("/{id}", orderHandler.GetById)
	router.Put("/{id}", orderHandler.UpdateById)
	router.Delete("/{id}", orderHandler.DeleteById)
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func minorChanges() {

}
