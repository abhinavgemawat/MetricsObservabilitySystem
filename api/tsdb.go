package api

import (
	"context"
	"fmt"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var influxToken = os.Getenv("INFLUXDB_TOKEN")

const (
	bucket    = "metrics"
	org       = "COMP41720"
	influxURL = "http://localhost:8086"
)

func WriteMetricToInfluxDB(metric Metric) error {
	fmt.Println("InfluxDB token:", influxToken)
	client := influxdb2.NewClient(influxURL, influxToken)
	defer client.Close()

	writeAPI := client.WriteAPIBlocking(org, bucket)

	tags := map[string]string{}
	for k, v := range metric.Tags {
		tags[k] = v
	}

	p := influxdb2.NewPoint(metric.Name,
		tags,
		map[string]interface{}{"value": metric.Value},
		metric.Timestamp)

	if err := writeAPI.WritePoint(context.Background(), p); err != nil {
		return fmt.Errorf("failed to write metric to InfluxDB: %v", err)
	}

	return nil
}
