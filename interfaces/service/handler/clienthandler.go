package handler

import (
	"context"
	"net/http"

	log "github.com/howood/kangaroochat/infrastructure/logger"
	"github.com/howood/kangaroochat/infrastructure/requestid"
	"github.com/labstack/echo/v4"
)

// ClientHandler struct
type ClientHandler struct {
	BaseHandler
}

// Request is chat request
func (ch ClientHandler) Request(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	xRequestID := requestid.GetRequestID(c.Request())
	ch.ctx = context.WithValue(context.Background(), echo.HeaderXRequestID, xRequestID)
	identifier := c.Param("identifier")
	log.Info(ch.ctx, "========= START REQUEST : "+requesturi)
	log.Info(ch.ctx, c.Request().Method)
	log.Info(ch.ctx, c.Request().Header)
	log.Info(ch.ctx, identifier)
	viewval := map[string]interface{}{
		"identifier": identifier,
	}
	return c.Render(http.StatusOK, "client.html", viewval)
}
