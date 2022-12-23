package natsimpl

import (
	"context"

	"github.com/nats-io/stan.go"
	"github.com/rezaAmiri123/nov-test/pkg/logger"
)

type Config struct {
	ClusterID    string
	Url          string
	ClientID     string
	PingInterval int
	PingMaxOut   int
}

func NewClientConn(ctx context.Context, logger logger.Logger, config Config) (stan.Conn, error) {

	//clusterID := "NATS"       // nats cluster id
	//url := "nats://nats:4222" // nats url
	//clientID := "800"
	// you can set client id anything
	sc, err := stan.Connect(config.ClusterID, config.ClientID, stan.NatsURL(config.Url),
		stan.Pings(config.PingInterval, config.PingMaxOut),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			logger.Errorf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		logger.Errorf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, config.Url)
		return nil, err
	}

	logger.Printf("Connected Nats")

	//Sc = sc
	return sc, nil
}
