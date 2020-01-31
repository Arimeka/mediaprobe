package mediaprobe

import (
	"time"

	"github.com/3d0c/gmf"
)

// FFProbe parses video or audio using ffmpeg bindings
// It uses github.com/3d0c/gmf package
func (info *Info) FFProbe() error {
	inputCtx, err := gmf.NewInputCtx(info.filename)
	if err != nil {
		return err
	}
	defer inputCtx.Free()

	info.Duration = time.Duration(inputCtx.Duration() * float64(time.Second))
	info.StartTime = time.Duration(inputCtx.StartTime() * int(time.Second))
	info.BitRate = inputCtx.BitRate()

	for idx := 0; idx < inputCtx.StreamsCnt(); idx++ {
		stream, err := inputCtx.GetStream(idx)
		if err != nil {
			return err
		}

		codecCtx := stream.CodecCtx()
		if codecCtx == nil {
			// do nothing if can't extract codec from stream
			continue
		}
		codec := codecCtx.Codec()

		streamInfo := Stream{
			ID:             stream.Id(),
			Index:          stream.Index(),
			Bitrate:        codecCtx.BitRate(),
			MediaType:      codecCtx.GetMediaType(),
			Codec:          codec.Name(),
			CodecLongName:  codec.LongName(),
			CodecTag:       codecCtx.GetCodecTagName(),
			IsExperimental: codec.IsExperimental(),
			SampleFmtName:  codecCtx.GetSampleFmtName(),
		}

		if stream.IsVideo() {
			streamInfo.Width = codecCtx.Width()
			streamInfo.Height = codecCtx.Height()
			streamInfo.FrameRate = stream.GetRFrameRate().AVR().Av2qd()
			streamInfo.AvgFrameRate = stream.GetAvgFrameRate().AVR().Av2qd()
			streamInfo.Profile = codecCtx.GetProfileName()
			streamInfo.ColorRangeName = codecCtx.GetColorRangeName()
			streamInfo.AspectRation = codecCtx.GetAspectRation().AVR().Av2qd()
			streamInfo.BFrames = codecCtx.GetBFrames()
			streamInfo.BitsPerSample = codecCtx.GetBitsPerSample()
		}

		codecCtx.Free()
		stream.Free()

		info.Streams = append(info.Streams, streamInfo)
	}

	return nil
}
