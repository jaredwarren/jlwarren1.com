package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.Handle("/.well-known/", http.StripPrefix("/.well-known/", fs))

	listenPort := os.Getenv("LISTEN_PORT")
	// for testing locally where port 80 is not available
	if listenPort == "" {
		listenPort = "8082"
	}
	addr := fmt.Sprintf(":%s", listenPort)
	fmt.Println("listening on:", addr)
	fmt.Println(http.ListenAndServe(addr, nil))
}
