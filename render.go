package fascii

import "strings"

type Renderer struct {
	font *font
}

func NewRendere(name fontName) (*Renderer, error) {
	f, err := loadFont(name)
	if err != nil {
		return nil, err
	}

	return &Renderer{
		font: f,
	}, nil
}

func (r *Renderer) Render(text string) string {
	chars := make([]*asciiChar, len(text))

	for i, char := range text {
		asciiChar, err := newAsciiChar(r.font, char)
		if err != nil {
			asciiChar, _ = newAsciiChar(r.font, '?')
		}

		chars[i] = asciiChar
	}

	result := strings.Builder{}

	for lineIndex := 0; lineIndex < r.font.height; lineIndex++ {
		for i := range chars {
			_, _ = result.WriteString(chars[i].LineAt(lineIndex))
		}
		_, _ = result.WriteString("\n")
	}

	return result.String()
}
