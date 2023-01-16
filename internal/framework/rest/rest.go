package rest

import (
	"analytics/internal"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// const otelName = "analytics/internal/rest"

// ErrorResponse represents a response containing an error message.
type ErrorResponse struct {
	Error       string            `json:"error"`
	Validations validation.Errors `json:"validations,omitempty"`
}

func renderErrorResponse(c *gin.Context, msg string, err error) {
	resp := ErrorResponse{Error: msg}
	status := http.StatusInternalServerError

	var ierr *internal.Error
	if !errors.As(err, &ierr) {
		resp.Error = "internal error"
	} else {
		switch ierr.Code() {
		case internal.ErrorCodeNotFound:
			status = http.StatusNotFound
		case internal.ErrorCodeInvalidArgument:
			status = http.StatusBadRequest

			var verrors validation.Errors
			if errors.As(ierr, &verrors) {
				resp.Validations = verrors
			}
		case internal.ErrorCodeUnknown:
			fallthrough
		default:
			status = http.StatusInternalServerError
		}
	}

	if err != nil {
		//_, span := otel.Tracer(otelName).Start(c, "renderErrorResponse")
		//defer span.End()
		//
		//span.RecordError(err)
		fmt.Println("otel called...")
	}

	fmt.Printf("Error: %v\n", err)

	c.JSON(status, resp)
}

func renderResponse(c *gin.Context, res interface{}, status int) {
	c.JSON(status, res)
}
