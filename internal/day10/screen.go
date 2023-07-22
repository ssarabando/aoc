package day10

import "strings"

type screen struct {
	data   []bool
	height int
	width  int
}

func makeScreen(width, height int) *screen {
	return &screen{
		make([]bool, width*height),
		height,
		width,
	}
}

func (s *screen) set(row, col int) {
	if row >= 0 && row < s.height && col >= 0 && col < s.width {
		s.data[s.width*row+col] = true
	}
}

func (s *screen) unset(row, col int) {
	if row >= 0 && row < s.height && col >= 0 && col < s.width {
		s.data[s.width*row+col] = false
	}
}

func (s *screen) String() string {
	m := map[bool]rune{
		true:  '#',
		false: '.',
	}

	result := strings.Builder{}

	for i := 0; i < len(s.data); i++ {
		if i > 0 && i%s.width == 0 {
			result.WriteRune('\n')
		}
		result.WriteRune(m[s.data[i]])
	}
	return result.String()
}
