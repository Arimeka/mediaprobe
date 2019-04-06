package mediaprobe

import (
	"image"
	"os"

	// Expanding image package
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

// ParseImage used for retrieve image data
// TODO: implement calculating aspect-ratio
func (probe *Info) ParseImage(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.DecodeConfig(file)
	if err != nil {
		return err
	}

	probe.Width = img.Width
	probe.Height = img.Height

	return nil
}
