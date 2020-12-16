package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"yam-api/source/utils/log"
)

func ResError(code int, w http.ResponseWriter, msg string) {
	ResJSON(code, w, map[string]interface{}{"code": code, "message": msg})
}

func ResByte(c int, h http.ResponseWriter, payload []byte) {
	h.Header().Set("Content-Type", "application/json")
	h.WriteHeader(c)
	h.Write(payload)
}

func ResJSON(c int, h http.ResponseWriter, payload interface{}) {
	response, _ := json.Marshal(payload)
	h.Header().Set("Content-Type", "application/json")
	h.WriteHeader(c)
	h.Write(response)
}

func MakeRequest(url string) ([]byte, error) {
	resp, respErr := http.Get(url)
	if respErr != nil {
		log.Error(respErr)
		return nil, respErr
	}
	body, bodyErr := ioutil.ReadAll(resp.Body)
	if bodyErr != nil {
		log.Error(bodyErr)
		return nil, bodyErr
	}
	return []byte(body), nil
}

func BuildUrl(r *http.Request, u string) string {
	url := u
	from := r.URL.Query().Get("from")
	if from != "" {
		url = url + "?from=" + from
	}
	to := r.URL.Query().Get("to")
	if to != "" {
		url = url + "?to=" + to
	}
	page := r.URL.Query().Get("page")
	if page != "" {
		url = url + "?page=" + page
	}
	log.Info("url", url)
	return url
}

func CheckSuccess(code int) bool {
	value := true
	switch code {
	case 200:
		value = true
	case 201:
		value = true
	case 429:
		value = false
	case 451:
		value = false
	}
	return value
}
