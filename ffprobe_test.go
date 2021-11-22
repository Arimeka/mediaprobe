package mediaprobe_test

import (
	"testing"

	"github.com/Arimeka/mediaprobe"
)

var bitrateFiles = map[string]int64{
	"./fixtures/video.mp4":       551193,
	"./fixtures/with-meta.mov":   481140,
	"./fixtures/audio.mp3":       128000,
	"./fixtures/not-a-video.mp4": 0,
	"./fixtures/corrupted.mp4":   0,
}

func TestInfo_FFProbe(t *testing.T) {

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

	t.Run("remote_file", ffprobeRemoteFile)
	t.Run("pixel format", ffprobePixelFormat)
}

func ffprobeRemoteFile(t *testing.T) {
	handler := &Handler{
		Filename: "./fixtures/video.mp4",
	}
	srv := ServeHttp(handler)
	defer srv.Stop()

	info, err := mediaprobe.New(srv.Endpoint())
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	err = info.FFProbe()
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	bitrate := info.BitRate
	if bitrate != 551193 {
		t.Errorf("Not expected video bitrate. Expected %d; got %d", 551193, bitrate)
	}
}

func ffprobePixelFormat(t *testing.T) {
	info, err := mediaprobe.New("./fixtures/video.mp4")
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	err = info.FFProbe()
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if info.Streams[0].PixFmtName != "yuv420p" {
		t.Errorf("Not expected pixel format. Expected %s; got %s", "yuv420p", info.Streams[0].PixFmtName)
	}
}
