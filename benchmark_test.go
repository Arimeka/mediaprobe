package mediaprobe_test

import (
	"os/exec"
	"testing"

	"github.com/Arimeka/mediaprobe"
)

const (
	benchmarkValidVideo = "./example/samples/video.mp4"
	benchmarkValidImage = "./example/samples/image.jpeg"
)

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mediaprobe.Parse(benchmarkValidVideo)
	}
}

func BenchmarkInfo_FFProbe(b *testing.B) {
	info, _ := mediaprobe.New(benchmarkValidVideo)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		info.FFProbe()
	}
}

func BenchmarkInfo_FFProbeCli(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd := exec.Command("ffprobe", "-v quiet", "-print_format json", "-show_format", "-show_streams", benchmarkValidVideo)
		cmd.Run()
	}
}

func BenchmarkInfo_ParseImage(b *testing.B) {
	info, _ := mediaprobe.New(benchmarkValidImage)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		info.FFProbe()
	}
}
