package handler

import (
	"context"
	"net/http"
	"time"

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

func (bh BaseHandler) setNewLatsModified() string {
	return time.Now().UTC().Format(http.TimeFormat)
}
