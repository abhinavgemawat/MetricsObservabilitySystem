package api

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

const (
	influxToken = ""
	bucket      = "cpu_metrics"
	org         = "example"
	influxURL   = "http://localhost:5001"
)

var (
	Client influxdb2.Client
	// WriteAPI influxdb2.WriteAPIBlocking
	// WriteAPI influxdb2.
)

// Look into nonblocking

func InitInfluxDB() {
	Client = influxdb2.NewClient(influxURL, influxToken)
	// WriteAPI = Client.WriteAPIBlocking(org, bucket)
}
