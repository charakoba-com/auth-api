package authapi

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func httpJSONWithStatus(w http.ResponseWriter, status int, v interface{}) {
	buf := bytes.Buffer{}
	if err := json.NewEncoder(&buf).Encode(v); err != nil {
		httpError(w, http.StatusInternalServerError, `encode json`, err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	log.Infof("%d %s", status, buf.String())
	buf.WriteTo(w)
}

func httpJSON(w http.ResponseWriter, v interface{}) {
	httpJSONWithStatus(w, http.StatusOK, v)
}

type jsonErr struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func httpError(w http.ResponseWriter, status int, message string, err error) {
	v := jsonErr{
		Message: message,
	}
	if err != nil {
		v.Error = err.Error()
	}
	httpJSONWithStatus(w, status, v)
}
