package main

import (
	"net/http"
)

func main() {
	http.HandlerFunc("/identify_image", handlers.identifyImage())

	srv, err := http.ListenAndServe()
}
