package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("created an account")
}
func (o *Order) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updated an account")
}
func (o *Order) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleted an account")
}
func (o *Order) Read(w http.ResponseWriter, r *http.Request) {
	fmt.Println("read an account")
}
