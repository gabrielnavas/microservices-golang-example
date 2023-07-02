package main

import (
	"io/ioutil"
	"log"
	"os"
)

func loadData() []byte {
	jsonFile, err := os.Open("./products.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Close()
	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}

func main() {
	log.Println(string(loadData()))
}
