package mediaprobe_test

import (
	"testing"

	"github.com/Arimeka/mediaprobe"
)

const (
	testImageValidImage          = "./example/samples/image.jpeg"
	testImageWithExifOrientation = "./example/samples/left.jpg"
	testImageInvalidImage        = "./example/samples/not-an-image.jpeg"
)

func TestInfo_ParseImage(t *testing.T) {
	info, _ := mediaprobe.New(testImageInvalidImage)
	err := info.ParseImage()
	if err == nil {
		t.Errorf("Filename: %s. Expected to return error but return nil", testImageInvalidImage)
	}

	info, _ = mediaprobe.New(testImageValidImage)
	err = info.ParseImage()
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testImageValidImage, err)
	}

	width := info.Width
	if width != 290 {
		t.Errorf("Filename: %s. Not expected width. Expected %d; got %d", testImageValidImage, 290, width)
	}
}

func TestInfo_ParseImageWithOrientation(t *testing.T) {
	info, _ := mediaprobe.New(testImageWithExifOrientation)
	err := info.ParseImage()
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testImageWithExifOrientation, err)
	}

	width := info.Width
	if width != 330 {
		t.Errorf("Filename: %s. Not expected width. Expected %d; got %d", testImageWithExifOrientation, 330, width)
	}
}
