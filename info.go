package mediaprobe

import (
	"os"
	"time"
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
		filename: filename,
		Name:     fileinfo.Name(),
		Size:     fileinfo.Size(),
	}

	return info, nil
}

// Info contains parsed information
type Info struct {
	filename string

	Name         string
	MediaType    string
	MediaSubtype string
	Size         int64
	StartTime    time.Duration
	Duration     time.Duration
	BitRate      int64
	Width        int
	Height       int
	Streams      []Stream
}

// Stream contains audio/video stream information
type Stream struct {
	ID             int
	Index          int
	MediaType      string
	Codec          string
	CodecTag       string
	CodecLongName  string
	IsExperimental bool
	Profile        string
	ColorRangeName string
	SampleFmtName  string
	Bitrate        int
	Width          int
	Height         int
	AspectRation   float64
	FrameRate      float64
	AvgFrameRate   float64
	BFrames        int
	BitsPerSample  int
}
