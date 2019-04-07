package mediaprobe_test

import (
	"testing"

	"mediaprobe"
)

var bitrateFiles = map[string]int64{
	"./example/samples/video.mp4":       551193,
	"./example/samples/audio.mp3":       128000,
	"./example/samples/not-a-video.mp4": 0,
	"./example/samples/corrupted.mp4":   0,
}

func TestInfo_FFProbe(t *testing.T) {
	for filename, expectedBitrate := range bitrateFiles {
		info, _ := mediaprobe.New(filename)

		err := info.FFProbe()
		if err != nil {
			if filename != "./example/samples/corrupted.mp4" {
				t.Errorf("Filename: %s. Unexpected error %v", filename, err)
			}
		} else {
			if filename == "./example/samples/corrupted.mp4" {
				t.Errorf("Filename: %s. Expected to return error but return nil", filename)
			}
		}

		bitrate := info.BitRate
		if bitrate != expectedBitrate {
			t.Errorf("Filename: %s. Not expected video bitrate. Expected %d; got %d", filename, expectedBitrate, bitrate)
		}
	}
}
