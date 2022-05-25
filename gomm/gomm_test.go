package gomm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGomm(t *testing.T) {
	left := []string{
		"1",
		"2",
		"3",
		"5",
		"6",
	}

	right := []string{
		"3",
		"4",
		"6",
		"7",
	}

	l := NewMemoryScanner(left)
	r := NewMemoryScanner(right)

	got := make(map[Position][]string)
	c := func(p Position, line string) {
		got[p] = append(got[p], line)
	}

	g := New(l, r, c)
	err := g.Compare()
	require.NoError(t, err)

	expected := map[Position][]string{
		LEFT: {
			"1",
			"2",
			"5",
		},
		RIGHT: {
			"4",
			"7",
		},
		BOTH: {
			"3",
			"6",
		},
	}

	require.Equal(t, expected, got)
}
