package mediaprobe_test

import (
	"net/http"
	"net/http/httptest"
	"time"
)

type Server struct {
	srv *httptest.Server
}

func (s *Server) Stop() {
	s.srv.Close()
}

func (s *Server) Endpoint() string {
	return s.srv.URL
}

func ServeHttp(handler *Handler) *Server {
	httptest.NewServer(handler)

	srv := httptest.NewServer(handler)
	time.Sleep(100 * time.Millisecond)

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
