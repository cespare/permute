// Package permute implements a generic method for in-place generation
// of all permutations for ordered collections.
package permute

// The algorithm used here is a non-recursive version of Heap's permutation method.
// See http://en.wikipedia.org/wiki/Heap's_algorithm
//
// Here's pseudocode from https://www.cs.princeton.edu/~rs/talks/perms.pdf
//
// generate(int N) {
//   int n;
//   for (n = 1; n <= N; n++) {
//     p[n] = n;
//     c[n] = 1;
//   }
//   doit();
//   for (n = 1; n <= N; ) {
//     if (c[n] < n) {
//       exch(N % 2 ? 1 : c, N)
//       c[n]++;
//       n = 1;
//       doit();
//     } else {
//       c[n++] = 1;
//     }
//   }
// }

// Interface is satisfied by types (usually ordered collections)
// which can have all permutations generated Permute.
type Interface interface {
	Len() int
	Swap(i, j int)
}

// Permute gives a Permuter to generate all permutations
// of the elements of v.
func Permute(v Interface) *Permuter {
	return &Permuter{
		iface:    v,
		started:  false,
		finished: false,
	}
}

// A Permuter holds state about an ongoing iteration of permutations.
type Permuter struct {
	iface    Interface
	started  bool
	finished bool
	n        int
	i        int
	p        []int
	c        []int
}

// Permute generates the next permutation of the contained collection in-place.
// If it returns false, the iteration is finished.
func (p *Permuter) Permute() bool {
	if p.finished {
		panic("Permute() called on finished Permuter")
	}
	if !p.started {
		p.started = true
		p.n = p.iface.Len()
		p.p = make([]int, p.n)
		p.c = make([]int, p.n)
		for i := 0; i < p.n; i++ {
			p.p[i] = i
			p.c[i] = 0
		}
		p.i = 0
		return true
	}
	for {
		if p.i >= p.n {
			p.finished = true
			return false
		}
		if c := p.c[p.i]; c < p.i {
			if p.i&1 == 0 {
				c = 0
			}
			p.iface.Swap(c, p.i)
			p.c[p.i]++
			p.i = 0
			return true
		} else {
			p.c[p.i] = 0
			p.i++
		}
	}
}

type intSlice []int

func (s intSlice) Len() int      { return len(s) }
func (s intSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Ints is a convenience function for generating permutations of []ints.
func Ints(s []int) *Permuter { return Permute(intSlice(s)) }

type stringSlice []string

func (s stringSlice) Len() int      { return len(s) }
func (s stringSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Strings is a convenience function for generating permutations of []strings.
func Strings(s []string) *Permuter { return Permute(stringSlice(s)) }
