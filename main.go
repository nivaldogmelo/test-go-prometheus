package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request to /service")
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/service", HelloHandler)

	log.Println("Starting to serve at port 8080...")
	http.ListenAndServe(":8080", nil)
	
}
