package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Units
const (
	_        = iota
	KB int64 = 1 << (10 * iota)
	MB
	GB
	TB
)

var cert string
var key string
var port string

func init() {
	flag.StringVar(&cert, "cert", "", "give me a certificate")
	flag.StringVar(&key, "key", "", "give me a key")
	flag.StringVar(&port, "port", "80", "give me a port number")
}

func main() {
	fmt.Println("........")
	flag.Parse()

	// fs := http.FileServer(http.Dir("static"))
	// http.Handle("/", fs)
	// http.HandleFunc("/", serveTemplate)
	// http.Handle("/.well-known/", http.StripPrefix("/.well-known/", fs))

	http.HandleFunc("/data", dataHandler)
	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/bench", benchHandler)
	http.HandleFunc("/", whoamiHandler)
	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/health", healthHandler)

	// for testing locally where port 80 is not available
	// if listenPort == "" {
	// 	listenPort = "8082"
	// }
	if len(cert) > 0 && len(key) > 0 {
		fmt.Println("Starting TLS:", port)
		log.Fatal(http.ListenAndServeTLS(":"+port, cert, key, nil))
	}
	fmt.Println("Starting:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--------")
	fmt.Printf("%+v\n", r)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("----data----")
	fmt.Printf("%+v\n", r)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("----echo----")
	fmt.Printf("%+v\n", r)
}

func benchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("----bench----")
	fmt.Printf("%+v\n", r)
}

func whoamiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("----whoami----")
	fmt.Printf("%+v\n", r)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("----api----")
	fmt.Printf("%+v\n", r)
}

type healthState struct {
	StatusCode int
}

var currentHealthState = healthState{http.StatusOK}
var mutexHealthState = &sync.RWMutex{}

func healthHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		var statusCode int

		if err := json.NewDecoder(req.Body).Decode(&statusCode); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("Update health check status code [%d]\n", statusCode)

		mutexHealthState.Lock()
		defer mutexHealthState.Unlock()
		currentHealthState.StatusCode = statusCode
	} else {
		mutexHealthState.RLock()
		defer mutexHealthState.RUnlock()
		w.WriteHeader(currentHealthState.StatusCode)
	}
}
