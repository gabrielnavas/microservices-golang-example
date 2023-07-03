package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"text/template"

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

var productUrl string

func init() {
	productUrl = os.Getenv("PRODUCT_URL")
}

func loadProducts() []Product {
	response, err := http.Get(productUrl + "/products")
	if err != nil {
		println(err)
	}
	data, _ := io.ReadAll(response.Body)

	var products Products
	json.Unmarshal(data, &products)

	return products.Products
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", ListProducts).Methods("GET")
	r.HandleFunc("/product/{id}", ShowProduct).Methods("GET")
	http.ListenAndServe(":8080", r)
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	products := loadProducts()
	t := template.Must(template.ParseFiles("templates/catalog.html"))
	t.Execute(w, products)
}

func ShowProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response, err := http.Get(productUrl + "/products/" + vars["id"])
	if err != nil {
		println("The HTTP request falied with herror %s\n", err)
	}
	data, _ := io.ReadAll(response.Body)

	var product Product
	json.Unmarshal(data, &product)

	t := template.Must(template.ParseFiles("templates/view.html"))
	t.Execute(w, product)
}
