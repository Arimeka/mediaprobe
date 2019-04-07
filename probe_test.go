package mediaprobe_test

import (
	"testing"

	"github.com/Arimeka/mediaprobe"
)

const (
	testProbeValidImage    = "./example/samples/image.jpeg"
	testProbeInvalidImage  = "./example/samples/not-an-image.jpeg"
	testProbeValidVideo    = "./example/samples/video.mp4"
	testProbeInvalidVideo  = "./example/samples/not-a-video.mp4"
	testProbeValidAudio    = "./example/samples/audio.mp3"
	testProbeCorruptedFile = "./example/samples/corrupted.mp4"
)

func TestParse(t *testing.T) {
	info, err := mediaprobe.Parse(testProbeValidImage)
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testProbeValidImage, err)
	}
	width := info.Width
	if width != 290 {
		t.Errorf("Filename: %s. Not expected width. Expected %d; got %d", testProbeValidImage, 290, width)
	}

	info, err = mediaprobe.Parse(testProbeInvalidImage)
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testProbeInvalidImage, err)
	}
	bitrate := info.BitRate
	if bitrate != 551193 {
		t.Errorf("Filename: %s. Not expected video bitrate. Expected %d; got %d", testProbeInvalidImage, 551193, bitrate)
	}

	info, err = mediaprobe.Parse(testProbeValidVideo)
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testProbeValidVideo, err)
	}
	bitrate = info.BitRate
	if bitrate != 551193 {
		t.Errorf("Filename: %s. Not expected video bitrate. Expected %d; got %d", testProbeValidVideo, 551193, bitrate)
	}

	info, err = mediaprobe.Parse(testProbeInvalidVideo)
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testProbeInvalidVideo, err)
	}
	width = info.Width
	if width != 290 {
		t.Errorf("Filename: %s. Not expected width. Expected %d; got %d", testProbeInvalidVideo, 290, width)
	}

	info, err = mediaprobe.Parse(testProbeValidAudio)
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testProbeValidAudio, err)
	}
	bitrate = info.BitRate
	if bitrate != 128000 {
		t.Errorf("Filename: %s. Not expected video bitrate. Expected %d; got %d", testProbeValidAudio, 128000, bitrate)
	}

	info, err = mediaprobe.Parse(testProbeCorruptedFile)
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testProbeCorruptedFile, err)
	}
	mediaSubtype := info.MediaSubtype
	if mediaSubtype != "octet-stream" {
		t.Errorf("Filename: %s. Not media subtype. Expected %s; got %s", testProbeValidAudio, "octet-stream", mediaSubtype)
	}
}
