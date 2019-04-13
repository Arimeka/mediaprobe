package mediaprobe_test

import (
	"testing"

	"github.com/Arimeka/mediaprobe"
)

const (
	testImageValidImage           = "./fixtures/image.jpeg"
	testImageWithExifOrientation  = "./fixtures/left.jpg"
	testImageInvalidImage         = "./fixtures/not-an-image.jpeg"
	testImageJPEGWithoutExifImage = "./fixtures/without-exif.jpg"
	testImagePNGWithoutExifImage  = "./fixtures/without-exif.png"
)

func TestInfo_ParseImageNotFound(t *testing.T) {
	info := &mediaprobe.Info{}
	err := info.ParseImage()
	if err == nil {
		t.Errorf("Expected to return error found but return nil")
	}
}

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

func TestInfo_ParseImageJPEGWithoutExif(t *testing.T) {
	info, _ := mediaprobe.New(testImageJPEGWithoutExifImage)
	err := info.ParseImage()
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testImageJPEGWithoutExifImage, err)
	}

	width := info.Width
	if width != 200 {
		t.Errorf("Filename: %s. Not expected width. Expected %d; got %d", testImageJPEGWithoutExifImage, 200, width)
	}
}

func TestInfo_ParseImagePNGWithoutExif(t *testing.T) {
	info, _ := mediaprobe.New(testImagePNGWithoutExifImage)
	err := info.ParseImage()
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testImagePNGWithoutExifImage, err)
	}

	width := info.Width
	if width != 100 {
		t.Errorf("Filename: %s. Not expected width. Expected %d; got %d", testImagePNGWithoutExifImage, 100, width)
	}
}
