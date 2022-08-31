package main

import (
	"fmt"
	"go-demo1/app"
	"log"
	"net"
	"net/http"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi Hello World!")
}

func BuildServer() (*http.Server, net.Listener) {
	fmt.Println("web server running...")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // 如果很簡單，你可以把實作直接寫進來，不需要額外再寫個function在外面
		http.FileServer(http.Dir("./app/urls/pages/")).ServeHTTP(w, r)
	})
	mux.HandleFunc("/about/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, app.Version)
	})
	mux.HandleFunc("/hi/", hi) // 建議結尾也補上 /

	// http.ListenAndServe(":8080", nil) // 對外網都有效
	// err := http.ListenAndServe("127.0.0.1:0", mux) // 純本機，所以防火牆不會彈出來 // 因為我們希望獲得動態port的埠號，所以不能靠ListenAndServer，它封裝在裡面，又是私有的所以取不到
	server := &http.Server{Addr: "127.0.0.1:0", Handler: mux} // port: 0會自動分配
	listener, err := net.Listen("tcp", server.Addr)
	if err != nil {
		log.Fatal("web server err:", err)
	}
	return server, listener

}

func main() {
	server, listener := BuildServer()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()
	time.Sleep(50 * time.Millisecond) // wait server ready
	port := fmt.Sprintf("%d", listener.Addr().(*net.TCPAddr).Port)
	if runtime.GOOS == "windows" {
		if err := exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://127.0.0.1:"+port).Start(); err != nil {
			panic(err)
		}
	}
	wg.Wait()
}
