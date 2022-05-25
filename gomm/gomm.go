package gomm

import (
	"errors"
	"io"
)

type Position int

const (
	_ Position = iota
	LEFT
	RIGHT
	BOTH
)

type Callback func(Position, string)

type Gomm struct {
	left  Scanner
	right Scanner

	callback Callback
}

func New(l, r Scanner, c Callback) *Gomm {
	return &Gomm{
		left:     l,
		right:    r,
		callback: c,
	}
}

func (g *Gomm) Compare() error {
	l, err := g.left.Next()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return g.finish(RIGHT, g.right)
		}
		return err
	}

	r, err := g.right.Next()
	if err != nil {
		if errors.Is(err, io.EOF) {
			g.callback(LEFT, l)
			return g.finish(LEFT, g.left)
		}
		return err
	}

	for {
		switch {
		case l == r:
			g.callback(BOTH, l)

			l, err = g.left.Next()
			if err != nil {
				break
			}

			r, err = g.right.Next()
			if err != nil {
				g.callback(LEFT, l)
				break
			}

		case l < r:
			g.callback(LEFT, l)

			l, err = g.left.Next()
			if err != nil {
				break
			}

		case l > r:
			g.callback(RIGHT, r)

			r, err = g.right.Next()
			if err != nil {
				break
			}
		}

		if err != nil {
			break
		}
	}

	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	err = g.finish(LEFT, g.left)
	if err != nil {
		return err
	}
	return g.finish(RIGHT, g.right)
}

func (g *Gomm) finish(d Position, s Scanner) error {
	for {
		l, err := s.Next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		g.callback(d, l)
	}

	return nil
}
