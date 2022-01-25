package main

import (
	"coredemo/framework"
	"fmt"
	"net/http"
)

func main() {
	server := &http.Server{
		Handler: framework.NewCore(),
		Addr:    "localhost:8080",
	}
	fmt.Println("i love baby")
	server.ListenAndServe()
}
