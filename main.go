package main

import (
	"log"
	"net/http"

	"github.com/abhinavgemawat/MetricsObservabilitySystem/api"
)

func main() {

	go api.ScrapeMetrics()
	// Start the HTTP server (you can add more handlers if needed)
	http.Handle("/metrics", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Metrics collection running"))
	}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
