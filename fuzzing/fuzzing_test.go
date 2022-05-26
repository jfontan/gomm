package fuzzing

import (
	"testing"

	"github.com/jfontan/gomm/gomm"
	"github.com/stretchr/testify/require"
)

func FuzzGomm(f *testing.F) {
	f.Add("one", "two", "three", "four")

	f.Fuzz(func(t *testing.T, a1, a2, b1, b2 string) {
		l := gomm.NewMemoryScanner([]string{a1, a2})
		r := gomm.NewMemoryScanner([]string{b1, b2})

		var total int
		g := gomm.New(l, r, func(p gomm.Position, s string) {
			total++
		})

		err := g.Compare()
		require.NoError(t, err)
		require.Equal(t, 4, total)
	})

}
