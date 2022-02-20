package main

import (
	"net/http"

	"github.com/adityagoel/product-api/handlers"
)

func main() {
	http.HandleFunc("/", handlers.DataHandler)
	http.ListenAndServe(":8081", nil)
}
