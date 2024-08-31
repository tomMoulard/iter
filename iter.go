// Package iter provides functions to work with sequences.
package iter

import (
	"cmp"
	"iter"
	"slices"
)

// Values returns a slice of elements from the input sequence.
func Values[T any](seq iter.Seq[T]) []T {
	var res []T

	seq(func(elem T) bool {
		res = append(res, elem)

		return true
	})

	return res
}

// Values2 returns a pair of slices of elements from the input sequence.
func Values2[T, U any](seq iter.Seq2[T, U]) ([]T, []U) {
	var (
		resT []T
		resU []U
	)

	seq(func(elem1 T, elem2 U) bool {
		resT = append(resT, elem1)
		resU = append(resU, elem2)

		return true
	})

	return resT, resU
}

// Values2Map returns a map of elements from the input sequence.
func Values2Map[T comparable, U any](seq iter.Seq2[T, U]) map[T]U {
	res := make(map[T]U)

	seq(func(elem1 T, elem2 U) bool {
		res[elem1] = elem2

		return true
	})

	return res
}

// Zip returns a sequence of pairs of elements from the input sequences.
// The resulting sequence is as long as the shortest input sequence.
func Zip[T, U any](a []T, b []U) iter.Seq2[T, U] {
	minLen := min(len(a), len(b))

	return func(yield func(T, U) bool) {
		for i := 0; i < minLen && yield(a[i], b[i]); i++ {
		}
	}
}

// IZip returns a sequence of pairs of elements from the input sequences.
// The resulting sequence is as long as the shortest input sequence.
func IZip[T, U any](seqA iter.Seq[T], seqB iter.Seq[U]) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		next, stop := iter.Pull(seqB)
		defer stop()

		b, ok2 := next()
		for a := range seqA {
			if !ok2 || !yield(a, b) {
				return
			}

			b, ok2 = next()
		}
	}
}

type either[T, U any] interface{}

// ZipLongest returns a sequence of pairs of elements from the input
// sequences.
// The resulting sequence is as long as the longest input sequence.
// If one sequence is shorter than the other, the missing values are filled
// with the provided fill value.
func ZipLongest[T, U any](a []T, b []U, fill either[T, U]) iter.Seq2[T, U] {
	maxLen := max(len(a), len(b))

	return func(yield func(T, U) bool) {
		for i := range maxLen {
			switch {
			case i < len(a) && i < len(b):
				if !yield(a[i], b[i]) {
					return
				}

			case i < len(a):
				fillU, ok := fill.(U)
				if !ok {
					return
				}

				if !yield(a[i], fillU) {
					return
				}

			case i < len(b):
				fillT, ok := fill.(T)
				if !ok {
					return
				}

				if !yield(fillT, b[i]) {
					return
				}
			}
		}
	}
}

// Accumulate returns a sequence of accumulated values.
// The first element is the same as the first element of the input sequence.
// The second element is the sum of the first and second elements of the input
// sequence.
// So on and so forth.
func Accumulate[T cmp.Ordered](a []T) iter.Seq[T] {
	if len(a) == 0 {
		return func(yield func(T) bool) {}
	}

	return func(yield func(T) bool) {
		acc := a[0]
		if !yield(acc) {
			return
		}

		for i := 1; i < len(a); i++ {
			acc += a[i]
			if !yield(acc) {
				return
			}
		}
	}
}

// Chain returns a sequence of elements from the input sequences.
// The resulting sequence is the concatenation of the input sequences.
func Chain[T any](seqs ...[]T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range seqs {
			for j := range seqs[i] {
				if !yield(seqs[i][j]) {
					return
				}
			}
		}
	}
}

// ChainSeq returns a sequence of elements from the input sequences.
// The resulting sequence is the concatenation of the input sequences.
func ChainSeq[T any](seqs ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range seqs {
			seqs[i](func(elem T) bool {
				return yield(elem)
			})
		}
	}
}

// Compress returns a sequence of elements from the input sequence.
// The resulting sequence contains only the elements where the corresponding
// selector is true.
func Compress[T any](data []T, selectors []bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range selectors {
			if selectors[i] && !yield(data[i]) {
				return
			}
		}
	}
}

// DropWhile returns a sequence of elements from the input sequence.
// The resulting sequence contains only the elements after the predicate is
// false.
func DropWhile[T any](pred func(T) bool, a []T) iter.Seq[T] {
	var i int
	for i = 0; i < len(a) && pred(a[i]); i++ {
	}

	return func(yield func(T) bool) {
		for ; i < len(a) && yield(a[i]); i++ {
		}
	}
}

// Filter returns a sequence of elements from the input sequence.
// The resulting sequence contains only the elements where the predicate is true.
func Filter[T any](pred func(T) bool, a []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range a {
			if pred(a[i]) && !yield(a[i]) {
				return
			}
		}
	}
}

// FilterFalse returns a sequence of elements from the input sequence.
// The resulting sequence contains only the elements where the predicate is false.
func FilterFalse[T any](pred func(T) bool, a []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range len(a) {
			if !pred(a[i]) && !yield(a[i]) {
				return
			}
		}
	}
}

// GroupBy returns a sequence of groups of elements from the input sequence.
// The resulting sequence contains groups of elements where the key function
// returns the same value.
func GroupBy[T any, K comparable](key func(T) K, a []T) iter.Seq2[K, iter.Seq[T]] {
	groups := make(map[K][]T)

	for i := range a {
		k := key(a[i])
		groups[k] = append(groups[k], a[i])
	}

	return func(yield func(K, iter.Seq[T]) bool) {
		for k, group := range groups {
			if !yield(k, slices.Values(group)) {
				return
			}
		}
	}
}

// Map returns a sequence of elements from the input sequence.
// The resulting sequence contains the elements after applying the function to
// each element of the input sequence.
func Map[T, U any](f func(T) U, a []T) iter.Seq[U] {
	return func(yield func(U) bool) {
		for i := 0; i < len(a) && yield(f(a[i])); i++ {
		}
	}
}

// Map2 returns a sequence of elements from the input sequences.
// The resulting sequence contains the elements after applying the function to
// each pair of elements from the input sequences.
func Map2[T, U, V any](f func(T, U) V, a []T, b []U) iter.Seq[V] {
	minLen := min(len(a), len(b))

	return func(yield func(V) bool) {
		for i := 0; i < minLen && yield(f(a[i], b[i])); i++ {
		}
	}
}

// TakeWhile returns a sequence of elements from the input sequence.
// The resulting sequence contains only the elements before the predicate is
// false.
func TakeWhile[T any](pred func(T) bool, a []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := 0; i < len(a) && pred(a[i]) && yield(a[i]); i++ {
		}
	}
}

type pair[T, U any] struct {
	key T
	val U
}

// ChainMap returns a sequence of elements from the input map.
// The resulting sequence is ordered by the keys of the input map.
func ChainMap[T cmp.Ordered, U any](maps ...map[T]U) iter.Seq2[T, U] {
	ordered := make([]pair[T, U], 0, len(maps))

	for i := range maps {
		for k, v := range maps[i] {
			ordered = append(ordered, pair[T, U]{key: k, val: v})
		}
	}

	slices.SortFunc(ordered, func(a, b pair[T, U]) int {
		return cmp.Compare(a.key, b.key)
	})

	return func(yield func(T, U) bool) {
		for _, v := range ordered {
			if !yield(v.key, v.val) {
				return
			}
		}
	}
}

// Permutations returns a sequence of permutations of the input sequence.
// The resulting sequence contains all possible permutations of the input
// sequence.
// This function is a helper for `PermutationsLen(a, len(a))`.
func Permutations[T any](a []T) iter.Seq[[]T] {
	return PermutationsLen(a, len(a))
}

// PermutationsLen returns a sequence of permutations of the input sequence.
// The resulting sequence contains all possible permutations of the input
// sequence with the specified length.
// It uses the Heap's algorithm to generate the permutations.
func PermutationsLen[T any](a []T, length int) iter.Seq[[]T] {
	if length == 0 {
		return func(yield func([]T) bool) {}
	}

	if length > len(a) {
		return func(yield func([]T) bool) {
			if !yield(a) {
				return
			}
		}
	}

	if length == 1 {
		return func(yield func([]T) bool) {
			for i := range a {
				if !yield([]T{a[i]}) {
					return
				}
			}
		}
	}

	c := make([]int, len(a))
	i := 0

	return func(yield func([]T) bool) {
		if !yield(a[:length]) {
			return
		}

		for i < len(a) {
			if c[i] >= i {
				c[i] = 0
				i++

				continue
			}

			if i%2 == 0 {
				a[i], a[0] = a[0], a[i]
			} else {
				a[i], a[c[i]] = a[c[i]], a[i]
			}

			if !yield(a[:length]) {
				return
			}

			c[i]++
			i = 0
		}
	}
}

// Append returns a sequence of elements from the concatenation of the input
// sequences.
func Append[T any](seqs ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range seqs {
			seqs[i](func(elem T) bool {
				return yield(elem)
			})
		}
	}
}

// Append2 returns a sequence of elements from the concatenation of the input
// sequences.
func Append2[T, U any](seqs ...iter.Seq2[T, U]) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		for i := range seqs {
			seqs[i](func(elem1 T, elem2 U) bool {
				return yield(elem1, elem2)
			})
		}
	}
}

// First returns a sequence of elements from the input sequence.
// The resulting sequence contains only the first element of the input sequence.
func First[T, U any](seq iter.Seq2[T, U]) iter.Seq[T] {
	return func(yield func(T) bool) {
		seq(func(elem T, _ U) bool {
			return yield(elem)
		})
	}
}

// Second returns a sequence of elements from the input sequence.
// The resulting sequence contains only the second element of the input sequence.
func Second[T, U any](seq iter.Seq2[T, U]) iter.Seq[U] {
	return func(yield func(U) bool) {
		seq(func(_ T, elem U) bool {
			return yield(elem)
		})
	}
}

func Equal[T comparable](seqA, seqB iter.Seq[T]) bool {
	for aa, bb := range IZip(seqA, seqB) {
		if aa != bb {
			return false
		}
	}

	return true
}
