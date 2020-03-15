package mediaprobe_test

import (
	"testing"

	"github.com/Arimeka/mediaprobe"
)

func TestInfo_CalculateMime(t *testing.T) {
	t.Run("not_found", calculateMimeNotFound)
	t.Run("local", calculateMimeLocal)
	t.Run("remote", calculateMimeRemote)
}

func calculateMimeNotFound(t *testing.T) {
	info := &mediaprobe.Info{}
	err := info.CalculateMime()
	if err == nil {
		t.Errorf("Expected to return error found but return nil")
	}
}

func calculateMimeLocal(t *testing.T) {
	info, err := mediaprobe.New("./fixtures/not-an-image.jpeg")
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	err = info.CalculateMime()
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	if info.MediaType != "video" {
		t.Errorf("Unexpected media type. Expected %s, got %s", "video", info.MediaType)
	}
}

func calculateMimeRemote(t *testing.T) {
	handler := &Handler{
		Filename: "./fixtures/not-an-image.jpeg",
	}
	srv := ServeHttp(handler)
	defer srv.Stop()

	info, err := mediaprobe.New("http://localhost:9090/not-an-image.jpeg")
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	err = info.CalculateMime()
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	if info.MediaType != "video" {
		t.Errorf("Unexpected media type. Expected %s, got %s", "video", info.MediaType)
	}
}
