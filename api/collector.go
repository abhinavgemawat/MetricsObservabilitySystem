package api

import (
	"fmt"
)

func WriteMetrics(cpuUsage, memoryUsage float64) {
	// Create a point using the fluent API
	// p := influxdb2.NewPointWithMeasurement("system").
	// 	AddTag("location", "server1").
	// 	AddField("cpu_usage", cpuUsage).
	// 	AddField("memory_usage", memoryUsage).
	// 	SetTime(time.Now())
	fmt.Print("Hello World")
	// Write point to InfluxDB
	// err := timeseriesdb.WriteAPI.WritePoint(context.Background(), p)
	// if err != nil {
	// 	log.Printf("Error writing to InfluxDB: %v", err)
	// } else {
	// 	log.Printf("Written data to InfluxDB: cpu_usage=%f, memory_usage=%f", cpuUsage, memoryUsage)
	// }
}
