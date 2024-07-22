package main

import (
	"fmt"
	"time"

	"github.com/abhinavgemawat/MetricsObservabilitySystem/api/api"
)

func main() {

	var metrics []api.Metric

	go api.ScrapeMetrics(&metrics)

	for {
		time.Sleep(10 * time.Second)
		for _, metric := range metrics {
			serializedMetric, err := metric.Serialize()
			if err != nil {
				fmt.Printf("Error serializing metric: %v\n", err)
				continue
			}

			message := api.KafkaMessage{
				Key:   metric.Name,
				Value: serializedMetric,
				Time:  metric.Timestamp,
			}

			err = api.ProduceMessage("metrics-topic", message)
			if err != nil {
				fmt.Printf("Error producing message: %v\n", err)
			}

			err = api.WriteMetricToInfluxDB(metric)
			if err != nil {
				fmt.Printf("Error writing metric to InfluxDB: %v\n", err)
			}
		}

		// Clear metrics slice after sending to Kafka
		metrics = nil
	}
}
