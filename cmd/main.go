package main

import (
	"log"
	"net/http"
	"todo/handler"
)

func main() {
	addr := ":8080"
	log.Printf("HTTP server started on %s", addr)
	log.Fatal(http.ListenAndServe(addr, handler.NewRouter()))
}
