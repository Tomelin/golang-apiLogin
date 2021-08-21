package controllers

import (
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	redirectURL := "http://127.0.0.1:8000/login"

	http.Redirect(w, r, redirectURL, 302)

}
