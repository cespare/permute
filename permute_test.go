package permute

import (
	"reflect"
	"strings"
	"testing"
)

func TestSimple(t *testing.T) {
	for _, tt := range []struct {
		in  string
		out []string
	}{
		{"", []string{""}},
		{"A", []string{"A"}},
		{"AB", []string{"AB", "BA"}},
		{"ABC", []string{"ABC", "BAC", "CAB", "ACB", "BCA", "CBA"}},
	} {
		want := make(map[string]struct{})
		for _, w := range tt.out {
			want[w] = struct{}{}
		}
		s := strings.Split(tt.in, "")
		p := Strings(s)
		got := make(map[string]struct{})
		for p.Permute() {
			got[strings.Join(s, "")] = struct{}{}
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v; want %v", got, want)
		}
	}
}

func fact(n int) int {
	m := 1
	for ; n > 1; n-- {
		m *= n
	}
	return m
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
		p := NewPermuter(byteSlice(s))
		perms := make(map[string]struct{})
		for p.Permute() {
			perms[string(s)] = struct{}{}
		}
		if got, want := len(perms), fact(n); got != want {
			t.Fatalf("fail at n=%d: got %d; want %d", n, got, want)
		}
	}
}
