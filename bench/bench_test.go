package bench

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/samber/lo"
	"github.com/tommoulard/iter"
)

func BenchmarkZipInt(b *testing.B) {
	aa, ab, ac := helperBuildIntSlices()
	ba, bb, bc := helperBuildIntSlices()
	m := map[string][][]int{
		"small":  {aa, ba},
		"medium": {ab, bb},
		"large":  {ac, bc},
	}

	for name, args := range m {
		b.Run("raw-"+name, func(b *testing.B) {
			b.ResetTimer()

			var acc int
			for i := 0; i < b.N; i++ {
				for _, a := range args[0] {
					acc += a
				}
				for _, b := range args[1] {
					acc += b
				}
			}
		})

		b.Run("iter-"+name, func(b *testing.B) {
			b.ResetTimer()

			var acc int
			for i := 0; i < b.N; i++ {
				for a, b := range iter.Zip(args[0], args[1]) {
					acc += a + b
				}
			}
		})

		b.Run("lo-"+name, func(b *testing.B) {
			b.ResetTimer()

			var acc int
			for i := 0; i < b.N; i++ {
				for _, a := range lo.Zip2(args[0], args[1]) {
					acc += a.A + a.B
				}
			}
		})
	}
}

// func BenchmarkZipSimple(b *testing.B) {
// aa, ab, ac := helperBuildSimpleSlices()
// ba, bb, bc := helperBuildSimpleSlices()
// m := map[string][][]simple{
// "small":  {aa, ba},
// "medium": {ab, bb},
// "large":  {ac, bc},
// }

// for name, args := range m {
// b.Run("raw-"+name, func(b *testing.B) {
// b.ResetTimer()

// var acc string
// for i := 0; i < b.N; i++ {
// for _, a := range args[0] {
// acc += a.attr1
// }
// for _, b := range args[1] {
// acc += b.attr1
// }
// }
// })

// b.Run("iter-"+name, func(b *testing.B) {
// b.ResetTimer()

// var acc string
// for i := 0; i < b.N; i++ {
// for a, b := range iter.Zip(args[0], args[1]) {
// acc += a.attr1 + b.attr1
// }
// }
// })

// b.Run("lo-"+name, func(b *testing.B) {
// b.ResetTimer()

// var acc string
// for i := 0; i < b.N; i++ {
// for _, a := range lo.Zip2(args[0], args[1]) {
// acc += a.A.attr1 + a.B.attr1
// }
// }
// })
// }
// }

// func BenchmarkZipComplexe(b *testing.B) {
// aa, ab, ac := helperBuildComplexSlices()
// ba, bb, bc := helperBuildComplexSlices()
// m := map[string][][]complex{
// "small":  {aa, ba},
// "medium": {ab, bb},
// "large":  {ac, bc},
// }

// for name, args := range m {
// b.Run("raw-"+name, func(b *testing.B) {
// b.ResetTimer()

// var acc int
// for i := 0; i < b.N; i++ {
// for _, a := range args[0] {
// acc += a.id
// }
// for _, b := range args[1] {
// acc += b.id
// }
// }
// })

// b.Run("iter-"+name, func(b *testing.B) {
// b.ResetTimer()

// var acc int
// for i := 0; i < b.N; i++ {
// for a, b := range iter.Zip(args[0], args[1]) {
// acc += a.id + b.id
// }
// }
// })

// b.Run("lo-"+name, func(b *testing.B) {
// b.ResetTimer()

// var acc int
// for i := 0; i < b.N; i++ {
// for _, a := range lo.Zip2(args[0], args[1]) {
// acc += a.A.id + a.B.id
// }
// }
// })
// }
// }

func isPair(a int) bool {
	return a%2 == 0
}

func isPairIdx(a, _ int) bool {
	return a%2 == 0
}

func BenchmarkFilterInt(b *testing.B) {
	aa, ab, ac := helperBuildIntSlices()
	m := map[string][]int{
		"small":  aa,
		"medium": ab,
		"large":  ac,
	}

	for name, args := range m {
		b.Run("raw-"+name, func(b *testing.B) {
			b.ResetTimer()

			var acc int
			for i := 0; i < b.N; i++ {
				for _, a := range args {
					if isPair(a) {
						acc += a
					}
				}
			}
		})

		b.Run("iter-"+name, func(b *testing.B) {
			b.ResetTimer()

			var acc int
			for i := 0; i < b.N; i++ {
				for a := range iter.Filter(isPair, args) {
					acc += a
				}
			}
		})

		b.Run("lo-"+name, func(b *testing.B) {
			b.ResetTimer()

			var acc int
			for i := 0; i < b.N; i++ {
				for _, a := range lo.Filter(args, isPairIdx) {
					acc += a
				}
			}
		})
	}
}

// func containsSimple(a simple) bool {
// return strings.Contains(a.attr1, "6")
// }

// func containsSimpleIdx(a simple, _ int) bool {
// return strings.Contains(a.attr1, "6")
// }

// func BenchmarkFilterSimple(b *testing.B) {
// aa, ab, ac := helperBuildSimpleSlices()
// m := map[string][]simple{
// "small":  aa,
// "medium": ab,
// "large":  ac,
// }

// for name, args := range m {
// b.Run("raw-"+name, func(b *testing.B) {
// b.ResetTimer()

// var acc string
// for i := 0; i < b.N; i++ {
// for _, a := range args {
// if strings.Contains(a.attr1, "6") {
// acc += a.attr1
// }
// }
// }
// })

// b.Run("iter-"+name, func(b *testing.B) {
// b.ResetTimer()

// var acc string
// for i := 0; i < b.N; i++ {
// for a := range iter.Filter(containsSimple, args) {
// acc += a.attr1
// }
// }
// })

// b.Run("lo-"+name, func(b *testing.B) {
// b.ResetTimer()

// var acc string
// for i := 0; i < b.N; i++ {
// for _, a := range lo.Filter(args, containsSimpleIdx) {
// acc += a.attr1
// }
// }
// })
// }
// }

// func BenchmarkFilterComplex(b *testing.B) {
// aa, ab, ac := helperBuildComplexSlices()
// m := map[string][]complex{
// "small":  aa,
// "medium": ab,
// "large":  ac,
// }

// for name, args := range m {
// b.Run("raw-"+name, func(b *testing.B) {
// b.ResetTimer()

// var acc int
// for i := 0; i < b.N; i++ {
// for _, a := range args {
// if isPair(a.id) {
// acc += a.id
// }
// }
// }
// })

// b.Run("iter-"+name, func(b *testing.B) {
// b.ResetTimer()

// var acc int
// for i := 0; i < b.N; i++ {
// for a := range iter.Filter(func(c complex) bool {
// return isPair(c.id)
// }, args) {
// acc += a.id
// }
// }
// })

// b.Run("lo-"+name, func(b *testing.B) {
// b.ResetTimer()

// var acc int
// for i := 0; i < b.N; i++ {
// for _, a := range lo.Filter(args, func(item complex, index int) bool {
// return isPair(item.id)
// }) {
// acc += a.id
// }
// }
// })
// }
// }

func BenchmarkGroupByInt(b *testing.B) {
	aa, ab, ac := helperBuildIntSlices()
	m := map[string][]int{
		"small":  aa,
		"medium": ab,
		"large":  ac,
	}

	for name, args := range m {
		b.Run("raw-"+name, func(b *testing.B) {
			b.ResetTimer()

			var acc int
			for i := 0; i < b.N; i++ {
				for _, a := range args {
					if isPair(a) {
						acc += a
					}
				}
			}
		})

		b.Run("iter-"+name, func(b *testing.B) {
			b.ResetTimer()

			var acc int
			for i := 0; i < b.N; i++ {
				for a, b := range iter.GroupBy(isPair, args) {
					if !a {
						continue
					}

					for c := range b {
						acc += c
					}
				}
			}
		})

		b.Run("lo-"+name, func(b *testing.B) {
			b.ResetTimer()

			var acc int
			for i := 0; i < b.N; i++ {
				for a, b := range lo.GroupBy(args, isPair) {
					if !a {
						continue
					}

					for c := range b {
						acc += c
					}
				}
			}
		})
	}
}

// TODO: BenchmarkGroupBySimple
// TODO: BenchmarkGroupBycomplex

func BenchmarkMapInt(b *testing.B) {
	aa, ab, ac := helperBuildIntSlices()
	m := map[string][]int{
		"small":  aa,
		"medium": ab,
		"large":  ac,
	}

	for name, args := range m {
		b.Run("raw-"+name, func(b *testing.B) {
			b.ResetTimer()

			var acc int
			for i := 0; i < b.N; i++ {
				for _, a := range args {
					acc += a * 2
				}
			}
		})

		b.Run("iter-"+name, func(b *testing.B) {
			b.ResetTimer()

			var acc int
			for i := 0; i < b.N; i++ {
				for a := range iter.Map(func(i int) int { return i * 2 }, args) {
					acc += a
				}
			}
		})

		b.Run("lo-"+name, func(b *testing.B) {
			b.ResetTimer()

			var acc int
			for i := 0; i < b.N; i++ {
				for _, a := range lo.Map(args, func(item, _ int) int { return item * 2 }) {
					acc += a
				}
			}
		})
	}
}

// TODO: BenchmarkMapSimple
// TODO: BenchmarkMapComplex

func sliceToMap[T any](a []T) map[int][]T {
	m := make(map[int][]T)
	for i := range a {
		m[i] = a[:i]
	}
	return m
}

func BenchmarkChainMapInt(b *testing.B) {
	aa, ab, ac := helperBuildIntSlices()
	m := map[string]map[int][]int{
		"small":  sliceToMap(aa),
		"medium": sliceToMap(ab),
		"large":  sliceToMap(ac),
	}

	for name, args := range m {
		b.Run("raw-"+name, func(b *testing.B) {
			b.ResetTimer()

			var acc int
			for i := 0; i < b.N; i++ {
				for _, a := range args {
					for _, b := range a {
						acc += b
					}
				}
			}
		})

		b.Run("iter-"+name, func(b *testing.B) {
			b.ResetTimer()

			var acc int
			for i := 0; i < b.N; i++ {
				for _, a := range iter.ChainMap(args) {
					for _, b := range a {
						acc += b
					}
				}
			}
		})

		b.Run("lo-"+name, func(b *testing.B) {
			b.ResetTimer()

			var acc int
			for i := 0; i < b.N; i++ {
				for _, a := range lo.Entries(args) {
					for _, b := range a.Value {
						acc += b
					}
				}
			}
		})
	}
}

// TODO: BenchmarkChainMapSimple
// TODO: BenchmarkChainMapComplex

// helperBuildEasySlices returns three slices of random integers.
// len(first)  == 10
// len(second) == 10 000
// len(third)  == 100 000 000
func helperBuildIntSlices() ([]int, []int, []int) {
	return generateSliceInt(10), generateSliceInt(10000), generateSliceInt(100000000)
}

func generateSliceInt(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = rand.Int()
	}
	return slice
}

type simple struct {
	attr1 string
}

func helperBuildSimpleSlices() ([]simple, []simple, []simple) {
	return generateSliceSimple(10), generateSliceSimple(10000), generateSliceSimple(100000000)
}

func generateSliceSimple(n int) []simple {
	slice := make([]simple, n)
	for i := 0; i < n; i++ {
		slice[i] = simple{attr1: strconv.Itoa(rand.Int())}
	}
	return slice
}

type complex struct {
	id      int
	name    string
	simple  simple
	complex []simple
}

func helperBuildComplexSlices() ([]complex, []complex, []complex) {
	return generateSliceComplex(10), generateSliceComplex(10000), generateSliceComplex(100000000)
}

func generateSliceComplex(n int) []complex {
	slice := make([]complex, n)
	for i := 0; i < n; i++ {
		slice[i] = complex{
			id:      rand.Int(),
			name:    strconv.Itoa(rand.Int()),
			simple:  simple{attr1: strconv.Itoa(rand.Int())},
			complex: generateSliceSimple(10),
		}
	}
	return slice
}
