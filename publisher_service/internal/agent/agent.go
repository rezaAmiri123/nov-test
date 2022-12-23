package agent

import (
	"context"
	"io"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/rezaAmiri123/nov-test/publisher_service/internal/app"
	// "github.com/rezaAmiri123/test-microservice/pkg/auth"

	"github.com/rezaAmiri123/nov-test/pkg/logger"
)

type Config struct {
	Debug     bool   `mapstructure:"DEBUG"`
	SecretKey string `mapstructure:"SECRET_KEY"`

	// NATS config
	NatsUrl          string `mapstructure:"NATS_URL"`
	NatsClusterID    string `mapstructure:"NATS_CLUSTER_ID"`
	NatsClientID     string `mapstructure:"NATS_CLIENT_ID"`
	NatsPingInterval int    `mapstructure:"NATS_PING_INTERVAL"`
	NatsPingMaxOut   int    `mapstructure:"NATS_PING_MAX_OUT"`

	// Http server address
	HttpServerAddr string `mapstructure:"HTTP_SERVER_ADDR"`
	HttpServerPort int    `mapstructure:"HTTP_SERVER_PORT"`

	// check alive
	HttpKeepAliveServerHostPort string `mapstructure:"HTTP_KEEP_ALIVE_SERVER_HOST_PORT"`

	// applogger.Config
	LogLevel   string `mapstructure:"LOG_LEVEL"`
	LogDevMode bool   `mapstructure:"LOG_DEV_MOD"`
	LogEncoder string `mapstructure:"LOG_ENCODER"`

	// tracing.Config
	TracerServiceName string `mapstructure:"TRACER_SERVICE_NAME"`
	TracerHostPort    string `mapstructure:"TRACER_HOST_PORT"`
	TracerEnable      bool   `mapstructure:"TRACER_ENABLE"`
	TracerLogSpans    bool   `mapstructure:"TRACER_LOG_SPANS"`

	//rabbitmq
	RabbitMQUser           string `mapstructure:"RABBITMQ_USER"`
	RabbitMQPassword       string `mapstructure:"RABBITMQ_PASSWORD"`
	RabbitMQHost           string `mapstructure:"RABBITMQ_HOST"`
	RabbitMQPort           int    `mapstructure:"RABBITMQ_PORT"`
	RabbitMQExchange       string `mapstructure:"RABBITMQ_EXCHANGE"`
	RabbitMQQueue          string `mapstructure:"RABBITMQ_QUEUE"`
	RabbitMQRoutingKey     string `mapstructure:"RABBITMQ_ROUTING_KEY"`
	RabbitMQConsumerTag    string `mapstructure:"RABBITMQ_CONSUMER_TAG"`
	RabbitMQWorkerPoolSize int    `mapstructure:"RABBITMQ_WORKER_POOL_SIZE"`
}

type Agent struct {
	Config

	// GRPCUserClientTLSConfig    *tls.Config
	//GRPCFinanceClientTLSConfig *tls.Config

	logger logger.Logger
	//metric     *metrics.ApiServiceMetric
	httpServer *echo.Echo
	//grpcServer *grpc.Server
	// repository  api.Repository
	Application *app.Application
	//Maker       token.Maker
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
		//a.setupMetric,
		//a.setupRepository,
		a.setupTracing,
		a.setupApplication,
		//a.setupAuthClient,
		a.setupKeepAlive,
		a.setupHttpServer,
		// a.setupGrpcServer,
		//a.setupGRPCServer,
		//a.setupTracer,
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
		func() error {
			return a.httpServer.Shutdown(context.Background())
		},
		//func() error {
		//	a.grpcServer.GracefulStop()
		//	return nil
		//},
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
