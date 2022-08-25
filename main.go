package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	fmt.Println("web server running...")

	http.HandleFunc("/", helloWorld)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("web server err:", err)
	}
}
