package main

import (
	"flag"
	"log"

	"github.com/Arimeka/mediaprobe"
)

func main() {
	var srcFileName string

	flag.StringVar(&srcFileName, "src", "./samples/video.mp4", "source")
	flag.Parse()

	info, err := mediaprobe.Parse(srcFileName)
	if err != nil {
		log.Fatalf("Error Parse - %s\n", err)
	}

	log.Printf("Safe Parse: %+v\n", info)
}
