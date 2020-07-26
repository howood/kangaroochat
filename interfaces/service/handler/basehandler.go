package handler

import (
	"context"

	"github.com/dgrijalva/jwt-go"
	"github.com/howood/kangaroochat/application/actor"
	"github.com/howood/kangaroochat/domain/entity"
	"github.com/labstack/echo/v4"
)

// BaseHandler struct
type BaseHandler struct {
	ctx context.Context
}

func (bh BaseHandler) errorResponse(c echo.Context, statudcode int, err error) error {
	c.Response().Header().Set(echo.HeaderXRequestID, bh.ctx.Value(echo.HeaderXRequestID).(string))
	return c.JSONPretty(statudcode, map[string]interface{}{"message": err.Error()}, "    ")
}

func (bh BaseHandler) setResponseHeader(c echo.Context, lastmodified, contentlength, xrequestud string) {
	c.Response().Header().Set(echo.HeaderLastModified, lastmodified)
	c.Response().Header().Set(echo.HeaderContentLength, contentlength)
	c.Response().Header().Set(echo.HeaderXRequestID, xrequestud)
}

func (bh BaseHandler) getClaimsFromToken(c echo.Context) *entity.JwtClaims {
	user := c.Get(actor.JWTContextKey).(*jwt.Token)
	claims := user.Claims.(*entity.JwtClaims)
	return claims
}
