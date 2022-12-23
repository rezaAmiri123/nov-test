package subscriber

import (
	"context"
	"github.com/nats-io/stan.go"
	eventMessages "github.com/rezaAmiri123/nov-test/publisher_service/proto/event"
	"github.com/rezaAmiri123/nov-test/subscriber_service/internal/domain"
	"google.golang.org/protobuf/proto"
)

func (mp *messageProcessor) CreateSensor() stan.MsgHandler {
	return func(msg *stan.Msg) {
		if err := msg.Ack(); err != nil {
			mp.Log.Printf("failed to ACK msg:%v", err)
		}
		var m eventMessages.CreateSensor
		if err := proto.Unmarshal(msg.Data, &m); err != nil {
			//s.log.WarnMsg("proto.Unmarshal", err)
			//s.commitErrMessage(ctx, r, m)
			return
		}

		var sensors []*domain.Sensor
		for _, val := range m.Sensors {
			sensors = append(sensors, &domain.Sensor{
				Name:      val.GetName(),
				Timestamp: val.GetTimestamp().AsTime(),
				Value:     val.GetValue(),
			})
		}

		err := mp.App.Commands.CreateSensor.Handle(context.Background(), sensors)
		if err != nil {
			mp.Log.Errorf("error create email consumer", err)
		}
		//s.commitMessage(ctx, r, m)
	}

}
