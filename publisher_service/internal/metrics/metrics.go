package metrics

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Config struct {
	MetricServiceName     string `mapstructure:"METRIC_SERVICE_NAME"`
	MetricServiceHostPort string `mapstructure:"METRIC_SERVICE_HOST_PORT"`
}

type PublisherServiceMetric struct {
	CreateSensorHttpRequests prometheus.Counter
	SuccessHttpRequests      prometheus.Counter
	ErrorHttpRequests        prometheus.Counter
}

func NewPublisherServiceMetric(cfg *Config) *PublisherServiceMetric {
	return &PublisherServiceMetric{
		CreateSensorHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_sensor_http_requests_total", cfg.MetricServiceName),
			Help: "The total of create sensor http requests",
		}),

		SuccessHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_http_requsts_total", cfg.MetricServiceName),
			Help: "The total number of success http requests",
		}),
		ErrorHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_error_http_requsts_total", cfg.MetricServiceName),
			Help: "The total number of error http requests",
		}),
	}
}
