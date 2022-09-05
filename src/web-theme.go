package main

import (
  "net/http"
  "html"
  "strings"
)

func theme(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		urlString := html.EscapeString(r.URL.Path)
		tags := strings.Split(urlString, "/")
		oneTheme := tags[2]

		AppConfig.Theme = oneTheme
		write_config()
	} 

	http.Redirect(w, r, "/", http.StatusSeeOther)
}