package mediaprobe

import (
	"strings"

	"github.com/rakyll/magicmime"
)

// CalculateMime calculates mime type by magic numbers
// Function uses libmagic bindings using github.com/rakyll/magicmime package.
func (info *Info) CalculateMime() error {
	if err := magicmime.Open(magicmime.MAGIC_MIME_TYPE | magicmime.MAGIC_SYMLINK | magicmime.MAGIC_ERROR); err != nil {
		return err
	}
	defer magicmime.Close()

	media, err := magicmime.TypeByFile(info.filename)
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
