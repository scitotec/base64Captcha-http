package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mojocn/base64Captcha"
)

func NewAudioHandler(store base64Captcha.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var driverConfig base64Captcha.DriverAudio
		handle(&driverConfig, store, w, r)
	}
}
func NewChineseHandler(store base64Captcha.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var driverConfig base64Captcha.DriverChinese
		handle(&driverConfig, store, w, r)
	}
}
func NewDigitHandler(store base64Captcha.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var driverConfig base64Captcha.DriverDigit
		handle(&driverConfig, store, w, r)
	}
}
func NewLanguageHandler(store base64Captcha.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var driverConfig base64Captcha.DriverLanguage
		handle(&driverConfig, store, w, r)
	}
}
func NewMathHandler(store base64Captcha.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var driverConfig base64Captcha.DriverMath
		handle(&driverConfig, store, w, r)
	}
}
func NewStringHandler(store base64Captcha.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var driverConfig base64Captcha.DriverString
		handle(&driverConfig, store, w, r)
	}
}

func handle(driverConfig base64Captcha.Driver, store base64Captcha.Store, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := decoder.Decode(&driverConfig)
	if err != nil {
		writeJsonError(w, err, 10, http.StatusBadRequest)
		return
	}
	c := base64Captcha.NewCaptcha(driverConfig, store)
	id, b64s, err := c.Generate()
	if err != nil {
		writeJsonError(w, err, 20, http.StatusInternalServerError)
		return
	}
	body := map[string]interface{}{"code": 0, "data": b64s, "captchaId": id, "msg": "success"}
	writeJson(w, body)
}
