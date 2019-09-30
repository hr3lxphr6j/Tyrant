package api

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"Tyrant/src/service"
)

type commonResp struct {
	ErrMsg    string      `json:"err_msg"`
	Data      interface{} `json:"data"`
	Page      int         `json:"page"`
	PageSize  int         `json:"page_size"`
	PageCount int         `json:"page_count"`
}

type Server struct {
	svc    service.ResultService
	router *mux.Router
	hs     *http.Server
}

func New(svc service.ResultService, bind string) (*Server, error) {
	if svc == nil {
		return nil, errors.New("ResultService is null")
	}
	if bind == "" {
		return nil, errors.New("bind address is null")
	}
	s := &Server{
		svc:    svc,
		router: mux.NewRouter(),
		hs: &http.Server{
			Addr: bind,
		},
	}
	s.initRouterMap()
	return s, nil
}

func (s *Server) initRouterMap() {
	s.hs.Handler = s.router
	s.router.Use(
		// log
		func(handler http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Do stuff here
				log.Println(r.RequestURI)
				// Call the next handler, which can be another middleware in the chain, or the final handler.
				handler.ServeHTTP(w, r)
			})
		},
		// CORS
		func(handler http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodOptions {
					w.Header().Add("Access-Control-Allow-Origin", r.Header.Get("Origin"))
					w.Header().Add("Access-Control-Allow-Methods", "GET")
					w.Write(nil)
				} else {
					w.Header().Add("Access-Control-Allow-Origin", "*")
					handler.ServeHTTP(w, r)
				}
			})
		},
	)
	r := s.router.PathPrefix("/api").Subrouter()
	r.HandleFunc("/result", s.handleGetResult).Methods(http.MethodGet)
}

func (s *Server) Start() error {
	go func() {
		s.hs.ListenAndServe()
	}()
	return nil
}

func (s *Server) Stop() {
	ctx, cancel := context.WithCancel(context.Background())
	s.hs.Shutdown(ctx)
	cancel()
}
