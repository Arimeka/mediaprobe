// Package mediaprobe provides functions for parsing media files for information,
// such as dimensions, codecs, duration, etc. It uses bindings to ffmpeg and libmagic.
package mediaprobe

import (
	"fmt"
)

// Parse file media data
// It determines the file type by magic bytes,
// and parses the media data of the video or image.
func Parse(filename string) (Info, error) {
	info, err := New(filename)
	if err != nil {
		return Info{}, fmt.Errorf("can't parse file: %v", err)
	}

	if err = info.CalculateMime(); err != nil {
		return *info, fmt.Errorf("can't parse file: %v", err)
	}

	switch info.MediaType {
	case "image":
		err = info.ParseImage()
	case "video", "audio":
		err = info.FFProbe()
	}
	if err != nil {
		return *info, fmt.Errorf("can't parse file: %v", err)
	}

	return *info, nil
}
