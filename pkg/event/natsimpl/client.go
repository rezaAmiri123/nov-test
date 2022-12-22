package natsimpl

import (
	"context"
	"github.com/nats-io/stan.go"
	"github.com/rezaAmiri123/nov-test/pkg/logger"
)

//// NewKafkaConn create new kafka connection
//func NewKafkaConn(ctx context.Context, kafkaCfg Config) (*kafka.Conn, error) {
//	return kafka.DialContext(ctx, "tcp", kafkaCfg.Brokers[0])
//}

func NewClientConn(ctx context.Context, logger logger.Logger) (stan.Conn, error) {

	clusterID := "test-cluster"    // nats cluster id
	url := "nats://127.0.0.1:4222" // nats url
	clientID := "123"
	// you can set client id anything
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(url),
		stan.Pings(1, 3),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			logger.Errorf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		logger.Errorf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, url)
	}

	logger.Printf("Connected Nats")

	//Sc = sc
	return sc, nil
}
