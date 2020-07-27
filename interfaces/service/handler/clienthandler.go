package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/howood/kangaroochat/application/actor"
	"github.com/howood/kangaroochat/application/actor/cacheservice"
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

	claims := ch.getClaimsFromToken(c)
	viewval := map[string]interface{}{
		"identifier": identifier,
		"username":   claims.Name,
	}
	if roomname, err := ch.getRoomname(identifier); err != nil {
		log.Error(ch.ctx, err)
	} else {
		viewval["roomname"] = roomname
	}
	return c.Render(http.StatusOK, "client.html", viewval)
}

func (ch ClientHandler) getRoomname(identifier string) (string, error) {
	cacheAssessor := cacheservice.NewCacheAssessor(ch.ctx)
	if cachedvalue, cachedfound := cacheAssessor.Get(identifier); cachedfound {
		chatroom := actor.NewChatRoomOperator(ch.ctx)
		var err error
		switch xi := cachedvalue.(type) {
		case []byte:
			err = chatroom.GobDecode(xi)
		case string:
			err = chatroom.GobDecode([]byte(xi))
		default:
			err = errors.New("get cache error")
		}
		if err != nil {
			log.Error(ch.ctx, err.Error())
			return "", err
		}
		return chatroom.GetRoomName(), nil
	}

	return "", errors.New("No Room")
}
