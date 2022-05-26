package gomm

import (
	"errors"
	"io"
)

type Position int

func (p Position) String() string {
	switch p {
	case LEFT:
		return "left"
	case RIGHT:
		return "right"
	case BOTH:
		return "both"

	default:
		return "unknown"
	}
}

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

	loop:
	for {
		switch {
		case l == r:
			g.callback(BOTH, l)

			l, err = g.left.Next()
			if err != nil {
				break loop
			}

			r, err = g.right.Next()
			if err != nil {
				g.callback(LEFT, l)
				break loop
			}

		case l < r:
			g.callback(LEFT, l)

			l, err = g.left.Next()
			if err != nil {
				break loop
			}

		case l > r:
			g.callback(RIGHT, r)

			r, err = g.right.Next()
			if err != nil {
				break loop
			}
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
