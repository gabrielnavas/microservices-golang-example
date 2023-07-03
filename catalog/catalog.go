package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
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
	println(string(data))

	var products Products
	json.Unmarshal(data, &products)

	return products.Products
}

func main() {
	loadProducts()
}
