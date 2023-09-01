package magic

type Magic uint32

const (
	// Fat binary magic is referenced  from MacOSX13.3.sdk/usr/include/mach-o/fat.h .

	// Big endian 32 bit magic.
	MagicFat32BE Magic = 0xcafebabe
	// Little endian 32 bit magic.
	MagicFat32LE Magic = 0xbebafeca
	// Big endian 64 bit magic.
	MagicFat64BE Magic = 0xcafebabf
	// Little endian 64 bit magic.
	MagicFat64LE Magic = 0xbfbafeca

	// Mach-O magic is referenced from MacOSX13.3.sdk/usr/include/mach-o/loader.h .

	// Big endian 32 bit magic.
	MagicMacho32BE Magic = 0xfeedface
	// Little endian 32 bit magic.
	MagicMacho32LE Magic = 0xcefaedfe
	// Big endian 64 bit magic.
	MagicMacho64BE Magic = 0xfeedfacf
	// Little endian 64 bit magic.
	MagicMacho64LE Magic = 0xcffaedfe

	// Shebang magic padded with zeros aka #!.
	MagicScript Magic = 0x2321_0000
)
