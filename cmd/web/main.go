package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /post/view/{id}", postView)
	mux.HandleFunc("GET /post/create", postCreate)
	mux.HandleFunc("POST /post/create", postCreatePOST)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
