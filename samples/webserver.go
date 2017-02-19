package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handleHello)
	log.Println("serving on http://localhost:7777/hello")
	http.ListenAndServe("localhost:7777", nil) // HL
}

func handleHello(w http.ResponseWriter, req *http.Request) {
	log.Println("serving", req.URL)
	fmt.Fprintln(w, "Hello, FinistGo")
}
