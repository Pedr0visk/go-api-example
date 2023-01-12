package rest

import (
	"analytics/internal/domain"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PageService interface {
	RetrievePageMetadata(ctx context.Context, url string) (domain.Page, error)
}

type PageHandler struct {
	svc PageService
}

func NewPageHandler(svc PageService) *PageHandler {
	return &PageHandler{
		svc: svc,
	}
}

func (p *PageHandler) Register(app *gin.Engine) {
	app.GET("/api/pages/:url", p.get)
}

type CreateTraceResponse struct {
	ID string `json:"id"`
}

func (p *PageHandler) get(c *gin.Context) {
	//ctx, cancel := context.WithTimeout(c, 5*time.Second)
	//defer cancel()

	fmt.Println(c.Param("url"))
	c.JSON(http.StatusOK, &struct{ Ok string }{Ok: "true"})
	//
	//page, err := p.svc.RetrievePageMetadata(c, "url")
	//if err != nil {
	//	renderErrorResponse(c, "get failed", err)
	//}
	//
	//c.JSON(http.StatusOK, &page)
}
