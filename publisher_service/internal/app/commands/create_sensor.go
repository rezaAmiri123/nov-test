package commands

import (
	"context"
	"github.com/opentracing/opentracing-go"
	messageClient "github.com/rezaAmiri123/nov-test/pkg/event"
	"github.com/rezaAmiri123/nov-test/pkg/logger"
	"github.com/rezaAmiri123/nov-test/publisher_service/internal/domain"
)

type CreateSensorHandler struct {
	logger          logger.Logger
	repo            domain.Repository
	messageProducer messageClient.Producer
}

func NewCreateSensorHandler(repo domain.Repository, logger logger.Logger, producer messageClient.Producer) *CreateSensorHandler {
	if repo == nil {
		panic("userRepo is nil")
	}
	return &CreateSensorHandler{repo: repo, logger: logger, messageProducer: producer}
}

func (h CreateSensorHandler) Handle(ctx context.Context, arg *[]domain.Sensor) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateSensorHandler.Handle")
	defer span.Finish()

	//// TODO we need to hash the password
	//hashedPassword, err := utils.HashPassword(arg.Password)
	//if err != nil {
	//	h.logger.Errorf("connot hash the password: %v", err)
	//	return &user.User{}, fmt.Errorf("connot hash the password: %w", err)
	//}
	//arg.Password = hashedPassword
	//
	//u, err := h.repo.CreateUser(ctx, arg)
	//if err != nil {
	//	h.logger.Errorf("connot create user: %v", err)
	//	return &user.User{}, fmt.Errorf("connot create user: %w", err)
	//}
	//
	//// remove the password from response
	//u.Password = ""
	//
	//err = h.sentEvent(ctx, u)
	//if err != nil {
	//	h.logger.Errorf("connot send create user event: %v", err)
	//}
	//
	return nil
}

//func (h CreateUserHandler) sentEvent(ctx context.Context, u *user.User) error {
//
//	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateUserHandler.sentEvent")
//	defer span.Finish()
//
//	req := &kafkaMessage.CreateUser{
//		UserID:   u.UserID.String(),
//		Username: u.Username,
//		Email:    u.Email,
//		Bio:      u.Bio,
//		Image:    u.Image,
//	}
//
//	message, err := proto.Marshal(req)
//	if err != nil {
//		return err
//	}
//	err = h.kafkaProducer.PublishMessage(ctx, kafka.Message{
//		Topic:   kafkaClient.CreateUserTopic,
//		Value:   message,
//		Time:    time.Now().UTC(),
//		Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
//	})
//	if err != nil {
//		h.logger.Errorf("can not send kafka message %v", err)
//	}
//	return err
//}
