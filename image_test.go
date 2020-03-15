package mediaprobe_test

import (
	"net/http"
	"testing"

	"github.com/Arimeka/mediaprobe"
)

const (
	testImageValidRemoteImage     = "http://localhost:9090/image.jpeg"
	testImageValidImage           = "./fixtures/image.jpeg"
	testImageWithExifOrientation  = "./fixtures/left.jpg"
	testImageInvalidImage         = "./fixtures/not-an-image.jpeg"
	testImageJPEGWithoutExifImage = "./fixtures/without-exif.jpg"
	testImagePNGWithoutExifImage  = "./fixtures/without-exif.png"
)

func TestInfo_ParseImage(t *testing.T) {
	t.Run("not_found", parseImageNotFound)
	t.Run("not_found_remote", parseImageNotFoundRemote)
	t.Run("conn_error_remote", parseImageConnectionErrorRemote)
	t.Run("valid", parseImageValid)
	t.Run("valid_remote", parseImageValidRemote)
	t.Run("invalid", parseImageInvalid)
	t.Run("with_orientation", parseImageWithOrientation)
	t.Run("jpeg_without_exif", parseImageJPEGWithoutExif)
	t.Run("png_without_exif", parseImagePNGWithoutExif)
}

func parseImageNotFound(t *testing.T) {
	info := &mediaprobe.Info{}
	err := info.ParseImage()
	if err == nil {
		t.Errorf("Expected to return error found but return nil")
	}
}

func parseImageNotFoundRemote(t *testing.T) {
	handler := &Handler{
		Filename: "./fixtures/image.jpeg",
	}
	srv := ServeHttp(handler)
	defer srv.Stop()

	info, err := mediaprobe.New(testImageValidRemoteImage)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	handler.Status = http.StatusNotFound
	err = info.ParseImage()
	if err == nil {
		t.Errorf("Expected to return error found but return nil")
	}
}

func parseImageConnectionErrorRemote(t *testing.T) {
	handler := &Handler{
		Filename: "./fixtures/image.jpeg",
	}
	srv := ServeHttp(handler)

	info, err := mediaprobe.New(testImageValidRemoteImage)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	srv.Stop()
	err = info.ParseImage()
	if err == nil {
		t.Errorf("Expected to return error found but return nil")
	}
}

func parseImageValid(t *testing.T) {
	info, err := mediaprobe.New(testImageValidImage)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	err = info.ParseImage()
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testImageValidImage, err)
	}

	width := info.Width
	if width != 290 {
		t.Errorf("Filename: %s. Not expected width. Expected %d; got %d", testImageValidImage, 290, width)
	}
}

func parseImageValidRemote(t *testing.T) {
	handler := &Handler{
		Filename: "./fixtures/image.jpeg",
	}
	srv := ServeHttp(handler)
	defer srv.Stop()

	info, err := mediaprobe.New(testImageValidRemoteImage)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	err = info.ParseImage()
	if err != nil {
		t.Errorf("Filename: %s. Unexpected error %v", testImageValidImage, err)
	}

	width := info.Width
	if width != 290 {
		t.Errorf("Filename: %s. Not expected width. Expected %d; got %d", testImageValidImage, 290, width)
	}
}

func parseImageInvalid(t *testing.T) {
	info, err := mediaprobe.New(testImageInvalidImage)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	err = info.ParseImage()
	if err == nil {
		t.Errorf("Filename: %s. Expected to return error but return nil", testImageInvalidImage)
	}
}

func parseImageWithOrientation(t *testing.T) {
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

func parseImageJPEGWithoutExif(t *testing.T) {
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

func parseImagePNGWithoutExif(t *testing.T) {
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
