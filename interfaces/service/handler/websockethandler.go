package handler

import (
	"context"

	"github.com/gorilla/websocket"
	"github.com/howood/kangaroochat/domain/repository"
	log "github.com/howood/kangaroochat/infrastructure/logger"
	"github.com/howood/kangaroochat/infrastructure/requestid"
	"github.com/labstack/echo/v4"
)

// WebSocket 更新用
var upgrader = websocket.Upgrader{}

// WebSockerHandler struct
type WebSockerHandler struct {
	BaseHandler
	BroadCaster repository.BroadCasterRepository
}

// Request is chat request
func (wsh WebSockerHandler) Request(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	xRequestID := requestid.GetRequestID(c.Request())
	wsh.ctx = context.WithValue(context.Background(), echo.HeaderXRequestID, xRequestID)
	identifier := c.Param("identifier")
	log.Info(wsh.ctx, "========= START REQUEST : "+requesturi)
	log.Info(wsh.ctx, c.Request().Method)
	log.Info(wsh.ctx, c.Request().Header)
	log.Info(wsh.ctx, identifier)

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()
	wsh.BroadCaster.SetNewClient(ws, identifier)

	for {
		if err := wsh.BroadCaster.ReadMessage(ws, identifier); err != nil {
			log.Error(wsh.ctx, err)
			break
		}
	}
	return nil
}
