package mediaprobe_test

import (
	"log"

	"github.com/Arimeka/mediaprobe"
)

func ExampleParse() {
	info, err := mediaprobe.Parse("./samples/video.mp4")
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	log.Printf("Result: %+v\n", info)
}
