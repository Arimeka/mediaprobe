package mediaprobe

import (
	"time"

	"github.com/3d0c/gmf"
)

func (probe *Info) FFProbe(filepath string) error {
	inputCtx, err := gmf.NewInputCtx(filepath)
	if err != nil {
		return err
	}
	defer inputCtx.Free()

	probe.Duration = time.Duration(inputCtx.Duration() * float64(time.Second))
	probe.StartTime = time.Duration(inputCtx.StartTime() * int(time.Second))
	probe.BitRate = inputCtx.BitRate()

	for idx := 0; idx < inputCtx.StreamsCnt(); idx++ {
		stream, err := inputCtx.GetStream(idx)
		if err != nil {
			return err
		}

		codec := stream.GetCodecPar()

		streamInfo := Stream{
			ID:      stream.Id(),
			Index:   stream.Index(),
			Bitrate: codec.GetBitRate(),
		}

		if stream.IsVideo() {
			streamInfo.Width = codec.GetWidth()
			streamInfo.Height = codec.GetHeight()
			streamInfo.FrameRate = stream.GetRFrameRate().AVR().Av2qd()
			streamInfo.AvgFrameRate = stream.GetAvgFrameRate().AVR().Av2qd()
		}

		stream.Free()

		probe.Streams = append(probe.Streams, streamInfo)
	}

	return nil
}
