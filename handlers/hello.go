package handlers

import (
	"net/http"
)

func NewHelloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body := map[string]interface{}{"code": 0, "whoami": "base64captcha-http"}
		writeJson(w, body)
	}
}
