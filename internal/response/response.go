package response

import (
	"github.com/gin-gonic/gin"
	"github.com/vincentweilasto16/wallet-api/internal/errors"
)

// Meta contains status code and message
type Meta struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// ErrorDetail represents a single error
type ErrorDetail struct {
	Code    string `json:"code"`            // internal code
	Field   string `json:"field,omitempty"` // optional, e.g., validation
	Message string `json:"message"`         // human-readable
}

// Response is a generic API response
type Response[T any] struct {
	Data   T             `json:"data,omitempty"`
	Errors []ErrorDetail `json:"errors,omitempty"`
	Meta   Meta          `json:"meta"`
}

// Success sends a standardized success response
func Success[T any](c *gin.Context, data T, message ...string) {
	msg := "Success"
	if len(message) > 0 {
		msg = message[0]
	}

	res := Response[T]{
		Data:   data,
		Errors: nil,
		Meta: Meta{
			Status:  200,
			Message: msg,
		},
	}
	c.JSON(200, res)
}

func Error(c *gin.Context, err error) {
	var httpCode int
	var code, message, field string

	if appErr, ok := err.(errors.AppError); ok {
		httpCode = appErr.HTTPCode
		code = appErr.Code
		message = appErr.Message
	} else {
		httpCode = 500
		code = "INTERNAL_ERROR"
		message = err.Error()
	}

	res := struct {
		Data   interface{}   `json:"data"`
		Errors []ErrorDetail `json:"errors"`
		Meta   Meta          `json:"meta"`
	}{
		Data: nil,
		Errors: []ErrorDetail{
			{Code: code, Message: message, Field: field},
		},
		Meta: Meta{
			Status:  httpCode,
			Message: message,
		},
	}

	c.JSON(httpCode, res)
}
