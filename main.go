package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/home", homepage)
	http.ListenAndServe("localhost:10000", nil)

}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello we are on home page")
}
