package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/goar/internal/handler"
)

func main() {
	http.HandleFunc("/", handler.HomeHandler)

	port := ":8080"

	http.ListenAndServe(port, nil)
}
