package auth

import "github.com/gorilla/sessions"

const (
	SessionName = "session"
)

type SessionOptions struct {
	CookiesKey string
	MaxAge     int
	HttpOnly   bool // Should be true if the site is served over HTTP (development environment)
	Secure     bool // Should be true if the site is served over HTTPS (production environment)
}

func NewCookieStore(opts SessionOptions) *sessions.CookieStore {
	store := sessions.NewCookieStore([]byte(opts.CookiesKey))

	store.MaxAge(opts.MaxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = opts.HttpOnly
	store.Options.Secure = opts.Secure

	return store
}
