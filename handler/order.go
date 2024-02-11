package handler

import "net/http"

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Order created..."))
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Order listed..."))
}

func (o *Order) ListById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Order listed by Id"))
}

func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Order updated by Id"))
}

func (o *Order) DeleteById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Order deleted by Id"))
}
