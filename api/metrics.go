package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	cpuUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_usage",
		Help: "CPU usage percentage",
	})
	memoryUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "memory_usage",
		Help: "Memory usage percentage",
	})
)

func init() {
	prometheus.MustRegister(cpuUsage)
	prometheus.MustRegister(memoryUsage)
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
		for _, line := range lines {
			if strings.HasPrefix(line, "cpu_usage") {
				var value float64
				_, err := fmt.Sscanf(line, "cpu_usage %f", &value)
				if err == nil {
					cpuUsage.Set(value)
					fmt.Println("CPU: ", value)
				}
			} else if strings.HasPrefix(line, "memory_usage") {
				var value float64
				_, err := fmt.Sscanf(line, "memory_usage %f", &value)
				if err == nil {
					memoryUsage.Set(value)
					fmt.Println("Mem: ", value)
				}
			}
		}

		time.Sleep(10 * time.Second)
	}
}

func main() {
	go scrapeMetrics()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
