package auth

import (
	"time"
)

// Conf - auth config
type Conf struct {
	Auth     bool
	User     string
	Password string
	ExpStr   string
	Expire   time.Duration
}

var allSessions = map[string]time.Time{}

var cookieName = "wyl_session_token"
