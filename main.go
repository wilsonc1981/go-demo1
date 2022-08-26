package main

import (
	"fmt"
	"go-demo1/app"
	"log"
	"net/http"
)

func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi Hello World!")
}

func main() {
	fmt.Println("web server running...")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // 如果很簡單，你可以把實作直接寫進來，不需要額外再寫個function在外面
		_, _ = w.Write([]byte("index!"))
	})
	mux.HandleFunc("/about/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, app.Version)
	})
	mux.HandleFunc("/hi/", hi) // 建議結尾也補上 /

	// http.ListenAndServe(":8080", nil) // 對外網都有效
	err := http.ListenAndServe("127.0.0.1:8080", mux) // 純本機，所以防火牆不會彈出來
	if err != nil {
		log.Fatal("web server err:", err)
	}
}
