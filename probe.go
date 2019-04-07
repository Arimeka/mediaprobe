package mediaprobe

import (
	"fmt"
	"os"
	"strings"

	"github.com/rakyll/magicmime"
)

// New initialized Info using magic bytes for calculating media type
func New(filename string) (*Info, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	info := &Info{
		Name: fileinfo.Name(),
		Size: fileinfo.Size(),
	}

	if err := magicmime.Open(magicmime.MAGIC_MIME_TYPE | magicmime.MAGIC_SYMLINK | magicmime.MAGIC_ERROR); err != nil {
		return nil, err
	}
	defer magicmime.Close()

	media, err := magicmime.TypeByFile(filename)
	if err != nil {
		return nil, err
	}

	kind := strings.Split(media, "/")
	if len(kind) == 2 {
		info.MediaType = kind[0]
		info.MediaSubtype = kind[1]
	} else {
		info.MediaType = "application"
		info.MediaSubtype = "octet-stream"
	}

	return info, nil
}

// Parse parsing file media data
func Parse(filename string) (Info, error) {
	info, err := New(filename)
	if err != nil {
		return Info{}, fmt.Errorf("can't probe file: %v", err)
	}

	switch info.MediaType {
	case "image":
		err = info.ParseImage(filename)
	case "video", "audio":
		err = info.FFProbe(filename)
	}
	if err != nil {
		return Info{}, fmt.Errorf("can't probe file: %v", err)
	}

	return *info, nil
}
