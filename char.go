package fascii

import (
	"errors"
)

type asciiChar struct {
	lines []string
}

func newAsciiChar(font *font, char rune) (*asciiChar, error) {
	if char < 0 || char > 127 {
		return nil, errors.New("not ASCII char")
	}

	lines := font.charLines(char)
	return &asciiChar{lines: lines}, nil
}

func (char *asciiChar) LineAt(index int) string {
	return char.lines[index]
}
