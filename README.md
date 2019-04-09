# mediaprobe

Parsing media files using bindings for getting information about codec, bitrate, dimensions, etc.

## Prerequisites

It uses [github.com/rakyll/magicmime](https://github.com/rakyll/magicmime) for detect mimetypes and
[github.com/3d0c/gmf](https://github.com/3d0c/gmf) for parsing audio and video. See these packages for installation info.

## Usage

See example folder.

## Benchmark

```
BenchmarkParse-4                     200           7433064 ns/op            2069 B/op         59 allocs/op
BenchmarkInfo_FFProbe-4              300           4776643 ns/op            2392 B/op         48 allocs/op
BenchmarkInfo_ParseImage-4          3000            465799 ns/op            1488 B/op         23 allocs/op
```

## Caveats

Detecting mimetype by magic number may not be accurate enough, bindings in Go often cause unexpected memory leaks.
