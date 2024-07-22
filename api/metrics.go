package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// func init() {
// 	// Initialize InfluxDB
// 	InitInfluxDB()
// }

// Metric represents a single metric data point
type Metric struct {
	Name      string            `json:"name"`
	Value     float64           `json:"value"`
	Timestamp time.Time         `json:"timestamp"`
	Tags      map[string]string `json:"tags"`
}

// Serialize serializes the Metric to JSON
func (m *Metric) Serialize() (string, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return "", fmt.Errorf("failed to serialize metric: %v", err)
	}
	return string(data), nil
}

// DeserializeMetric deserializes the JSON string to a Metric
func DeserializeMetric(data string) (*Metric, error) {
	var metric Metric
	err := json.Unmarshal([]byte(data), &metric)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize metric: %v", err)
	}
	return &metric, nil
}

// func ScrapeMetrics() {
// 	for {
// 		resp, err := http.Get("http://localhost:5000/metrics")
// 		if err != nil {
// 			log.Printf("Error fetching metrics: %v", err)
// 			time.Sleep(10 * time.Second)
// 			continue
// 		}
// 		// defer resp.Body.Close()

// 		body, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			log.Printf("Error reading response body: %v", err)
// 			time.Sleep(10 * time.Second)
// 			continue
// 		}

// 		lines := strings.Split(string(body), "\n")
// 		var cpuUsage, memoryUsage, latency float64
// 		for _, line := range lines {
// 			if strings.HasPrefix(line, "cpu_usage") {
// 				_, err := fmt.Sscanf(line, "cpu_usage %f", &cpuUsage)
// 				if err != nil {
// 					log.Printf("Error parsing cpu_usage: %v", err)
// 				}
// 			} else if strings.HasPrefix(line, "memory_usage") {
// 				_, err := fmt.Sscanf(line, "memory_usage %f", &memoryUsage)
// 				if err != nil {
// 					log.Printf("Error parsing memory_usage: %v", err)
// 				}
// 			} else if strings.HasPrefix(line, "latency") {
// 				_, err := fmt.Sscanf(line, "latency %f", &latency)
// 				if err != nil {
// 					log.Printf("Error parsing latency: %v", err)
// 				}
// 			}
// 		}

// 		// Write metrics to InfluxDB
// 		// TODO: Add latency
// 		WriteMetrics()

// 		time.Sleep(10 * time.Second)
// 	}
// }

// latency
// average downtime per week
// traffic
// what percentage of time are the servers 80%/50%/20% of the load
// Additional - dashboard

// / ScrapeMetrics fetches and parses metrics, storing them in the provided Metrics slice
func ScrapeMetrics(metrics *[]Metric) {
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
		var cpuUsage, memoryUsage, latency float64
		for _, line := range lines {
			if strings.HasPrefix(line, "cpu_usage") {
				_, err := fmt.Sscanf(line, "cpu_usage %f", &cpuUsage)
				if err != nil {
					log.Printf("Error parsing cpu_usage: %v", err)
				} else {
					*metrics = append(*metrics, Metric{
						Name:      "cpu_usage",
						Value:     cpuUsage,
						Timestamp: time.Now(),
						Tags:      map[string]string{},
					})
				}
			} else if strings.HasPrefix(line, "memory_usage") {
				_, err := fmt.Sscanf(line, "memory_usage %f", &memoryUsage)
				if err != nil {
					log.Printf("Error parsing memory_usage: %v", err)
				} else {
					*metrics = append(*metrics, Metric{
						Name:      "memory_usage",
						Value:     memoryUsage,
						Timestamp: time.Now(),
						Tags:      map[string]string{},
					})
				}
			} else if strings.HasPrefix(line, "latency") {
				_, err := fmt.Sscanf(line, "latency %f", &latency)
				if err != nil {
					log.Printf("Error parsing latency: %v", err)
				} else {
					*metrics = append(*metrics, Metric{
						Name:      "latency",
						Value:     latency,
						Timestamp: time.Now(),
						Tags:      map[string]string{},
					})
				}
			}
		}

		time.Sleep(10 * time.Second)
	}
}
