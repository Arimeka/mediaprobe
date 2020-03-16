package mediaprobe_test

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
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

	FailOnAttempt uint64
	Counter       uint64
	mu            sync.RWMutex
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if h.Status != 0 {
		rw.WriteHeader(h.Status)
		return
	}
	if h.FailOnAttempt > 0 {
		h.mu.RLock()
		counter := h.Counter
		h.mu.RUnlock()
		if counter >= h.FailOnAttempt {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		atomic.AddUint64(&h.Counter, 1)
	}

	http.ServeFile(rw, req, h.Filename)
}
