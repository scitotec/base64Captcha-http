package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mojocn/base64Captcha"
)

type verifyBody struct {
	Id     string
	Answer string
}

func NewVerifyHandler(store base64Captcha.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		var param verifyBody
		err := decoder.Decode(&param)
		if err != nil {
			writeJsonError(w, err, 10, http.StatusBadRequest)
			return
		}
		verified := false
		if param.Id != "" {
			verified = store.Verify(param.Id, param.Answer, true)
		}
		body := map[string]interface{}{"code": 0, "msg": "ok", "verified": verified}
		writeJson(w, body)
	}
}
