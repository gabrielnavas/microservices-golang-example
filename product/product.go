package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

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

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products", ListProducts)
	http.ListenAndServe(":8081", r)
}
