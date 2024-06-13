package api

import (
	"log"
	"net/http"
)

func main() {
	go scrapeMetrics()

	// Start the HTTP server (you can add more handlers if needed)
	http.Handle("/metrics", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Metrics collection running"))
	}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
