package mediaprobe

import (
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"strings"

	"github.com/h2non/filetype"
)

// UnsafeNew initialized Info without using magic bytes for calculating media type
func UnsafeNew(filename string) (*Info, error) {
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
