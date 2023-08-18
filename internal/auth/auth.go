package auth

import (
	"net/http"
	"time"
)

// Auth - main auth func
func Auth(next http.HandlerFunc, conf *Conf) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !conf.Auth || conf.User == "" || conf.Password == "" {
			next.ServeHTTP(w, r)
			return
		}

		sessionToken := getTokenFromCookie(r)

		userSession, exists := allSessions[sessionToken]
		if !exists {
			http.Redirect(w, r, "/login/", 302)
			return
		}
		if userSession.Before(time.Now()) {
			delete(allSessions, sessionToken)
			http.Redirect(w, r, "/login/", 302)
			return
		}

		userSession = time.Now().Add(conf.Expire)
		allSessions[sessionToken] = userSession

		next.ServeHTTP(w, r)
	}
}
