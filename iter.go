package iter

import (
	"cmp"
	stdIter "iter"
	"slices"
)

// Values returns a slice of elements from the input sequence.
func Values[T any](seq stdIter.Seq[T]) []T {
	var res []T
	seq(func(elem T) bool {
		res = append(res, elem)
		return true
	})
	return res
}

// Values2 returns a pair of slices of elements from the input sequence.
func Values2[T, U any](seq stdIter.Seq2[T, U]) ([]T, []U) {
	var resT []T
	var resU []U
	seq(func(elem1 T, elem2 U) bool {
		resT = append(resT, elem1)
		resU = append(resU, elem2)
		return true
	})
	return resT, resU
}

// Zip returns a sequence of pairs of elements from the input sequences.
// The resulting sequence is as long as the shortest input sequence.
func Zip[T, U any](a []T, b []U) stdIter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		minLen := min(len(a), len(b))
		for i := 0; i < minLen; i++ {
			if !yield(a[i], b[i]) {
				return
			}
		}
	}
}

// ZipLongest returns a sequence of pairs of elements from the input
// sequences.
// The resulting sequence is as long as the longest input sequence.
// If one sequence is shorter than the other, the missing values are filled
// with the provided fill value.
func ZipLongest[T, U any](a []T, b []U, fill any) stdIter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		maxLen := max(len(a), len(b))
		for i := 0; i < maxLen; i++ {
			if i < len(a) && i < len(b) {
				if !yield(a[i], b[i]) {
					return
				}
			} else if i < len(a) {
				if !yield(a[i], fill.(U)) {
					return
				}
			} else if i < len(b) {
				if !yield(fill.(T), b[i]) {
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
func Accumulate[T cmp.Ordered](a []T) stdIter.Seq[T] {
	return func(yield func(T) bool) {
		if len(a) == 0 {
			return
		}

		acc := a[0]
		if !yield(acc) {
			return
		}

		for i := 1; i < len(a); i++ {
			acc = acc + a[i]
			if !yield(acc) {
				return
			}
		}
	}
}

// Chain returns a sequence of elements from the input sequences.
// The resulting sequence is the concatenation of the input sequences.
func Chain[T any](seqs ...[]T) stdIter.Seq[T] {
	return func(yield func(T) bool) {
		for _, seq := range seqs {
			for _, elem := range seq {
				if !yield(elem) {
					return
				}
			}
		}
	}
}

// ChainSeq returns a sequence of elements from the input sequences.
// The resulting sequence is the concatenation of the input sequences.
func ChainSeq[T any](seqs ...stdIter.Seq[T]) stdIter.Seq[T] {
	return func(yield func(T) bool) {
		for _, seq := range seqs {
			seq(func(elem T) bool {
				return yield(elem)
			})
		}
	}
}

// Compress returns a sequence of elements from the input sequence.
// The resulting sequence contains only the elements where the corresponding
// selector is true.
func Compress[T any](data []T, selectors []bool) stdIter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range selectors {
			if selectors[i] {
				if !yield(data[i]) {
					return
				}
			}
		}
	}
}

// DropWhile returns a sequence of elements from the input sequence.
// The resulting sequence contains only the elements after the predicate is
// false.
func DropWhile[T any](pred func(T) bool, a []T) stdIter.Seq[T] {
	return func(yield func(T) bool) {
		var i int
		for i = 0; i < len(a) && pred(a[i]); i++ {
		}

		for ; i < len(a); i++ {
			if !yield(a[i]) {
				return
			}
		}
	}
}

// Filter returns a sequence of elements from the input sequence.
// The resulting sequence contains only the elements where the predicate is true.
func Filter[T any](pred func(T) bool, a []T) stdIter.Seq[T] {
	return func(yield func(T) bool) {
		for _, elem := range a {
			if pred(elem) && !yield(elem) {
				return
			}
		}
	}
}

// FilterFalse returns a sequence of elements from the input sequence.
// The resulting sequence contains only the elements where the predicate is false.
func FilterFalse[T any](pred func(T) bool, a []T) stdIter.Seq[T] {
	return func(yield func(T) bool) {
		for _, elem := range a {
			if !pred(elem) && !yield(elem) {
				return
			}
		}
	}
}

// GroupBy returns a sequence of groups of elements from the input sequence.
// The resulting sequence contains groups of elements where the key function
// returns the same value.
func GroupBy[T any, K comparable](key func(T) K, a []T) stdIter.Seq2[K, stdIter.Seq[T]] {
	return func(yield func(K, stdIter.Seq[T]) bool) {
		groups := make(map[K][]T)
		for _, elem := range a {
			k := key(elem)
			groups[k] = append(groups[k], elem)
		}

		// fmt.Println(groups)

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
func Map[T, U any](f func(T) U, a []T) stdIter.Seq[U] {
	return func(yield func(U) bool) {
		for _, elem := range a {
			if !yield(f(elem)) {
				return
			}
		}
	}
}

// Map2 returns a sequence of elements from the input sequences.
// The resulting sequence contains the elements after applying the function to
// each pair of elements from the input sequences.
func Map2[T, U, V any](f func(T, U) V, a []T, b []U) stdIter.Seq[V] {
	return func(yield func(V) bool) {
		minLen := min(len(a), len(b))
		for i := 0; i < minLen; i++ {
			if !yield(f(a[i], b[i])) {
				return
			}
		}
	}
}

// TakeWhile returns a sequence of elements from the input sequence.
// The resulting sequence contains only the elements before the predicate is
// false.
func TakeWhile[T any](pred func(T) bool, a []T) stdIter.Seq[T] {
	return func(yield func(T) bool) {
		for i := 0; i < len(a) && pred(a[i]); i++ {
			if !yield(a[i]) {
				return
			}
		}
	}
}
