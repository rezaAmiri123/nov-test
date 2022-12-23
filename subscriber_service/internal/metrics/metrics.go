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

type SubscriberServiceMetric struct {
	CreateSensorAsyncRequests prometheus.Counter
	SuccessAsyncRequests      prometheus.Counter
	ErrorAsyncRequests        prometheus.Counter
}

func NewPublisherServiceMetric(cfg *Config) *SubscriberServiceMetric {
	return &SubscriberServiceMetric{
		CreateSensorAsyncRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_sensor_async_requests_total", cfg.MetricServiceName),
			Help: "The total of create sensor async requests",
		}),

		SuccessAsyncRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_async_requsts_total", cfg.MetricServiceName),
			Help: "The total number of success async requests",
		}),
		ErrorAsyncRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_error_async_requsts_total", cfg.MetricServiceName),
			Help: "The total number of error async requests",
		}),
	}
}
