package gomm

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMemoryScanner(t *testing.T) {
	expected := []string{
		"one",
		"two",
		"three",
		"four",
	}

	s := NewMemoryScanner(expected)

	var got []string
	var err error
	var line string
	for {
		line, err = s.Next()
		if line != "" {
			require.NoError(t, err)
		} else {
			require.ErrorIs(t, err, io.EOF)
			break
		}

		got = append(got, line)
	}

	require.Equal(t, expected, got)
}

var ()

func TestFileScanner(t *testing.T) {
	contents := `1
2
3
4
5
6
7
8
9
`

	expected := []string{
		"1\n",
		"2\n",
		"3\n",
		"4\n",
		"5\n",
		"6\n",
		"7\n",
		"8\n",
		"9\n",
	}

	f, err := os.CreateTemp("", "goom-")
	require.NoError(t, err)
	defer f.Close()

	n, err := f.WriteString(contents)
	require.NoError(t, err)
	require.Equal(t, len(contents), n)

	s, err := NewFileScanner(f.Name())
	require.NoError(t, err)

	var got []string
	var line string
	for {
		line, err = s.Next()
		if line != "" {
			require.NoError(t, err)
		} else {
			require.ErrorIs(t, err, io.EOF)
			break
		}

		got = append(got, line)
	}

	require.Equal(t, expected, got)
}
