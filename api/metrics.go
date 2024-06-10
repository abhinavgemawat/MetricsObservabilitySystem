package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// var (
// 	mode = flag.String("mode", "pull", "Mode of operation: pull or push")
// )

func main() {
	flag.Parse()

	// Initialize Prometheus metrics

	counter := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "example_counter",
		Help: "Example of a counter metric",
	})

	prometheus.MustRegister(counter)

	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Metrics Collection Server started in pull mode at http://localhost:9301/metrics")
	http.ListenAndServe(":5000", nil)

	// if *mode == "pull" {
	// 	// Pull mode: expose metrics for scraping by Prometheus
	// 	http.Handle("/metrics", promhttp.Handler())

	// 	fmt.Println("Metrics Collection Server started in pull mode at http://localhost:9301/metrics")
	// 	http.ListenAndServe(":9301", nil)

	// } else if *mode == "push" {
	// 	// Push mode: accept metric data via HTTP endpoint
	// 	http.HandleFunc("/push", func(w http.ResponseWriter, r *http.Request) {
	// 		// Parse metric data from request body and update Prometheus metrics
	// 		// Example: counter.Inc()

	// 		fmt.Println("Received pushed metrics")
	// 		w.WriteHeader(http.StatusOK)
	// 	})

	// 	fmt.Println("Metrics Collection Server started in push mode at http://localhost:8080/push")
	// 	http.ListenAndServe(":8080", nil)
	// } else {
	// 	fmt.Println("Invalid mode specified. Please use 'pull' or 'push'.")
	// }

}
