package api

import (
	"math"
	"net/http"
)

func (s *Server) handleGetResult(w http.ResponseWriter, r *http.Request) {
	var (
		page     = 1
		pageSize = 10
		desc     = false
	)
	loadIntReqParams(&page, r, "page")
	loadIntReqParams(&pageSize, r, "page_size")
	loadBoolReqParams(&desc, r, "desc")

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
