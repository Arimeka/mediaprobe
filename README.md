# mediaprobe

Parsing media files using binding for getting information about codec, bitrate, dimensions, etc.

## Prerequisites

It uses [github.com/rakyll/magicmime](https://github.com/rakyll/magicmime) for detect mimetypes and
[github.com/3d0c/gmf](https://github.com/3d0c/gmf) for parsing audio and video. See these packages for installation info.

## Usage

See example folder.

## Benchmark

With `-benchtime 1m`

```
BenchmarkParse-4                   10000           7602969 ns/op            2058 B/op         59 allocs/op
BenchmarkInfo_FFProbe-4            20000           5010409 ns/op            2941 B/op         48 allocs/op
BenchmarkInfo_FFProbeCli-4          3000          30534269 ns/op           20197 B/op        163 allocs/op
BenchmarkInfo_ParseImage-4        200000            515128 ns/op            1548 B/op         23 allocs/op
```

## Caveats

Detecting mimetype by magic number may not be accurate enough, binding in Go often cause unexpected memory leaks,
exif-rotation ignored (I hate exif-rotation).
