package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"html/template"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/hello", basicHandler)
	router.Get("/", dashboardHandler)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to listen to server", err)
	}
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML template
	tmpl, err := template.ParseFiles("templates/dashboard.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template, passing any necessary data (nil in this case)
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func minorChanges() {

}
