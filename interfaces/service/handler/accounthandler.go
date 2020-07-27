package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/howood/kangaroochat/application/actor"
	"github.com/howood/kangaroochat/application/actor/cacheservice"
	"github.com/howood/kangaroochat/domain/entity"
	log "github.com/howood/kangaroochat/infrastructure/logger"
	"github.com/howood/kangaroochat/infrastructure/requestid"
	"github.com/labstack/echo/v4"
)

// AccountHandler struct
type AccountHandler struct {
	BaseHandler
}

// CreateGet is shown to create chatroom
func (ah AccountHandler) CreateGet(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	xRequestID := requestid.GetRequestID(c.Request())
	ah.ctx = context.WithValue(context.Background(), echo.HeaderXRequestID, xRequestID)
	log.Info(ah.ctx, "========= START REQUEST : "+requesturi)
	log.Info(ah.ctx, c.Request().Method)
	log.Info(ah.ctx, c.Request().Header)
	viewval := map[string]interface{}{
		"csrftoken": c.Get("csrf").(string),
	}
	return c.Render(http.StatusOK, "create.html", viewval)
}

// Create is request to create chatroom
func (ah AccountHandler) Create(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	xRequestID := requestid.GetRequestID(c.Request())
	ah.ctx = context.WithValue(context.Background(), echo.HeaderXRequestID, xRequestID)
	log.Info(ah.ctx, "========= START REQUEST : "+requesturi)
	log.Info(ah.ctx, c.Request().Method)
	log.Info(ah.ctx, c.Request().Header)
	form := entity.CreateRoomForm{}
	var err error
	if err == nil {
		err = c.Bind(&form)
	}
	if err == nil {
		err = ah.validate(form)
	}
	if err != nil {
		return ah.errorResponse(c, http.StatusBadRequest, "create.html", err)
	}
	var identifier string
	if identifier, err = ah.setRoom(form.RoomName, form.Password); err != nil {
		return ah.errorResponse(c, http.StatusBadRequest, "create.html", err)
	}
	redirecturl := "/login/" + identifier
	return c.Redirect(http.StatusSeeOther, redirecturl)
}

// LoginGet is shown to login
func (ah AccountHandler) LoginGet(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	xRequestID := requestid.GetRequestID(c.Request())
	ah.ctx = context.WithValue(context.Background(), echo.HeaderXRequestID, xRequestID)
	identifier := c.Param("identifier")
	log.Info(ah.ctx, "========= START REQUEST : "+requesturi)
	log.Info(ah.ctx, c.Request().Method)
	log.Info(ah.ctx, c.Request().Header)
	log.Info(ah.ctx, identifier)
	viewval := map[string]interface{}{
		"csrftoken": c.Get("csrf").(string),
	}
	return c.Render(http.StatusOK, "login.html", viewval)
}

// Login is request to login
func (ah AccountHandler) Login(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	xRequestID := requestid.GetRequestID(c.Request())
	ah.ctx = context.WithValue(context.Background(), echo.HeaderXRequestID, xRequestID)
	identifier := c.Param("identifier")
	log.Info(ah.ctx, "========= START REQUEST : "+requesturi)
	log.Info(ah.ctx, c.Request().Method)
	log.Info(ah.ctx, c.Request().Header)
	log.Info(ah.ctx, identifier)
	form := entity.LoginRoomForm{}
	var token string
	var err error
	if err == nil {
		err = c.Bind(&form)
	}
	if err == nil {
		err = ah.validate(form)
	}
	if err == nil {
		err = ah.loginRoom(identifier, form.Password)
	}
	if err == nil {
		token, err = ah.createToken(identifier, form.UserName)
	}
	if err != nil {
		return ah.errorResponse(c, http.StatusBadRequest, "login.html", err)
	}

	cookieop := actor.NewCookieOperator(ah.ctx, c.Request())
	cookieop.Set(actor.RoomTokenKey, token)
	c.SetCookie(cookieop.GetCookie())

	redirecturl := "/client/" + identifier
	return c.Redirect(http.StatusSeeOther, redirecturl)
}

func (ah AccountHandler) loginRoom(identifier, password string) error {
	cacheAssessor := cacheservice.NewCacheAssessor(ah.ctx)
	if cachedvalue, cachedfound := cacheAssessor.Get(identifier); cachedfound {
		chatroom := actor.NewChatRoomOperator(ah.ctx)
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
			log.Error(ah.ctx, err.Error())
			return err
		}
		if err := chatroom.ComparePassword(password); err != nil {
			return err
		}
		return nil
	}

	return errors.New("No Room")
}

func (ah AccountHandler) setRoom(roomname, password string) (string, error) {
	chatroom := actor.NewChatRoomOperator(ah.ctx)
	chatroom.Set(roomname, password)
	encodedcached, err := chatroom.GobEncode()
	if err != nil {
		return "", err

	}

	cacheAssessor := cacheservice.NewCacheAssessor(ah.ctx)
	cacheAssessor.Set(chatroom.GetIdentifier(), encodedcached, cacheservice.GetChacheExpired())
	return chatroom.GetIdentifier(), nil
}

func (ah AccountHandler) createToken(identifier, username string) (string, error) {
	jwtinstance := actor.NewJwtOperator(ah.ctx, username, false, identifier)
	tokenstr := jwtinstance.CreateToken(actor.TokenSecret)
	return tokenstr, nil
}

func (ah AccountHandler) errorResponse(c echo.Context, statuscode int, html string, err error) error {
	viewval := map[string]interface{}{
		"csrftoken": c.Get("csrf").(string),
		"errormsg":  err.Error(),
	}
	return c.Render(statuscode, html, viewval)

}
