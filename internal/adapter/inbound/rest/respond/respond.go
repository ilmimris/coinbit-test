package respond

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const ErrMaxStack = 5

type (
	Causer interface {
		Cause() error
	}

	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	Response struct {
		Data       interface{} `json:"data,omitempty"`
		Message    string      `json:"message"`
		Code       string      `json:"code"`
		StatusCode int         `json:"statusCode"`
		Error      string      `json:"error,omitempty"`
		Latency    string      `json:"latency"`
	}
)

func (err *Error) Error() string {
	return fmt.Sprintf("error with code: %d; message: %s", err.Code, err.Message)
}

func Success(c *fiber.Ctx, status int, content interface{}) error {
	requestID := c.GetRespHeader(fiber.HeaderXRequestID)
	c.Set("X-Request-ID", requestID)
	c.Status(status)

	return c.JSON(&Response{
		Data:       content,
		Message:    "success",
		Code:       http.StatusText(status),
		StatusCode: status,
	})
}

func Fail(c *fiber.Ctx, status, errorCode int, err error) error {
	var (
		message = err.Error()
	)

	// if error masked, get detail!
	if ec, ok := err.(Causer); ok {
		err = ec.Cause()
	}

	requestID := c.GetRespHeader(fiber.HeaderXRequestID)
	c.Set("X-Request-ID", requestID)
	c.Status(status)

	return c.JSON(&Response{
		Message:    message,
		Code:       http.StatusText(errorCode),
		StatusCode: status,
		Error:      err.Error(),
	})
}
