package actor

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/howood/kangaroochat/domain/entity"
	"github.com/howood/kangaroochat/domain/repository"
)

// SessionExpired is token's expired
var SessionExpired = os.Getenv("SESSION_EXPIED")

const RoomTokenKey = "room_token"

var sessionstore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

// JwtOperator struct
type SessionOperator struct {
	cookie  *entity.Cookie
	expired time.Duration
	ctx     context.Context
}

// NewSessionOperator creates a new JwtClaimsRepository
func NewCookieOperator(ctx context.Context, r *http.Request) (repository.CookieRepository, error) {
	expired, err := strconv.Atoi(SessionExpired)
	if err != nil {
		return nil, err
	}
	return &SessionOperator{
		cookie: &entity.Cookie{
			Cookie: new(http.Cookie),
		},
		expired: time.Duration(int64(expired)),
		ctx:     ctx,
	}, nil
}

// Set creates a new token
func (so *SessionOperator) Set(key, value string) {
	so.cookie.Cookie.Name = key
	so.cookie.Cookie.Value = value
	so.cookie.Cookie.Expires = time.Now().Add(so.expired * time.Minute)
	so.cookie.Cookie.Path = "/"
}

func (so *SessionOperator) GetCookie() *http.Cookie {
	return so.cookie.Cookie
}
