package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	"hive-data-collector/internal/application/dto"
	"hive-data-collector/internal/domain"
	"net/http"
)

type TraceService interface {
	CreateNewTrace(ctx context.Context, params dto.TraceCreateRequestBody) (domain.Trace, error)
}

type TraceHandler struct {
	svc TraceService
}

func NewTraceHandler(svc TraceService) *TraceHandler {
	return &TraceHandler{
		svc: svc,
	}
}

func (t *TraceHandler) Register(app *gin.Engine) {
	app.POST("/api/traces/create", t.create)
}

type CreateTraceRequest struct {
	UserWalletAddress string `json:"user_wallet_address" binding:"required"`
	PublisherUrl      string `json:"publisher_url" binding:"required"`
	Payload           string `json:"payload" binding:"required"`
	Date              int64  `json:"date" binding:"required"`
}

type CreateTraceResponse struct {
	ID string `json:"id"`
}

func (t *TraceHandler) create(c *gin.Context) {
	var req CreateTraceRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		renderErrorResponse(c, "invalid request", domain.WrapErrorf(err, domain.ErrorCodeInvalidArgument, "json decoder"))
		return
	}

	trace, err := t.svc.CreateNewTrace(c, dto.TraceCreateRequestBody{
		UserWalletAddress: req.UserWalletAddress,
		Payload:           req.Payload,
		Date:              req.Date,
		PublisherUrl:      req.PublisherUrl,
	})

	if err != nil {
		renderErrorResponse(c, "create failed", err)
	}

	c.JSON(http.StatusCreated, &CreateTraceResponse{ID: trace.ID})

}
