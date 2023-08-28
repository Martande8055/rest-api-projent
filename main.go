package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type product struct {
	Id       string
	Name     string
	Quantity int
	Price    int
}

var products []product

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello we are on home page")
}

func productList(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}

func productInfo(w http.ResponseWriter, r *http.Request) {

	myvar := mux.Vars(r)

	key := myvar["id"]
	for _, prod := range products {
		if string(prod.Id) == key {
			json.NewEncoder(w).Encode(prod)
		}
	}
}

func httpResponceHandler() {
	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/products", productList)
	myrouter.HandleFunc("/product/{id}", productInfo)
	myrouter.HandleFunc("/home", homepage)
	http.ListenAndServe("localhost:10000", myrouter)

}

func main() {

	products = []product{
		product{"1", "apple", 10, 100},
		product{"2", "mango", 12, 80},
		product{"3", "orange", 19, 180},
	}

	httpResponceHandler()

}
