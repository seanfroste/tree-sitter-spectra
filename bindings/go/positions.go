package tree_sitter_spectra

import (
	"errors"
	"unicode/utf8"
)

// PositionEncoding identifies the units in a Position's Character field.
// Its values match the corresponding LSP position-encoding names.
type PositionEncoding string

const (
	UTF8  PositionEncoding = "utf-8"
	UTF16 PositionEncoding = "utf-16"
)

// Position is a zero-based line and character boundary. Character counts bytes
// for UTF8 and UTF-16 code units for UTF16, not display columns or graphemes.
// Tabs count as one unit. LF and CRLF end lines; their bytes do not contribute
// to Character. A lone CR is an ordinary character.
type Position struct {
	Line      uint
	Character uint
}

var (
	positionEncodingError = errors.New("unsupported position encoding")
	positionUTF8Error     = errors.New("source contains invalid UTF-8")
	positionBoundaryError = errors.New("position splits an encoded character or CRLF")
	positionBoundsError   = errors.New("position is outside the source")
)

// PositionAt converts a byte offset into a position without copying source.
// EOF and the boundary immediately before a line terminator are valid. Offsets
// inside a rune or between CR and LF are invalid. The entire source must be
// valid UTF-8, including the portion after offset.
func PositionAt(source []byte, offset uint, encoding PositionEncoding) (Position, error) {
	if encoding != UTF8 && encoding != UTF16 {
		return Position{}, positionEncodingError
	}
	if offset > uint(len(source)) {
		return Position{}, positionBoundsError
	}
	var position Position
	for i := uint(0); i < offset; {
		next, units, newline, err := positionStep(source, i, encoding)
		if err != nil {
			return Position{}, err
		}
		if next > offset {
			return Position{}, positionBoundaryError
		}
		if newline {
			position.Line++
			position.Character = 0
		} else {
			position.Character += units
		}
		i = next
	}
	if !utf8.Valid(source[offset:]) {
		return Position{}, positionUTF8Error
	}
	return position, nil
}

// ByteOffset converts a position to a byte offset without copying source.
// It rejects out-of-range lines/characters, UTF-8 rune and UTF-16 surrogate
// splits, and positions extending into a line terminator. It never clamps a
// position. The entire source must be valid UTF-8; EOF is a valid boundary.
func ByteOffset(source []byte, position Position, encoding PositionEncoding) (uint, error) {
	if encoding != UTF8 && encoding != UTF16 {
		return 0, positionEncodingError
	}
	var current Position
	for i := uint(0); ; {
		if current == position {
			if !utf8.Valid(source[i:]) {
				return 0, positionUTF8Error
			}
			return i, nil
		}
		if current.Line == position.Line && current.Character > position.Character {
			return 0, positionBoundaryError
		}
		if i == uint(len(source)) {
			return 0, positionBoundsError
		}
		next, units, newline, err := positionStep(source, i, encoding)
		if err != nil {
			return 0, err
		}
		if newline {
			if current.Line == position.Line {
				return 0, positionBoundsError
			}
			current.Line++
			current.Character = 0
		} else {
			current.Character += units
		}
		i = next
	}
}

// positionStep consumes one rune or one complete line terminator. Callers only
// use it while offset < len(source), and validate the unvisited suffix once
// they find their answer, avoiding both a copy and a second prefix scan.
func positionStep(source []byte, offset uint, encoding PositionEncoding) (next, units uint, newline bool, err error) {
	b := source[offset]
	if b == '\n' {
		return offset + 1, 0, true, nil
	}
	if b == '\r' && offset+1 < uint(len(source)) && source[offset+1] == '\n' {
		return offset + 2, 0, true, nil
	}
	if b < utf8.RuneSelf {
		return offset + 1, 1, false, nil
	}
	r, size := utf8.DecodeRune(source[offset:])
	if r == utf8.RuneError && size == 1 {
		return 0, 0, false, positionUTF8Error
	}
	units = uint(size)
	if encoding == UTF16 {
		units = 1
		if r > 0xffff {
			units = 2
		}
	}
	return offset + uint(size), units, false, nil
}
