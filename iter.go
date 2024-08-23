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

// Zip returns a sequence of pairs of elements from the input sequences.
// The resulting sequence is as long as the shortest input sequence.
func Zip[T, U any](a []T, b []U) iter.Seq2[T, U] {
	minLen := min(len(a), len(b))

	return func(yield func(T, U) bool) {
		for i := 0; i < minLen && yield(a[i], b[i]); i++ {
		}
	}
}

type Either[T, U any] interface{}

// ZipLongest returns a sequence of pairs of elements from the input
// sequences.
// The resulting sequence is as long as the longest input sequence.
// If one sequence is shorter than the other, the missing values are filled
// with the provided fill value.
func ZipLongest[T, U any](a []T, b []U, fill Either[T, U]) iter.Seq2[T, U] {
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
