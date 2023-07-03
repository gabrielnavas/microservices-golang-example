package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Product struct {
	Uuid        string  `json:"uuid"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price,string"`
}

type Products struct {
	Products []Product
}

func loadData() []byte {
	jsonFile, err := os.Open("./products.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Close()
	data, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	products := loadData()
	w.Write(products)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := loadData()

	var products Products
	json.Unmarshal(data, &products)

	for _, p := range products.Products {
		if p.Uuid == vars["id"] {
			product, _ := json.Marshal(p)
			w.Write([]byte(product))
		}
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products/{id}", GetProductById).Methods("GET")
	r.HandleFunc("/products", ListProducts).Methods("GET")
	http.ListenAndServe(":8081", r)
}
