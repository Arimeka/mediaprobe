package mediaprobe

import (
	"image"
	"io"
	"os"

	// Expanding image package
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/rwcarlsen/goexif/exif"
)

// ParseImage used for retrieve image data
func (probe *Info) ParseImage() error {
	file, err := os.Open(probe.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.DecodeConfig(file)
	if err != nil {
		return err
	}
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	x, err := exif.Decode(file)
	if err != nil {
		return err
	}

	if orientationTag, err := x.Get(exif.Orientation); err == nil {
		switch orientationTag.String() {
		case "5", "6", "7", "8":
			probe.Width = img.Height
			probe.Height = img.Width
		default:
			probe.Width = img.Width
			probe.Height = img.Height
		}

		return nil
	}

	probe.Width = img.Width
	probe.Height = img.Height

	return nil
}
