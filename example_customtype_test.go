package permute_test

import (
	"fmt"

	"github.com/cespare/permute"
)

type ByteSlice []byte

func (s ByteSlice) Len() int      { return len(s) }
func (s ByteSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func Example_customType() {
	b := []byte{'A', 'B', 'C'}
	p := permute.NewPermuter(ByteSlice(b))
	for p.Permute() {
		fmt.Println(string(b))
	}
	// Output:
	// ABC
	// BAC
	// CAB
	// ACB
	// BCA
	// CBA
}
