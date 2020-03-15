package mediaprobe_test

import (
	"context"
	"net/http"
)

type Server struct {
	srv *http.Server
}

func (s *Server) Stop() {
	_ = s.srv.Shutdown(context.Background())
}

func ServeHttp(handler *Handler) *Server {
	srv := &http.Server{
		Addr:    ":9090",
		Handler: handler,
	}
	go func(s *http.Server) {
		_ = s.ListenAndServe()
	}(srv)

	return &Server{srv: srv}
}

type Handler struct {
	Status   int
	Filename string
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if h.Status != 0 {
		rw.WriteHeader(h.Status)
		return
	}

	http.ServeFile(rw, req, h.Filename)
}
