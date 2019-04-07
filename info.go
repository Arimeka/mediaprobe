package mediaprobe

import "time"

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

// Stream used for contain audio/video stream information
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
