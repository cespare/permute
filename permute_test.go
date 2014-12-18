package permute

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func ExamplePermute() {
	s := []string{"A", "B", "C"}
	p := Strings(s)
	for p.Permute() {
		fmt.Println(strings.Join(s, ""))
	}
	// Output:
	// ABC
	// BAC
	// CAB
	// ACB
	// BCA
	// CBA
}

func TestSimple(t *testing.T) {
	for _, tc := range []struct {
		in  string
		out []string
	}{
		{"", []string{""}},
		{"A", []string{"A"}},
		{"AB", []string{"AB", "BA"}},
		{"ABC", []string{"ABC", "BAC", "CAB", "ACB", "BCA", "CBA"}},
	} {
		want := make(map[string]struct{})
		for _, w := range tc.out {
			want[w] = struct{}{}
		}
		s := strings.Split(tc.in, "")
		p := Strings(s)
		got := make(map[string]struct{})
		for p.Permute() {
			got[strings.Join(s, "")] = struct{}{}
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v; got: %v", want, got)
		}
	}
}

func fact(n int) int {
	result := 1
	for ; n > 1; n-- {
		result *= n
	}
	return result
}

type byteSlice []byte

func (s byteSlice) Len() int      { return len(s) }
func (s byteSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func TestBrute(t *testing.T) {
	for n := 0; n <= 9; n++ {
		s := make([]byte, n)
		for i := range s {
			s[i] = byte(i) + 'a'
		}
		p := Permute(byteSlice(s))
		perms := make(map[string]struct{})
		for p.Permute() {
			perms[string(s)] = struct{}{}
		}
		want := fact(n)
		got := len(perms)
		if want != got {
			t.Errorf("fail on n=%d: want %d; got %d", n, want, got)
		}
	}
}
