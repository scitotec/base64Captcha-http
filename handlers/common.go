package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func HijackNotFound(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			writeJsonError(w, errors.New("path not mapped"), 1, http.StatusNotFound)
			return
		}
		handler(w, r)
	}
}

func writeJsonError(w http.ResponseWriter, err error, code int, status int) {
	log.Println(err)
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(status)
	body := map[string]interface{}{"code": code, "msg": err.Error()}
	writeJson(w, body)
}

func writeUnknownError(w http.ResponseWriter, err error) {
	writeJsonError(w, err, 1, http.StatusInternalServerError)
}

func writeJson(w http.ResponseWriter, body any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		writeUnknownError(w, err)
	}
}

func OnlyPost(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			writeJsonError(w, errors.New("method not allowed"), 1, http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}
