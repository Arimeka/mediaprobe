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
func (info *Info) ParseImage() error {
	file, err := os.Open(info.filename)
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

	info.Width = img.Width
	info.Height = img.Height

	if x, err := exif.Decode(file); err == nil {
		if orientationTag, err := x.Get(exif.Orientation); err == nil {
			switch orientationTag.String() {
			case "5", "6", "7", "8":
				info.Width = img.Height
				info.Height = img.Width
			}
		}
	}

	return nil
}
