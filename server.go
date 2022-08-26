package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index!")
}

func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi Hello World!")
}

func main() {
	fmt.Println("web server running...")

	http.HandleFunc("/", index)
	http.HandleFunc("/hi", hi)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("web server err:", err)
	}
}
