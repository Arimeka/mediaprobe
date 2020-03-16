package mediaprobe_test

import (
	"net/http"
	"testing"

	"github.com/Arimeka/mediaprobe"
)

func TestNew(t *testing.T) {
	t.Run("not_found", newNotFound)
	t.Run("not_found_remote", newNotFoundRemote)
	t.Run("conn_error_remote", newConnErrorRemote)
	t.Run("server_error_remote", newServerErrorRemote)
	t.Run("local", newLocalFile)
	t.Run("remote", newRemoteFile)
}

func newNotFound(t *testing.T) {
	_, err := mediaprobe.New("")
	if err == nil {
		t.Errorf("Expected to return error found but return nil")
	}
}

func newNotFoundRemote(t *testing.T) {
	handler := &Handler{
		Status: http.StatusNotFound,
	}
	srv := ServeHttp(handler)
	defer srv.Stop()

	_, err := mediaprobe.New(srv.Endpoint())
	if err == nil {
		t.Errorf("Expected to return error found but return nil")
	}
}

func newConnErrorRemote(t *testing.T) {
	_, err := mediaprobe.New("http://localhost:9091/not-exist")
	if err == nil {
		t.Errorf("Expected to return error found but return nil")
	}
}

func newServerErrorRemote(t *testing.T) {
	handler := &Handler{
		Filename:      "./fixtures/video.mp4",
		FailOnAttempt: 1,
	}
	srv := ServeHttp(handler)
	defer srv.Stop()

	_, err := mediaprobe.New(srv.Endpoint())
	if err == nil {
		t.Errorf("Expected to return error found but return nil")
	}
}

func newLocalFile(t *testing.T) {
	info, err := mediaprobe.New("./fixtures/video.mp4")
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	if info.Size != 383631 {
		t.Errorf("Unexpected size. Expected %d, got %d", 383631, info.Size)
	}
}

func newRemoteFile(t *testing.T) {
	handler := &Handler{
		Filename: "./fixtures/video.mp4",
	}
	srv := ServeHttp(handler)
	defer srv.Stop()

	info, err := mediaprobe.New(srv.Endpoint())
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	if info.Size != 383631 {
		t.Errorf("Unexpected size. Expected %d, got %d", 383631, info.Size)
	}
}
