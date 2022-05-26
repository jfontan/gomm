package gomm

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

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

type FileScanner struct {
	f   *os.File
	b   *bufio.Reader
	eof bool
}

func NewFileScanner(file string) (*FileScanner, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("could not open file %s: %w", file, err)
	}

	b := bufio.NewReader(f)

	return &FileScanner{
		f: f,
		b: b,
	}, nil
}

func (s *FileScanner) Next() (string, error) {
	if s.eof {
		return "", io.EOF
	}

	line, err := s.b.ReadString('\n')
	if err != nil {
		if errors.Is(err, io.EOF) {
			s.eof = true
			if len(line) > 0 {
				return line, nil
			}
			_ = s.f.Close()
			return "", io.EOF
		}
		_ = s.f.Close()
		return "", fmt.Errorf("could not read file: %w", err)
	}
	return line, nil
}
