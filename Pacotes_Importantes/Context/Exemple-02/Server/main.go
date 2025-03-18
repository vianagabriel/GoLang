package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Carro struct {
	ID     int
	Modelo string
	Marca  string
	Ano    string
	Preco  float64
}

var Carros = []Carro{
	{ID: 1, Modelo: "Vectra", Marca: "Chevrolet", Ano: "2002", Preco: 22000.00},
	{ID: 2, Modelo: "Astra", Marca: "Chevrolet", Ano: "2007", Preco: 40000.00},
}

func main() {
	http.HandleFunc("/", ListarCarros)
	http.ListenAndServe(":8081", nil)
}

func ListarCarros(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	case <-time.After(1 * time.Second):
		json.NewEncoder(w).Encode(Carros)
	case <-ctx.Done():
		http.Error(w, "Requisição cancelada", http.StatusRequestTimeout)
	}
}
