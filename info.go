package mediaprobe

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

// New initialized Info and calculate file size
func New(filename string) (*Info, error) {
	uri, err := url.Parse(filename)
	if err != nil || !uri.IsAbs() {
		return openFile(filename)
	}

	return openURL(filename, uri)
}

func openFile(filename string) (*Info, error) {
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

func openURL(filename string, uri *url.URL) (*Info, error) {
	headReq := &http.Request{
		Method: http.MethodHead,
		URL:    uri,
	}
	head, err := httpClient.Do(headReq)
	if err != nil {
		return nil, err
	}
	if head.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", head.StatusCode)
	}
	head.Body.Close()

	body, err := getRemoteFile(uri)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	info := &Info{
		data:     make([]byte, 1024),
		filename: filename,
		uri:      uri,

		Name: filename,
		Size: head.ContentLength,
	}

	_, err = body.Read(info.data)
	if err != nil {
		return nil, err
	}

	return info, nil
}

// Info contains parsed information
type Info struct {
	filename string
	uri      *url.URL
	data     []byte

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
