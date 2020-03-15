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
	if info.data == nil {
		file, err := os.Open(info.filename)
		if err != nil {
			return err
		}
		defer file.Close()
		if err := info.decodeImage(file); err != nil {
			return err
		}
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			return err
		}
		info.decodeExif(file)

		return nil
	}

	body, err := getRemoteFile(info.uri)
	if err != nil {
		return err
	}
	err = info.decodeImage(body)
	body.Close()
	if err != nil {
		return err
	}

	body, err = getRemoteFile(info.uri)
	if err != nil {
		return err
	}
	info.decodeExif(body)
	body.Close()

	return nil
}

func (info *Info) decodeImage(reader io.Reader) error {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		return err
	}

	info.Width = img.Width
	info.Height = img.Height

	return nil
}

func (info *Info) decodeExif(reader io.Reader) {
	if x, err := exif.Decode(reader); err == nil {
		if orientationTag, err := x.Get(exif.Orientation); err == nil {
			switch orientationTag.String() {
			case "5", "6", "7", "8":
				w, h := info.Width, info.Height
				info.Width = h
				info.Height = w
			}
		}
	}
}
