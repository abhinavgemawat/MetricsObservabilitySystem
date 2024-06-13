package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func init() {
	// Initialize InfluxDB
	InitInfluxDB()
}

func scrapeMetrics() {
	for {
		resp, err := http.Get("http://localhost:5000/metrics")
		if err != nil {
			log.Printf("Error fetching metrics: %v", err)
			time.Sleep(10 * time.Second)
			continue
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading response body: %v", err)
			time.Sleep(10 * time.Second)
			continue
		}

		lines := strings.Split(string(body), "\n")
		var cpuUsage, memoryUsage float64
		for _, line := range lines {
			if strings.HasPrefix(line, "cpu_usage") {
				_, err := fmt.Sscanf(line, "cpu_usage %f", &cpuUsage)
				if err != nil {
					log.Printf("Error parsing cpu_usage: %v", err)
				}
			} else if strings.HasPrefix(line, "memory_usage") {
				_, err := fmt.Sscanf(line, "memory_usage %f", &memoryUsage)
				if err != nil {
					log.Printf("Error parsing memory_usage: %v", err)
				}
			}
		}

		// Write metrics to InfluxDB
		WriteMetrics(cpuUsage, memoryUsage)

		time.Sleep(10 * time.Second)
	}
}

func main() {
	go scrapeMetrics()

	// Start the HTTP server (you can add more handlers if needed)
	http.Handle("/metrics", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Metrics collection running"))
	}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
