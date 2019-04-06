package mediaprobe

import (
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"strings"

	"github.com/h2non/filetype"
)

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

	ext := filepath.Ext(fileinfo.Name())
	media := strings.Split(mime.TypeByExtension(ext), "/")
	if len(media) == 2 {
		info.MediaType = media[0]
		info.MediaSubtype = media[1]
	} else {
		info.MediaType = "octet"
		info.MediaSubtype = "stream"
	}

	return info, nil
}

func SafeNew(filepath string) (*Info, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	head := make([]byte, 261)
	_, err = file.Read(head)
	if err != nil {
		return nil, err
	}

	kind, err := filetype.Match(head)
	if err != nil {
		return nil, err
	}

	return &Info{
		Name:         fileinfo.Name(),
		MediaType:    kind.MIME.Type,
		MediaSubtype: kind.MIME.Subtype,
		Size:         fileinfo.Size(),
	}, nil
}

func Parse(filepath string) (Info, error) {
	info, err := SafeNew(filepath)
	if err != nil {
		return Info{}, fmt.Errorf("can't probe file: %v", err)
	}

	switch info.MediaType {
	case "image":
	case "video", "audio":
		err = info.FFProbe(filepath)
	}
	if err != nil {
		return Info{}, fmt.Errorf("can't probe file: %v", err)
	}

	return *info, nil
}
