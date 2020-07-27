package repository

import "net/http"

// CookieRepository interface
type CookieRepository interface {
	Set(key, value string)
	GetCookie() *http.Cookie
}
