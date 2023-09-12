package magic

import (
	"encoding/binary"
	"errors"
	"os"
)

var ErrUnknownMagic = errors.New("magic: unknown magic value")

func IsAnExecutable(path string) (bool, error) {
	magicFd, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer magicFd.Close()

	const magicLen int = 4
	magic := make([]byte, magicLen)

	if _, err := magicFd.Read(magic); err != nil {
		return false, err
	}

	magicNumShort := Magic(binary.BigEndian.Uint32(magic))

	switch magicNumShort {
	case MagicFat32BE, MagicFat32LE, MagicFat64BE, MagicFat64LE,
		MagicMacho32BE, MagicMacho32LE, MagicMacho64BE, MagicMacho64LE,
		MagicScript:
		return true, nil
	default:
		// Set the last bytes to zero, to check for a shebang,
		// whose magic is only 2 bytes long.
		magicNumShort &= 0xFFFF0000
		if magicNumShort == MagicScript {
			return true, nil
		}
	}

	return false, ErrUnknownMagic
}
