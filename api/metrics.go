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

// ScrapeMetrics fetches and parses metrics, storing them in the provided Metrics slice
func ScrapeMetrics(metrics *[]Metric) {
	for {
		resp, err := http.Get("http://localhost:5000/metrics")
		if err != nil {
			log.Printf("Error fetching metrics: %v", err)
			time.Sleep(10 * time.Second)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading response body: %v", err)
			time.Sleep(10 * time.Second)
			continue
		}

		lines := strings.Split(string(body), "\n")
		var cpuUsage, memoryUsage, latency, traffic float64
		var downtime int

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
						Tags:      map[string]string{"type": "cpu"},
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
						Tags:      map[string]string{"type": "memory"},
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
						Tags:      map[string]string{"type": "latency"},
					})
				}
			} else if strings.HasPrefix(line, "traffic") {
				_, err := fmt.Sscanf(line, "traffic %f", &traffic)
				if err != nil {
					log.Printf("Error parsing traffic: %v", err)
				} else {
					*metrics = append(*metrics, Metric{
						Name:      "traffic",
						Value:     traffic,
						Timestamp: time.Now(),
						Tags:      map[string]string{"type": "traffic"},
					})
				}
			} else if strings.HasPrefix(line, "downtime") {
				_, err := fmt.Sscanf(line, "downtime %d", &downtime)
				if err != nil {
					log.Printf("Error parsing downtime: %v", err)
				} else {
					*metrics = append(*metrics, Metric{
						Name:      "downtime",
						Value:     float64(downtime),
						Timestamp: time.Now(),
						Tags:      map[string]string{"type": "downtime"},
					})
				}
			}
		}

		time.Sleep(10 * time.Second)
	}
}
