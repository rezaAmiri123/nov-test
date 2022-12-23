package agent

import (
	"github.com/rezaAmiri123/nov-test/pkg/logger"
	"github.com/rezaAmiri123/nov-test/subscriber_service/internal/app"
	"github.com/rezaAmiri123/nov-test/subscriber_service/internal/domain"
	"io"
	"sync"
)

type Config struct {
	SecretKey string `mapstructure:"SECRET_KEY"`

	// NATS config
	NatsUrl          string `mapstructure:"NATS_URL"`
	NatsClusterID    string `mapstructure:"NATS_CLUSTER_ID"`
	NatsClientID     string `mapstructure:"NATS_CLIENT_ID"`
	NatsPingInterval int    `mapstructure:"NATS_PING_INTERVAL"`
	NatsPingMaxOut   int    `mapstructure:"NATS_PING_MAX_OUT"`

	// postgres.Config
	PGDriver   string `mapstructure:"POSTGRES_DRIVER"`
	PGHost     string `mapstructure:"POSTGRES_HOST"`
	PGPort     string `mapstructure:"POSTGRES_PORT"`
	PGUser     string `mapstructure:"POSTGRES_USER"`
	PGDBName   string `mapstructure:"POSTGRES_DB_NAME"`
	PGPassword string `mapstructure:"POSTGRES_PASSWORD"`

	// kafka config
	KafkaBrokers    []string `mapstructure:"KAFKA_BROKERS"`
	KafkaGroupID    string   `mapstructure:"KAFKA_GROUP_ID"`
	KafkaInitTopics bool     `mapstructure:"KAFKA_INIT_TOPICS"`

	// applogger.Config
	LogLevel   string `mapstructure:"LOG_LEVEL"`
	LogDevMode bool   `mapstructure:"LOG_DEV_MOD"`
	LogEncoder string `mapstructure:"LOG_ENCODER"`

	// tracing.Config
	TracerServiceName string `mapstructure:"TRACER_SERVICE_NAME"`
	TracerHostPort    string `mapstructure:"TRACER_HOST_PORT"`
	TracerEnable      bool   `mapstructure:"TRACER_ENABLE"`
	TracerLogSpans    bool   `mapstructure:"TRACER_LOG_SPANS"`

	// metrics.Config
	//MetricServiceName     string `mapstructure:"METRIC_SERVICE_NAME"`
	//MetricServiceHostPort string `mapstructure:"METRIC_SERVICE_HOST_PORT"`
}

type Agent struct {
	Config

	logger logger.Logger
	//metric *metrics.MessageServiceMetric
	// httpServer  *http.Server
	// grpcServer  *grpc.Server
	repository  domain.Repository
	Application *app.Application
	// AuthClient  auth.AuthClient

	shutdown     bool
	shutdowns    chan struct{}
	shutdownLock sync.Mutex
	closers      []io.Closer
}

func NewAgent(config Config) (*Agent, error) {
	a := &Agent{
		Config:    config,
		shutdowns: make(chan struct{}),
	}
	setupsFn := []func() error{
		a.setupLogger,
		//a.setupTracing,
		a.setupApplication,
		a.setupNats,
	}
	for _, fn := range setupsFn {
		if err := fn(); err != nil {
			return nil, err
		}
	}
	return a, nil
}

func (a *Agent) Shutdown() error {
	a.shutdownLock.Lock()
	defer a.shutdownLock.Unlock()

	if a.shutdown {
		return nil
	}
	a.shutdown = true
	close(a.shutdowns)
	shutdown := []func() error{
		//func() error {
		//	return a.httpServer.Shutdown(context.Background())
		//},
		// func() error {
		// 	a.grpcServer.GracefulStop()
		// 	return nil
		// },
		//func() error {
		//	return a.jaegerCloser.Close()
		//},
	}
	for _, fn := range shutdown {
		if err := fn(); err != nil {
			return err
		}
	}
	for _, closer := range a.closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}
