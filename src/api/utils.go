package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func loadIntReqParams(dst *int, r *http.Request, name string) {
	if dst == nil || r == nil || name == "" {
		return
	}
	num, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		return
	}
	*dst = num
}

func loadBoolReqParams(dst *bool, r *http.Request, name string) {
	if dst == nil || r == nil || name == "" {
		return
	}
	switch r.URL.Query().Get(name) {
	case "1", "true":
		*dst = true
	}
}

func writeJSON(w http.ResponseWriter, obj interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(obj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	_, err = w.Write(b)
	return err
}

func writeMessage(w http.ResponseWriter, statusCode int, msg string) {
	w.WriteHeader(statusCode)
	writeJSON(w, &commonResp{ErrMsg: msg})
}
