package custom

import (
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	EchoRequest interface {
		Bind(obj any) error
	}

	customEchoRequest struct {
		ctx       echo.Context
		validator *validator.Validate
	}
)

var (
	once              sync.Once
	validatorInstance *validator.Validate
)

func (c *customEchoRequest) Bind(obj any) error {
	if err := c.ctx.Bind(obj); err != nil {
		return err
	}
	if err := c.validator.Struct(obj); err != nil {
		return err
	}
	return nil
}

func NewCustomEchoRequest(ctx echo.Context) EchoRequest {
	once.Do(func() {
		validatorInstance = validator.New()
	})

	return &customEchoRequest{
		ctx:       ctx,
		validator: validatorInstance,
	}
}
