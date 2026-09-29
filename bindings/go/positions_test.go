package tree_sitter_spectra_test

import (
	"testing"

	spectra "github.com/seanfroste/tree-sitter-spectra/bindings/go"
)

// These literal boundaries catch byte/rune confusion, surrogate splitting,
// newline normalization, tab expansion, and accidental EOF clamping.
func TestPositionLiteralBoundaries(t *testing.T) {
	tests := []struct {
		name   string
		source string
		offset uint
		utf8   spectra.Position
		utf16  spectra.Position
	}{
		{"empty", "", 0, spectra.Position{0, 0}, spectra.Position{0, 0}},
		{"start", "a😀β\r\nZ", 0, spectra.Position{0, 0}, spectra.Position{0, 0}},
		{"before supplementary", "a😀β\r\nZ", 1, spectra.Position{0, 1}, spectra.Position{0, 1}},
		{"after supplementary", "a😀β\r\nZ", 5, spectra.Position{0, 5}, spectra.Position{0, 3}},
		{"before CRLF", "a😀β\r\nZ", 7, spectra.Position{0, 7}, spectra.Position{0, 4}},
		{"after CRLF", "a😀β\r\nZ", 9, spectra.Position{1, 0}, spectra.Position{1, 0}},
		{"EOF without newline", "a😀β\r\nZ", 10, spectra.Position{1, 1}, spectra.Position{1, 1}},
		{"tab is one unit", "\t界é\n", 1, spectra.Position{0, 1}, spectra.Position{0, 1}},
		{"three byte rune", "\t界é\n", 4, spectra.Position{0, 4}, spectra.Position{0, 2}},
		{"before combining mark", "\t界é\n", 5, spectra.Position{0, 5}, spectra.Position{0, 3}},
		{"after combining mark", "\t界é\n", 7, spectra.Position{0, 7}, spectra.Position{0, 4}},
		{"EOF after LF", "\t界é\n", 8, spectra.Position{1, 0}, spectra.Position{1, 0}},
		{"empty first line", "\r\n\nX\r\n", 0, spectra.Position{0, 0}, spectra.Position{0, 0}},
		{"empty middle line", "\r\n\nX\r\n", 2, spectra.Position{1, 0}, spectra.Position{1, 0}},
		{"mixed newline content", "\r\n\nX\r\n", 3, spectra.Position{2, 0}, spectra.Position{2, 0}},
		{"EOF after CRLF", "\r\n\nX\r\n", 6, spectra.Position{3, 0}, spectra.Position{3, 0}},
		{"literal replacement rune", "�", 3, spectra.Position{0, 3}, spectra.Position{0, 1}},
		{"embedded NUL", "a\x00b", 2, spectra.Position{0, 2}, spectra.Position{0, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, encoding := range []spectra.PositionEncoding{spectra.UTF8, spectra.UTF16} {
				want := tt.utf8
				if encoding == spectra.UTF16 {
					want = tt.utf16
				}
				got, err := spectra.PositionAt([]byte(tt.source), tt.offset, encoding)
				if err != nil || got != want {
					t.Fatalf("PositionAt(%q, %d, %q) = %v, %v; want %v", tt.source, tt.offset, encoding, got, err, want)
				}
				offset, err := spectra.ByteOffset([]byte(tt.source), want, encoding)
				if err != nil || offset != tt.offset {
					t.Fatalf("ByteOffset(%q, %v, %q) = %d, %v; want %d", tt.source, want, encoding, offset, err, tt.offset)
				}
			}
		})
	}
}

func TestPositionRejectsInvalidOffsets(t *testing.T) {
	for _, encoding := range []spectra.PositionEncoding{spectra.UTF8, spectra.UTF16} {
		for _, offset := range []uint{2, 3, 4, 6, 8, 11, ^uint(0)} {
			if got, err := spectra.PositionAt([]byte("a😀β\r\nZ"), offset, encoding); err == nil {
				t.Errorf("PositionAt offset %d with %q accepted invalid boundary as %v", offset, encoding, got)
			}
		}
		if _, err := spectra.PositionAt(nil, 1, encoding); err == nil {
			t.Errorf("PositionAt empty source with %q accepted offset 1", encoding)
		}
	}
}

func TestPositionRejectsInvalidCharacters(t *testing.T) {
	tests := []struct {
		encoding spectra.PositionEncoding
		position spectra.Position
	}{
		{spectra.UTF8, spectra.Position{0, 2}},
		{spectra.UTF8, spectra.Position{0, 3}},
		{spectra.UTF8, spectra.Position{0, 4}},
		{spectra.UTF8, spectra.Position{0, 6}},
		{spectra.UTF8, spectra.Position{0, 8}},
		{spectra.UTF8, spectra.Position{0, 9}},
		{spectra.UTF16, spectra.Position{0, 2}},
		{spectra.UTF16, spectra.Position{0, 5}},
		{spectra.UTF16, spectra.Position{0, 6}},
	}
	for _, tt := range tests {
		if got, err := spectra.ByteOffset([]byte("a😀β\r\nZ"), tt.position, tt.encoding); err == nil {
			t.Errorf("ByteOffset(%v, %q) accepted invalid character as %d", tt.position, tt.encoding, got)
		}
	}
	for _, encoding := range []spectra.PositionEncoding{spectra.UTF8, spectra.UTF16} {
		for _, position := range []spectra.Position{{1, 2}, {2, 0}, {0, ^uint(0)}, {^uint(0), 0}} {
			if _, err := spectra.ByteOffset([]byte("a😀β\r\nZ"), position, encoding); err == nil {
				t.Errorf("ByteOffset(%v, %q) accepted out-of-bounds position", position, encoding)
			}
		}
		for _, position := range []spectra.Position{{0, 1}, {1, 0}} {
			if _, err := spectra.ByteOffset(nil, position, encoding); err == nil {
				t.Errorf("ByteOffset empty source accepted %v with %q", position, encoding)
			}
		}
	}
}

func TestPositionRejectsInvalidUTF8Anywhere(t *testing.T) {
	// The requested origin is valid by itself. Invalid bytes later in the
	// document must still be rejected, rather than silently truncating validation.
	for _, source := range [][]byte{
		{0xff}, {0x80}, {0xc0, 0xaf}, {0xe2, 0x82},
		{0xed, 0xa0, 0x80}, {0xf4, 0x90, 0x80, 0x80},
		{'a', '\n', 0xff}, {0xff, '\n', 'a'},
	} {
		for _, encoding := range []spectra.PositionEncoding{spectra.UTF8, spectra.UTF16} {
			for _, offset := range []uint{0, uint(len(source))} {
				if _, err := spectra.PositionAt(source, offset, encoding); err == nil {
					t.Errorf("PositionAt(%x, %d, %q) accepted invalid UTF-8", source, offset, encoding)
				}
			}
			for _, position := range []spectra.Position{{0, 0}, {1, 1}} {
				if _, err := spectra.ByteOffset(source, position, encoding); err == nil {
					t.Errorf("ByteOffset(%x, %v, %q) accepted invalid UTF-8", source, position, encoding)
				}
			}
		}
	}
}

func TestPositionRejectsUnsupportedEncoding(t *testing.T) {
	for _, encoding := range []spectra.PositionEncoding{"", "utf-32", "UTF-8"} {
		for _, source := range [][]byte{nil, []byte("abc")} {
			if _, err := spectra.PositionAt(source, 0, encoding); err == nil {
				t.Errorf("PositionAt accepted unsupported encoding %q", encoding)
			}
			if _, err := spectra.ByteOffset(source, spectra.Position{}, encoding); err == nil {
				t.Errorf("ByteOffset accepted unsupported encoding %q", encoding)
			}
		}
	}
}
