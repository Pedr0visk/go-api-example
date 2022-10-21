package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	"hive-data-collector/internal/application/dto"
	"hive-data-collector/internal/domain"
	"hive-data-collector/internal/framework/rest/http_err"
	"net/http"
	"time"
)

type TraceService interface {
	Create(ctx context.Context, params dto.TraceCreateParams) (domain.Trace, error)
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
	app.POST("/api/trace/create", t.create)
}

type CreateTraceRequest struct {
	UserWalletAddress string `json:"user_wallet_address"`
	PublisherUrl      string `json:"publisher_url"`
	Payload           string `json:"payload"`
	Date              int64  `json:"date"`
}

func (t *TraceHandler) create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	var request CreateTraceRequest

	defer cancel()

	if err := c.BindJSON(&request); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		return
	}

	_, err := t.svc.Create(ctx, dto.TraceCreateParams{
		UserWalletAddress: request.UserWalletAddress,
		Payload:           request.Payload,
		Date:              request.Date,
		PublisherUrl:      request.PublisherUrl,
	})

	if err != nil {
		http_err.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, nil)

}
