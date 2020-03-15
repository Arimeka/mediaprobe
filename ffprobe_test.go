package mediaprobe_test

import (
	"testing"

	"github.com/Arimeka/mediaprobe"
)

var bitrateFiles = map[string]int64{
	"http://localhost:9090/video.mp4": 551193,
	"./fixtures/video.mp4":            551193,
	"./fixtures/with-meta.mov":        481140,
	"./fixtures/audio.mp3":            128000,
	"./fixtures/not-a-video.mp4":      0,
	"./fixtures/corrupted.mp4":        0,
}

func TestInfo_FFProbe(t *testing.T) {
	handler := &Handler{
		Status:   0,
		Filename: "./fixtures/video.mp4",
	}
	srv := ServeHttp(handler)
	defer srv.Stop()

	for filename, expectedBitrate := range bitrateFiles {
		info, err := mediaprobe.New(filename)
		if err != nil {
			t.Fatalf("Filename: %s. Unexpected error %v", filename, err)
		}

		err = info.FFProbe()
		if err != nil {
			if filename != "./fixtures/corrupted.mp4" {
				t.Errorf("Filename: %s. Unexpected error %v", filename, err)
			}
		} else {
			if filename == "./fixtures/corrupted.mp4" {
				t.Errorf("Filename: %s. Expected to return error but return nil", filename)
			}
		}

		bitrate := info.BitRate
		if bitrate != expectedBitrate {
			t.Errorf("Filename: %s. Not expected video bitrate. Expected %d; got %d", filename, expectedBitrate, bitrate)
		}
	}
}
