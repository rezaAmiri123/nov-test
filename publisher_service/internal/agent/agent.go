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

	//GRPCUserClientAddr string `mapstructure:"GRPC_USER_CLIENT_ADDR"`
	//GRPCUserClientPort int    `mapstructure:"GRPC_USER_CLIENT_PORT"`
	//
	//GRPCUserClientTLSEnabled       bool   `mapstructure:"GRPC_USER_CLIENT_TLS_ENABLED"`
	//GRPCUserClientTLSCertFile      string `mapstructure:"GRPC_USER_CLIENT_TLS_CERT_FILE"`
	//GRPCUserClientTLSKeyFile       string `mapstructure:"GRPC_USER_CLIENT_TLS_KEY_FILE"`
	//GRPCUserClientTLSCAFile        string `mapstructure:"GRPC_USER_CLIENT_TLS_CA_FILE"`
	//GRPCUserClientTLSServerAddress string `mapstructure:"GRPC_USER_CLIENT_TLS_SERVER_ADDRESS"`
	//
	//GRPCFinanceClientAddr string `mapstructure:"GRPC_FINANCE_CLIENT_ADDR"`
	//GRPCFinanceClientPort int    `mapstructure:"GRPC_FINANCE_CLIENT_PORT"`

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

	// metrics.Config
	//MetricServiceName     string `mapstructure:"METRIC_SERVICE_NAME"`
	//MetricServiceHostPort string `mapstructure:"METRIC_SERVICE_HOST_PORT"`
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
