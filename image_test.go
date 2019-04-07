package mediaprobe_test

import (
	"mediaprobe"
	"testing"
)

const (
	testImageValidImage   = "./example/samples/image.jpeg"
	testImageInvalidImage = "./example/samples/not-an-image.jpeg"
)

func TestInfo_ParseImage(t *testing.T) {
	info, _ := mediaprobe.UnsafeNew(testImageInvalidImage)
	err := info.ParseImage(testImageInvalidImage)
	if err == nil {
		t.Errorf("Filename: %s. Expected to return error but return nil", testImageInvalidImage)
	}

	info, _ = mediaprobe.UnsafeNew(testImageValidImage)
	err = info.ParseImage(testImageValidImage)
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testImageValidImage, err)
	}

	width := info.Width
	if width != 290 {
		t.Errorf("Filename: %s. Not expected width. Expected %d; got %d", testImageValidImage, 290, width)
	}
}
