package main

import (
	"net/http"
	"github.com/Troldefar/goar/internal/handler"
)

func main() {
	http.HandleFunc("/", handler.HomeHandler)

	port := ":8080"

	http.ListenAndServe(port, nil)
}
