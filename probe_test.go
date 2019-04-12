package mediaprobe_test

import (
	"testing"

	"github.com/Arimeka/mediaprobe"
)

const (
	testProbeValidImage    = "./fixtures/image.jpeg"
	testProbeInvalidImage  = "./fixtures/not-an-image.jpeg"
	testProbeValidVideo    = "./fixtures/video.mp4"
	testProbeInvalidVideo  = "./fixtures/not-a-video.mp4"
	testProbeValidAudio    = "./fixtures/audio.mp3"
	testProbeCorruptedFile = "./fixtures/corrupted.mp4"
)

func TestParseNotFound(t *testing.T) {
	_, err := mediaprobe.Parse("")
	if err == nil {
		t.Errorf("Expected to return error found but return nil")
	}
}

func TestParseValidImage(t *testing.T) {
	info, err := mediaprobe.Parse(testProbeValidImage)
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testProbeValidImage, err)
	}
	width := info.Width
	if width != 290 {
		t.Errorf("Filename: %s. Not expected width. Expected %d; got %d", testProbeValidImage, 290, width)
	}
}

func TestParseInvalidImage(t *testing.T) {
	info, err := mediaprobe.Parse(testProbeInvalidImage)
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testProbeInvalidImage, err)
	}
	bitrate := info.BitRate
	if bitrate != 551193 {
		t.Errorf("Filename: %s. Not expected video bitrate. Expected %d; got %d", testProbeInvalidImage, 551193, bitrate)
	}
}

func TestParseValidVideo(t *testing.T) {
	info, err := mediaprobe.Parse(testProbeValidVideo)
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testProbeValidVideo, err)
	}
	bitrate := info.BitRate
	if bitrate != 551193 {
		t.Errorf("Filename: %s. Not expected video bitrate. Expected %d; got %d", testProbeValidVideo, 551193, bitrate)
	}
}

func TestParseInvalidVideo(t *testing.T) {
	info, err := mediaprobe.Parse(testProbeInvalidVideo)
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testProbeInvalidVideo, err)
	}
	width := info.Width
	if width != 290 {
		t.Errorf("Filename: %s. Not expected width. Expected %d; got %d", testProbeInvalidVideo, 290, width)
	}
}

func TestParseValidAudio(t *testing.T) {
	info, err := mediaprobe.Parse(testProbeValidAudio)
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testProbeValidAudio, err)
	}
	bitrate := info.BitRate
	if bitrate != 128000 {
		t.Errorf("Filename: %s. Not expected video bitrate. Expected %d; got %d", testProbeValidAudio, 128000, bitrate)
	}
}

func TestParseCorruptedFile(t *testing.T) {
	info, err := mediaprobe.Parse(testProbeCorruptedFile)
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testProbeCorruptedFile, err)
	}
	mediaSubtype := info.MediaSubtype
	if mediaSubtype != "octet-stream" {
		t.Errorf("Filename: %s. Not media subtype. Expected %s; got %s", testProbeValidAudio, "octet-stream", mediaSubtype)
	}
}
