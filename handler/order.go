package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

type Stock struct {
	Symbol string
	Price  float64
}
type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("created an account")
}
func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updated an account")
}
func (o *Order) DeleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleted an account")
}
func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("read an account")
}
func (o *Order) GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("read an account")
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
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

func StockHandler(w http.ResponseWriter, r *http.Request) {
	//TODO implement api
	stocks := []Stock{
		{"AAPL", 130.21},
		{"GOOGL", 2512.56},
	}

	tmpl, err := template.ParseFiles("templates/stocks.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template, passing any necessary data (nil in this case)
	err = tmpl.Execute(w, stocks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
