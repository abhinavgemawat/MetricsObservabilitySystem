// metrics_test.go

package api

import (
	"testing"
	"time"

	"github.com/abhinavgemawat/MetricsObservabilitySystem/api/api"
	"github.com/stretchr/testify/assert"
)

func TestSerializeSuccess(t *testing.T) {
	metric := api.Metric{
		Name:      "cpu_usage",
		Value:     0.85,
		Timestamp: time.Now(),
		Tags:      map[string]string{"host": "localhost"},
	}

	jsonString, err := metric.Serialize()
	assert.NoError(t, err)
	assert.Contains(t, jsonString, `"name":"cpu_usage"`)
	assert.Contains(t, jsonString, `"value":0.85`)
	assert.Contains(t, jsonString, `"host":"localhost"`)
}

func TestSerializeFailure(t *testing.T) {
	metric := api.Metric{
		Name:      "cpu_usage",
		Value:     0.85,
		Timestamp: time.Now(),
		Tags:      map[string]string{"host": string([]byte{0xff})}, // Invalid UTF-8 sequence
	}

	_, nil := metric.Serialize()
	assert.Error(t, nil)
}
