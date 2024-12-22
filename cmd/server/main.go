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

	fmt.Println("we ok at port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
