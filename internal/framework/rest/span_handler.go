package rest

import (
	"analytics/internal"
	"analytics/internal/application/service"
	"context"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const uuidRegEx string = `[0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12}`

// SpanService ...
type SpanService interface {
	Create(ctx context.Context, params service.SpanCreateParams) error
}

// SpanHandler ...
type SpanHandler struct {
	svc SpanService
}

// NewSpanHandler ...
func NewSpanHandler(svc SpanService) *SpanHandler {
	return &SpanHandler{
		svc: svc,
	}
}

func (s *SpanHandler) Register(r *gin.Engine) {
	r.POST("/api/spans/create", s.create)
}

// CreateSpanRequest defines the request used for creating spans.
type CreateSpanRequest struct {
	SessionID string `json:"session_id"`
	PageID    string `json:"page_id"`
	UserAgent string `json:"user_agent"`
	Date      int    `json:"date"`
	Hostname  string `json:"hostname"`
	Pathname  string `json:"pathname"`
	Search    string `json:"search"`
}

func (c CreateSpanRequest) Validate() error {
	if err := validation.ValidateStruct(&c,
		validation.Field(&c.SessionID, validation.Required, validation.Match(regexp.MustCompile(uuidRegEx))), validation.Field(&c.PageID, validation.Required, validation.Match(regexp.MustCompile(uuidRegEx)))); err != nil {
		return err
	}

	return nil
}

func (s *SpanHandler) create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	var req CreateSpanRequest

	// Call BindJSON to bind the received JSON to req
	if err := c.BindJSON(&req); err != nil {
		renderErrorResponse(c, "invalid request",
			internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "json decoder"))

		return
	}

	// Call req.Validate to validate received req body
	if err := req.Validate(); err != nil {
		renderErrorResponse(c, "invalid request",
			internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "req.Validate"))

		return
	}

	err := s.svc.Create(ctx, service.SpanCreateParams{
		PageID:    req.PageID,
		SessionID: req.SessionID,
		Hostname:  req.Hostname,
		Pathname:  req.Pathname,
		Search:    req.Search,
		Date:      req.Date,
		UserAgent: req.UserAgent,
	})
	if err != nil {
		renderErrorResponse(c, "create failed", err)
		return
	}

	renderResponse(c, nil, http.StatusAccepted)
}
