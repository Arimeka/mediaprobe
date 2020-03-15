package mediaprobe_test

import (
	"testing"

	"github.com/Arimeka/mediaprobe"
)

const (
	testProbeValidRemoteImage = "http://localhost:9090/image.jpeg"
	testProbeValidImage       = "./fixtures/image.jpeg"
	testProbeInvalidImage     = "./fixtures/not-an-image.jpeg"
	testProbeValidVideo       = "./fixtures/video.mp4"
	testProbeInvalidVideo     = "./fixtures/not-a-video.mp4"
	testProbeValidAudio       = "./fixtures/audio.mp3"
	testProbeCorruptedFile    = "./fixtures/corrupted.mp4"
)

func TestParse(t *testing.T) {
	t.Run("not_found", parseNotFound)
	t.Run("valid_remote_image", parseValidRemoteImage)
	t.Run("valid_image", parseValidImage)
	t.Run("invalid_image", parseInvalidImage)
	t.Run("valid_video", parseValidVideo)
	t.Run("invalid_video", parseInvalidVideo)
	t.Run("valid_audio", parseValidAudio)
	t.Run("corrupted_file", parseCorruptedFile)
}

func parseNotFound(t *testing.T) {
	_, err := mediaprobe.Parse("")
	if err == nil {
		t.Errorf("Expected to return error found but return nil")
	}
}

func parseValidImage(t *testing.T) {
	info, err := mediaprobe.Parse(testProbeValidImage)
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testProbeValidImage, err)
	}
	width := info.Width
	if width != 290 {
		t.Errorf("Filename: %s. Not expected width. Expected %d; got %d", testProbeValidImage, 290, width)
	}
}

func parseValidRemoteImage(t *testing.T) {
	handler := &Handler{
		Filename: "./fixtures/image.jpeg",
	}
	srv := ServeHttp(handler)
	defer srv.Stop()

	info, err := mediaprobe.Parse(testProbeValidRemoteImage)
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testProbeValidRemoteImage, err)
	}
	width := info.Width
	if width != 290 {
		t.Errorf("Filename: %s. Not expected width. Expected %d; got %d", testProbeValidRemoteImage, 290, width)
	}
}

func parseInvalidImage(t *testing.T) {
	info, err := mediaprobe.Parse(testProbeInvalidImage)
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testProbeInvalidImage, err)
	}
	bitrate := info.BitRate
	if bitrate != 551193 {
		t.Errorf("Filename: %s. Not expected video bitrate. Expected %d; got %d", testProbeInvalidImage, 551193, bitrate)
	}
}

func parseValidVideo(t *testing.T) {
	info, err := mediaprobe.Parse(testProbeValidVideo)
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testProbeValidVideo, err)
	}
	bitrate := info.BitRate
	if bitrate != 551193 {
		t.Errorf("Filename: %s. Not expected video bitrate. Expected %d; got %d", testProbeValidVideo, 551193, bitrate)
	}
}

func parseInvalidVideo(t *testing.T) {
	info, err := mediaprobe.Parse(testProbeInvalidVideo)
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testProbeInvalidVideo, err)
	}
	width := info.Width
	if width != 290 {
		t.Errorf("Filename: %s. Not expected width. Expected %d; got %d", testProbeInvalidVideo, 290, width)
	}
}

func parseValidAudio(t *testing.T) {
	info, err := mediaprobe.Parse(testProbeValidAudio)
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testProbeValidAudio, err)
	}
	bitrate := info.BitRate
	if bitrate != 128000 {
		t.Errorf("Filename: %s. Not expected video bitrate. Expected %d; got %d", testProbeValidAudio, 128000, bitrate)
	}
}

func parseCorruptedFile(t *testing.T) {
	info, err := mediaprobe.Parse(testProbeCorruptedFile)
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testProbeCorruptedFile, err)
	}
	mediaSubtype := info.MediaSubtype
	if mediaSubtype != "octet-stream" {
		t.Errorf("Filename: %s. Not media subtype. Expected %s; got %s", testProbeValidAudio, "octet-stream", mediaSubtype)
	}
}
