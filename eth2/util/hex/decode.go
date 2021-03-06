package hex

import (
	hexenc "encoding/hex"
	"errors"
)

func DecodeHex(src []byte, dst []byte) error {
	offset, byteCount := DecodeHexOffsetAndLen(src)
	if len(dst) != byteCount {
		return errors.New("cannot decode hex, incorrect length")
	}
	_, err := hexenc.Decode(dst, src[offset:])
	return err
}

func DecodeHexOffsetAndLen(src []byte) (offset int, length int) {
	if src[0] == '0' && src[1] == 'x' {
		offset = 2
	}
	return offset, hexenc.DecodedLen(len(src) - offset)
}