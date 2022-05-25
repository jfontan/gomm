package gomm

import (
	"io"
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
