package iter_test

import (
	"fmt"
	"slices"

	"github.com/tommoulard/iter"
)

func ExampleValues() {
	intSeq := slices.Values([]int{1, 2, 3})
	fmt.Println(iter.Values(intSeq))

	// Output:
	// [1 2 3]
}

func ExampleValues2() {
	intSeq2 := iter.Zip([]int{1, 2, 3}, []int{4, 5, 6})
	fmt.Println(iter.Values2(intSeq2))

	// Output:
	// [1 2 3] [4 5 6]
}

func ExampleValues2Map() {
	intSeq2 := iter.Zip([]int{1, 2, 3}, []int{4, 5, 6})
	fmt.Println(iter.Values2Map(intSeq2))

	// Output:
	// map[1:4 2:5 3:6]
}

func ExampleZip() {
	for a, b := range iter.Zip([]int{1, 2, 3}, []int{4, 5, 6}) {
		fmt.Println(a, b)
	}

	// Output:
	// 1 4
	// 2 5
	// 3 6
}

func ExampleIZip() {
	for a, b := range iter.IZip(iter.Chain([]int{1, 2, 3}), iter.Chain([]int{4, 5, 6})) {
		fmt.Println(a, b)
	}

	// Output:
	// 1 4
	// 2 5
	// 3 6
}
func ExampleZipLongest() {
	for a, b := range iter.ZipLongest([]int{1, 2, 3}, []int{4}, 0) {
		fmt.Println(a, b)
	}

	// Output:
	// 1 4
	// 2 0
	// 3 0
}

func ExampleAccumulate() {
	for a := range iter.Accumulate([]int{1, 2, 3, 4, 5}) {
		fmt.Println(a)
	}

	// Output:
	// 1
	// 3
	// 6
	// 10
	// 15
}

func ExampleChain() {
	for a := range iter.Chain([]int{1, 2, 3}, []int{4, 5, 6}) {
		fmt.Println(a)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
}

func ExampleChainSeq() {
	for a := range iter.ChainSeq(slices.Values([]int{1, 2, 3}), slices.Values([]int{4, 5, 6})) {
		fmt.Println(a)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
}

func ExampleCompress() {
	for a := range iter.Compress([]int{1, 2, 3}, []bool{true, false, true}) {
		fmt.Println(a)
	}

	// Output:
	// 1
	// 3
}

func ExampleDropWhile() {
	for a := range iter.DropWhile(func(x int) bool { return x < 5 }, []int{1, 4, 6, 3, 8}) {
		fmt.Println(a)
	}

	// Output:
	// 6
	// 3
	// 8
}

func ExampleFilter() {
	for a := range iter.Filter(func(x int) bool { return x%2 == 0 }, []int{1, 2, 3, 4, 5}) {
		fmt.Println(a)
	}

	// Output:
	// 2
	// 4
}

func ExampleFilterFalse() {
	for a := range iter.FilterFalse(func(x int) bool { return x%2 == 0 }, []int{1, 2, 3, 4, 5}) {
		fmt.Println(a)
	}

	// Output:
	// 1
	// 3
	// 5
}

func ExampleGroupBy() {
	out := make(map[int][]int)
	for k, v := range iter.GroupBy(func(x int) int { return x % 2 }, []int{1, 2, 3, 4, 5}) {
		out[k] = iter.Values(v)
	}

	fmt.Println(out[0])
	fmt.Println(out[1])

	// Output:
	// [2 4]
	// [1 3 5]
}

func ExampleMap() {
	for a := range iter.Map(func(x int) int { return x * 2 }, []int{1, 2, 3, 4, 5}) {
		fmt.Println(a)
	}

	// Output:
	// 2
	// 4
	// 6
	// 8
	// 10
}

func ExampleMap2() {
	for a := range iter.Map2(func(x, y int) int { return x + y }, []int{1, 2, 3}, []int{4, 5, 6}) {
		fmt.Println(a)
	}

	// Output:
	// 5
	// 7
	// 9
}

func ExampleTakeWhile() {
	for a := range iter.TakeWhile(func(x int) bool { return x < 5 }, []int{1, 4, 6, 3, 8}) {
		fmt.Println(a)
	}

	// Output:
	// 1
	// 4
}

func ExampleChainMap() {
	m := map[int]string{
		int('b'): "b",
		int('c'): "c",
		int('a'): "a",
	}
	for key, value := range iter.ChainMap(m) {
		fmt.Println(key, value)
	}

	// Output:
	// 97 a
	// 98 b
	// 99 c
}

func ExamplePermutations() {
	for value := range iter.Permutations([]int{1, 2, 3}) {
		fmt.Println(value)
	}

	// Output:
	// [1 2 3]
	// [2 1 3]
	// [3 1 2]
	// [1 3 2]
	// [2 3 1]
	// [3 2 1]
}

func ExamplePermutationsLen() {
	for value := range iter.PermutationsLen([]int{1, 2, 3}, 2) {
		fmt.Println(value)
	}

	// Output:
	// [1 2]
	// [2 1]
	// [3 1]
	// [1 3]
	// [2 3]
	// [3 2]
}
