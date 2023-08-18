package auth

import (
	"net/http"
)

func setTokenCookie(w http.ResponseWriter, token string) {

	cookie := http.Cookie{Name: cookieName, Value: token, Path: "/"}
	http.SetCookie(w, &cookie)
}

func getTokenFromCookie(r *http.Request) string {

	cookie, err := r.Cookie(cookieName)
	if err != nil {

		return ""
	}

	return cookie.Value
}
