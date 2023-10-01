package actor

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/howood/kangaroochat/domain/entity"
	"github.com/howood/kangaroochat/domain/repository"
	"github.com/howood/kangaroochat/library/utils"
)

// CookieExpired is cookie's expired
var CookieExpired = utils.GetOsEnv("COOKIE_EXPIED", "3600")

// RoomTokenKey is cookie's key name of RoomToken
const RoomTokenKey = "kc_room_token"

// CookieOperator struct
type CookieOperator struct {
	repository.CookieRepository
}

// NewCookieOperator creates a new CookieRepository
func NewCookieOperator(ctx context.Context, r *http.Request) *CookieOperator {
	expired, _ := strconv.ParseInt(CookieExpired, 10, 64)
	return &CookieOperator{
		&cookieJar{
			cookie: &entity.Cookie{
				Cookie: new(http.Cookie),
			},
			expired: time.Duration(expired),
			ctx:     ctx,
		},
	}
}

// cookieJar struct
type cookieJar struct {
	cookie  *entity.Cookie
	expired time.Duration
	ctx     context.Context
}

// Set sets to Cookie
func (co *cookieJar) Set(key, value string) {
	co.cookie.Cookie.Name = key
	co.cookie.Cookie.Value = value
	co.cookie.Cookie.Expires = time.Now().Add(co.expired * time.Second)
	co.cookie.Cookie.Path = "/"
}

// GetCookie get cookie struct
func (co *cookieJar) GetCookie() *http.Cookie {
	return co.cookie.Cookie
}
