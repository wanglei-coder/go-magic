package magic

import (
	"os"
	"testing"
)

type TestMagic struct {
	mimeType   string
	descriptor string
}

func TestNewMagic(t *testing.T) {

	tests := []struct {
		in  string
		out TestMagic
	}{
		{
			"test/sample.elf",
			TestMagic{mimeType: "application/x-sharedlib", descriptor: "ELF 64-bit LSB shared object, x86-64, version 1 (SYSV), dynamically linked, interpreter /lib64/ld-linux-x86-64.so.2, for GNU/Linux 2.6.32, BuildID[sha1]=c828586e6e7cf929500a5b9c04faece9eceed5cc, stripped"},
		},
		{
			"test/sample.jpg",
			TestMagic{mimeType: "image/jpeg", descriptor: `JPEG image data, JFIF standard 1.01, aspect ratio, density 1x1, segment length 16, comment: "CREATOR: gd-jpeg v1.0 (using IJG JPEG v62), quality = 95", baseline, precision 8, 2560x1440, components 3`},
		},
		{
			"test/sample.pdf",
			TestMagic{mimeType: "application/pdf", descriptor: `PDF document, version 1.7`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			m, err := NewMagic(MagicMimeType | MagicSymLink | MagicError)
			if err != nil {
				t.Fatalf("NewMagic failed, reason: %v", err)
			}
			defer m.Close()
			{
				got, err := m.FromFile(tt.in)
				if err != nil {
					t.Errorf("mimeType(%s) got %s, want %s", tt.in, got, tt.out.mimeType)
				}
			}

			{
				m.SetFlags(MagicSymLink | MagicError)
				got, err := m.FromFile(tt.in)
				if err != nil {
					t.Errorf("descriptor(%s) got %s, want %s", tt.in, got, tt.out.descriptor)
				}
			}

			{
				f, err := os.Open(tt.in)
				if err != nil {
					t.Error(err)
				}
				defer f.Close()

				got, err := m.FromDescriptor(int(f.Fd()))
				if err != nil {
					t.Errorf("mimeType(%s) got %s, want %s", tt.in, got, tt.out.descriptor)
				}
			}

		})
	}
}
