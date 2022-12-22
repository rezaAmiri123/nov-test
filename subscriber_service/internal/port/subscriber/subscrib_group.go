package subscriber

import (
	"github.com/nats-io/stan.go"
	"github.com/rezaAmiri123/nov-test/pkg/logger"
	"github.com/rezaAmiri123/nov-test/subscriber_service/internal/app"
)

type messageProcessor struct {
	Log logger.Logger
	App *app.Application
	Sc  stan.Conn
}
type Worker func(msg *stan.Msg)

func NewMessageProcessor(log logger.Logger, app *app.Application, sc stan.Conn) *messageProcessor {
	return &messageProcessor{Log: log, App: app, Sc: sc}
}

func (mp *messageProcessor) ProcessMessage(subject, qgroup, durable string) stan.Subscription {
	sub, err := mp.Sc.QueueSubscribe(subject,
		qgroup, mp.CreateSensor(),
		stan.DeliverAllAvailable(),
		stan.SetManualAckMode(),
		stan.DurableName(durable))
	if err != nil {
		mp.Log.Printf(err.Error())
	}
	return sub
}
