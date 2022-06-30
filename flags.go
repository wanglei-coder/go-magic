package magic

// #cgo CFLAGS: -I/usr/local/include
// #cgo LDFLAGS: -lmagic -L/usr/local/lib
// #include <stdlib.h>
// #include <magic.h>
import "C"

type Flag int

const (
	MagicNone            Flag = C.MAGIC_NONE              // No flags
	MagicDebug           Flag = C.MAGIC_DEBUG             // Turn on debugging
	MagicSymLink         Flag = C.MAGIC_SYMLINK           // Follow symlinks
	MagicCompress        Flag = C.MAGIC_COMPRESS          // Check inside compressed files
	MagicDevices         Flag = C.MAGIC_DEVICES           // Look at the contents of devices
	MagicMimeType        Flag = C.MAGIC_MIME_TYPE         // Return the MIME type
	MagicContinue        Flag = C.MAGIC_CONTINUE          // Return all matches
	MagicCheck           Flag = C.MAGIC_CHECK             // Print warnings to stderr
	MagicPreserveATime   Flag = C.MAGIC_PRESERVE_ATIME    // Restore access time on exit
	MagicRaw             Flag = C.MAGIC_RAW               // Don't convert unprintable chars
	MagicError           Flag = C.MAGIC_ERROR             // Handle ENOENT etc as real errors
	MagicMimeEncoding    Flag = C.MAGIC_MIME_ENCODING     // Return the MIME encoding
	MagicMime            Flag = C.MAGIC_MIME              // MAGIC_MIME_TYPE | MAGIC_MIME_ENCODING
	MagicApple           Flag = C.MAGIC_APPLE             // Return the Apple creator/type
	MagicExtension       Flag = C.MAGIC_EXTENSION         // Return a /-separated list of extensions
	MagicCompressTRANSP  Flag = C.MAGIC_COMPRESS_TRANSP   // Check inside compressed files but not report compression
	MagicNoDesc          Flag = C.MAGIC_NODESC            // MAGIC_EXTENSION|MAGIC_MIME|MAGIC_APPLE
	MagicNoCheckCompress Flag = C.MAGIC_NO_CHECK_COMPRESS // Don't check for compressed files
	MagicNoCheckTar      Flag = C.MAGIC_NO_CHECK_TAR      // Don't check for tar files
	MagicNoCheckSoft     Flag = C.MAGIC_NO_CHECK_SOFT     // Don't check magic entries
	MagicNoCheckAppType  Flag = C.MAGIC_NO_CHECK_APPTYPE  // Don't check application type
	MagicNoCheckELF      Flag = C.MAGIC_NO_CHECK_ELF      // Don't check for elf details
	MagicNoCheckText     Flag = C.MAGIC_NO_CHECK_TEXT     // Don't check for text files
	MagicNoCheckCDF      Flag = C.MAGIC_NO_CHECK_CDF      // Don't check for cdf files
	MagicNoCheckCSV      Flag = C.MAGIC_NO_CHECK_CSV      // Don't check for CSV files
	MagicNoCheckTokens   Flag = C.MAGIC_NO_CHECK_TOKENS   // Don't check tokens
	MagicNoCheckEncoding Flag = C.MAGIC_NO_CHECK_ENCODING // Don't check text encodings
	MagicNoCheckJSON     Flag = C.MAGIC_NO_CHECK_JSON     // Don't check for JSON files
	MagicNoCheckASCII    Flag = C.MAGIC_NO_CHECK_ASCII
	MagicNoCheckFortran  Flag = C.MAGIC_NO_CHECK_FORTRAN // Don't check ascii/fortran
	MagicNoCheckTROFF    Flag = C.MAGIC_NO_CHECK_TROFF   // Don't check ascii/troff
)
