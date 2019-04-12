package mediaprobe_test

import (
	"testing"

	"github.com/Arimeka/mediaprobe"
)

func TestInfo_CalculateMimeNotFound(t *testing.T) {
	info := &mediaprobe.Info{}
	err := info.CalculateMime()
	if err == nil {
		t.Errorf("Expected to return error found but return nil")
	}
}
