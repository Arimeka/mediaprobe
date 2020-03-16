package mediaprobe_test

import (
	"testing"

	"github.com/Arimeka/mediaprobe"
)

const (
	benchmarkValidVideo = "./fixtures/video.mp4"
	benchmarkValidImage = "./fixtures/image.jpeg"
)

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = mediaprobe.Parse(benchmarkValidVideo)
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = mediaprobe.New(benchmarkValidVideo)
	}
}

func BenchmarkInfo_CalculateMime(b *testing.B) {
	info, _ := mediaprobe.New(benchmarkValidVideo)
	for i := 0; i < b.N; i++ {
		_ = info.CalculateMime()
	}
}

func BenchmarkInfo_FFProbe(b *testing.B) {
	info, _ := mediaprobe.New(benchmarkValidVideo)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = info.FFProbe()
	}
}

func BenchmarkInfo_ParseImage(b *testing.B) {
	info, _ := mediaprobe.New(benchmarkValidImage)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = info.FFProbe()
	}
}
