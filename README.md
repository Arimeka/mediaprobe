[![Build Status](https://travis-ci.org/Arimeka/mediaprobe.svg?branch=master)](https://travis-ci.org/Arimeka/mediaprobe)
[![Coverage Status](https://coveralls.io/repos/github/Arimeka/mediaprobe/badge.svg?branch=master)](https://coveralls.io/github/Arimeka/mediaprobe?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/Arimeka/mediaprobe)](https://goreportcard.com/report/github.com/Arimeka/mediaprobe)
[![GoDoc](https://godoc.org/github.com/Arimeka/mediaprobe?status.svg)](https://pkg.go.dev/github.com/Arimeka/mediaprobe)

# mediaprobe

Parsing media files using bindings for getting information about codec, bitrate, dimensions, etc.

## Prerequisites

It uses [github.com/rakyll/magicmime](https://github.com/rakyll/magicmime) for detect mimetypes and
[github.com/3d0c/gmf](https://github.com/3d0c/gmf) for parsing audio and video. See these packages for installation info.

## TL;DR Installing on Ubuntu

1. You need go version 1.10 or higher
1. You need `libmagic-dev`
    ```bash
    sudo apt-get install libmagic-dev
    ```
1. You need ffmpeg libraries version 4.x. You may [compile it from sources](https://trac.ffmpeg.org/wiki/CompilationGuide/Ubuntu),
or use PPA. For example, [ppa:jonathonf/ffmpeg-4](https://launchpad.net/~jonathonf/+archive/ubuntu/ffmpeg-4)
    ```bash
    sudo add-apt-repository ppa:jonathonf/ffmpeg-4
    sudo apt-get update
    sudo apt-get install libavcodec-dev libavdevice-dev libavfilter-dev \
                          libavformat-dev libavresample-dev libavutil-dev \
                          libpostproc-dev libswresample-dev libswscale-dev
    ```
## Usage

See [godoc](https://godoc.org/github.com/Arimeka/mediaprobe) examples.

## Benchmark

```
BenchmarkParse-4                             200           7877832 ns/op            2071 B/op         59 allocs/op
BenchmarkNew-4                             10000            123999 ns/op             488 B/op          5 allocs/op
BenchmarkInfo_CalculateMime-4                500           2303269 ns/op              75 B/op          4 allocs/op
BenchmarkInfo_FFProbe-4                      300           5137768 ns/op            2392 B/op         48 allocs/op
BenchmarkInfo_ParseImage-4                  2000            578220 ns/op            1320 B/op         23 allocs/op
```

## Caveats

Detecting mimetype by magic number may not be accurate enough, bindings in Go often cause unexpected memory leaks.
