package gomm

import (
	"os"
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

func TestGommFile(t *testing.T) {
	contentA := `1
2
3
4
6
7
8
`

	contentB := `3
4
5
6
9
`

	expected := map[Position][]string{
		LEFT: {
			"1\n",
			"2\n",
			"7\n",
			"8\n",
		},
		RIGHT: {
			"5\n",
			"9\n",
		},
		BOTH: {
			"3\n",
			"4\n",
			"6\n",
		},
	}

	fa, err := os.CreateTemp("", "gomm-")
	require.NoError(t, err)
	defer fa.Close()

	n, err := fa.WriteString(contentA)
	require.NoError(t, err)
	require.Equal(t, len(contentA), n)

	fb, err := os.CreateTemp("", "gomm-")
	require.NoError(t, err)
	defer fb.Close()

	n, err = fb.WriteString(contentB)
	require.NoError(t, err)
	require.Equal(t, len(contentB), n)

	l, err := NewFileScanner(fa.Name())
	require.NoError(t, err)
	r, err := NewFileScanner(fb.Name())
	require.NoError(t, err)

	got := make(map[Position][]string)
	c := func(p Position, line string) {
		got[p] = append(got[p], line)
	}

	g := New(l, r, c)
	err = g.Compare()
	require.NoError(t, err)

	require.Equal(t, expected, got)
}
