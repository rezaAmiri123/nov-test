package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/nov-test/publisher_service/internal/domain"
	"net/http"
	"time"
)

func (h *HttpServer) CreateSensor() echo.HandlerFunc {
	return func(c echo.Context) error {
		//h.metrics.CreateUserHttpRequests.Inc()

		span, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "HttpServer.CreateUser")
		defer span.Finish()

		sensors := []Sensor{}
		if err := c.Bind(&sensors); err != nil {
			h.log.WarnMsg("Bind", err)
			h.traceErr(span, err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		//if err := h.validate.StructCtx(ctx, sensors); err != nil {
		//	h.log.WarnMsg("validate", err)
		//	h.traceErr(span, err)
		//	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		//}

		var req []domain.Sensor
		for _, val := range sensors {
			req = append(req, val.getDomainSensor())
		}
		err := h.app.Commands.CreateSensor.Handle(ctx, req)
		if err != nil {
			h.log.WarnMsg("CreateSensor", err)
			h.traceErr(span, err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		//h.metrics.SuccessHttpRequests.Inc()
		return c.JSON(http.StatusOK, nil)
	}
}

type Sensor struct {
	Name      string  `json:"name" db:"name" validate:"required"`
	Timestamp int64   `json:"timestamp" db:"timestamp" validate:"required"`
	Value     float64 `json:"value" db:"value" validate:"required"`
}

func (s Sensor) getDomainSensor() domain.Sensor {
	return domain.Sensor{
		Name:      s.Name,
		Timestamp: time.Unix(s.Timestamp, 0),
		Value:     s.Value,
	}
}
