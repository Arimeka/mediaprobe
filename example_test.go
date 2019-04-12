package mediaprobe_test

import (
	"log"

	"github.com/Arimeka/mediaprobe"
)

func ExampleParse() {
	info, err := mediaprobe.Parse("./fixtures/video.mp4")
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	log.Printf("Result: %+v\n", info)
}

func ExampleInfo_FFProbe() {
	info, err := mediaprobe.New("./fixtures/video.mp4")
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	// Skipping calculate mime-type by magic numbers and immediately parse video
	if err = info.FFProbe(); err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	log.Printf("Result: %+v\n", info)
}

func ExampleInfo_ParseImage() {
	info, err := mediaprobe.New("./fixtures/image.jpeg")
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	// Skipping calculate mime-type by magic numbers and immediately parse image
	if err = info.ParseImage(); err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	log.Printf("Result: %+v\n", info)
}
