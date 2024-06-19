package api

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

const (
	// oCB_0br7Ev6M16YeMowICZaToJECOb94sSPaKSsnk9xgYOhJVhK9y7sag-fWaeXv50vzX2ej86sII0vns_HlFQ==
	influxToken = "oCB_0br7Ev6M16YeMowICZaToJECOb94sSPaKSsnk9xgYOhJVhK9y7sag-fWaeXv50vzX2ej86sII0vns_HlFQ=="
	bucket      = "metrics"
	org         = "COMP41720"
	influxURL   = "http://localhost:8086"
)

var (
	influxClient = influxdb2.NewClient(influxURL, influxToken)
	writeAPI     = influxClient.WriteAPIBlocking(org, bucket)
)

// Look into nonblocking

func InitInfluxDB() {
	influxClient = influxdb2.NewClient(influxURL, influxToken)
	writeAPI = influxClient.WriteAPIBlocking(org, bucket)
}

func WriteMetrics(cpuUsage, memoryUsage float64) {
	// // Create a point using the fluent API
	p := influxdb2.NewPointWithMeasurement("system").
		AddTag("location", "server1").
		AddField("cpu_usage", cpuUsage).
		AddField("memory_usage", memoryUsage).
		SetTime(time.Now())
	// // Write point to InfluxDB
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		panic(err)
	}
}
