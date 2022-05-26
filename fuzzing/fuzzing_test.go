package fuzzing

import (
	"sort"
	"testing"

	"github.com/jfontan/gomm/gomm"
	"github.com/stretchr/testify/require"
)

func FuzzGomm(f *testing.F) {
	f.Add("one", "two", "three", "four")

	f.Fuzz(func(t *testing.T, a1, a2, b1, b2 string) {
		ll := []string{a1, a2}
		sort.Strings(ll)
		l := gomm.NewMemoryScanner(ll)

		rl := []string{b1, b2}
		sort.Strings(rl)
		r := gomm.NewMemoryScanner(rl)

		var total int
		g := gomm.New(l, r, func(p gomm.Position, s string) {
			total++
			if p == gomm.BOTH {
				total++
			}
		})

		err := g.Compare()
		require.NoError(t, err)
		require.Equal(t, 4, total)
	})

}
