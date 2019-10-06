package api

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
)

func (s *Server) handleGetResult(w http.ResponseWriter, r *http.Request) {
	var (
		page     = 1
		pageSize = 10
		desc     = false
	)
	if i, err := strconv.Atoi(r.URL.Query().Get("page")); err == nil {
		page = i
	}
	if i, err := strconv.Atoi(r.URL.Query().Get("page_size")); err == nil {
		pageSize = i
	}
	switch r.URL.Query().Get("desc") {
	case "1", "true":
		desc = true
	}
	count, err := s.svc.Count()
	if err != nil {
		writeMessage(w, http.StatusInternalServerError, err.Error())
		return
	}
	results, err := s.svc.GetAllResult(pageSize, (page-1)*pageSize, desc)
	resp := &commonResp{
		Data:      results,
		Page:      page,
		PageSize:  pageSize,
		PageCount: int(math.Round(float64(count) / float64(pageSize))),
	}
	if err != nil {
		resp.ErrMsg = err.Error()
	}
	writeJSON(w, resp)
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
