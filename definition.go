package fascii

import (
	"embed"
	"strings"
)

//go:embed fonts/*
var fontsFS embed.FS

func loadFont(name fontName) (*font, error) {
	fontContent, err := fontsFS.ReadFile("fonts/" + string(name))
	if err != nil {
		return nil, err
	}
	
	return parseFontContent(strings.ReplaceAll(string(fontContent), "\r", ""))
}

type fontName string

const (
	OneRow   fontName = "1Row.flf"
	Standard fontName = "Standard.flf"
)
