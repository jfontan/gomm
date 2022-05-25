package examples

import (
	"fmt"

	"github.com/jfontan/gomm/gomm"
)

func ExampleGomm() {
	left := gomm.NewMemoryScanner([]string{"1", "2", "3"})
	right := gomm.NewMemoryScanner([]string{"2", "3", "4"})

	callback := func(p gomm.Position, line string) {
		if p == gomm.BOTH {
			fmt.Println(line)
		}
	}

	g := gomm.New(left, right, callback)
	_ = g.Compare()
	// Output:
	// 2
	// 3
}
