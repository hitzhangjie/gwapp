package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", cgiHello)

	log.Printf("http serve localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("http serve error %v", err)
	}
}

func cgiHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
