package gomm

import "io"

type Scanner interface {
	Next() (string, error)
}

type MemoryScanner struct {
	lines []string
	pos   int
}

func NewMemoryScanner(lines []string) *MemoryScanner {
	return &MemoryScanner{
		lines: lines,
		pos:   0,
	}
}

func (s *MemoryScanner) Next() (string, error) {
	if s.pos >= len(s.lines) {
		return "", io.EOF
	}

	l := s.lines[s.pos]
	s.pos++

	return l, nil
}
