package fascii

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//
// Explanation of the .flf file header
// THE HEADER LINE
//
// The header line gives information about the FIGfont.  Here is an example
// showing the names of all parameters:
//
//           flf2a$ 6 5 20 15 3 0 143 229    NOTE: The first five characters in
//             |  | | | |  |  | |  |   |     the entire file must be "flf2a".
//            /  /  | | |  |  | |  |   \
//   Signature  /  /  | |  |  | |   \   Codetag_Count
//     Hardblank  /  /  |  |  |  \   Full_Layout*
//          Height  /   |  |   \  Print_Direction
//          Baseline   /    \   Comment_Lines
//           Max_Length      Old_Layout*
//
//   * The two layout parameters are closely related and fairly complex.
//       (See "INTERPRETATION OF LAYOUT PARAMETERS".)
//

type font struct {
	hardblank string
	height    int
	fontChars []string
}

func (f *font) charLines(char rune) []string {
	height := f.height
	beginRow := (int(char) - 32) * height

	lines := make([]string, height)

	for i := 0; i < height; i++ {
		row := f.fontChars[beginRow+i]
		row = strings.Replace(row, "@", "", -1)
		row = strings.Replace(row, f.hardblank, " ", -1)
		lines[i] = row
	}

	return lines
}

func parseFontContent(fontContent string) (*font, error) {
	lines := strings.Split(fontContent, "\n")

	if len(lines) < 1 {
		return nil, errors.New("empty font content")
	}

	header := strings.Split(lines[0], " ")

	commentLines, err := strconv.Atoi(header[5])
	if err != nil {
		return nil, fmt.Errorf("failed to parse header comment lines parameter: %w", err)
	}

	height, err := strconv.Atoi(header[1])
	if err != nil {
		return nil, fmt.Errorf("failed to parse header height paremeter: %w", err)
	}

	font := &font{
		hardblank: header[0][len(header[0])-1:],
		height:    height,
		fontChars: lines[commentLines+1:],
	}

	return font, nil
}

