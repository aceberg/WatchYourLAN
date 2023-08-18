package auth

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

// StartSession for new login
func StartSession(w http.ResponseWriter, r *http.Request) {

	sessionToken := uuid.NewString()

	allSessions[sessionToken] = time.Now().Add(60 * time.Second)

	setTokenCookie(w, sessionToken)

	http.Redirect(w, r, "/", 302)

}

// LogOut - log out
func LogOut(w http.ResponseWriter, r *http.Request) {

	sessionToken := getTokenFromCookie(r)

	delete(allSessions, sessionToken)

	setTokenCookie(w, "")

	http.Redirect(w, r, "/", 302)
}
