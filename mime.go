package mediaprobe

import (
	"strings"

	"github.com/rakyll/magicmime"
)

func init() {
	if err := magicmime.Open(magicmime.MAGIC_MIME_TYPE | magicmime.MAGIC_SYMLINK | magicmime.MAGIC_ERROR); err != nil {
		panic(err)
	}
}

// CalculateMime calculates mime type by magic numbers
// Function uses libmagic bindings using github.com/rakyll/magicmime package.
func (info *Info) CalculateMime() (err error) {
	var media string
	if info.data != nil {
		media, err = magicmime.TypeByBuffer(info.data)
	} else {
		media, err = magicmime.TypeByFile(info.filename)
	}
	if err != nil {
		return err
	}

	info.MediaType = "application"
	info.MediaSubtype = "octet-stream"

	kind := strings.Split(media, "/")
	if len(kind) == 2 {
		info.MediaType = kind[0]
		info.MediaSubtype = kind[1]
	}

	return nil
}
