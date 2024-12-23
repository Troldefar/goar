package main

import (
	"net/http"
	"github.com/Troldefar/goar/internal/handler"
  "github.com/Troldefar/goar/functions"
  "fmt"
  "os"
)

func main() {
	http.HandleFunc("/", handler.HomeHandler)

	port := ":8080"

	http.ListenAndServe(port, nil)

  dat, err := os.ReadFile("test.txt")

  functions.CheckErr(err)

  fmt.Print(string(dat))
}
