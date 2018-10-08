package main

import (
	"fmt"
	"github.com/theodesp/bisect"
)

func main() {
	var ints bisect.IntSlice
	for i := 0; i < 10; i += 2 {
		ints = append(ints, i)
	}

	ints.InsortLeft(-1)
	ints.InsortRight(16)

	fmt.Println(ints) // [16 -1 0 2 4 6 8]
}
