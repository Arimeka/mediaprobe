package mediaprobe_test

import (
	"testing"

	"github.com/Arimeka/mediaprobe"
)

func TestNewNotFound(t *testing.T) {
	_, err := mediaprobe.New("")
	if err == nil {
		t.Errorf("Expected to return error found but return nil")
	}
}
