package permute_test

import (
	"fmt"

	"github.com/cespare/permute/v2"
)

func Example_ints() {
	s := []int{5, 7}
	p := permute.Slice(s)
	for p.Permute() {
		fmt.Println(s)
	}
	// Output:
	// [5 7]
	// [7 5]
}
