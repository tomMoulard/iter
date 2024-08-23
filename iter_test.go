package iter_test

import (
	stdIter "iter"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tommoulard/iter"
)

func TestValues(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      stdIter.Seq[any]
		expect []any
	}{
		{
			name:   "values int",
			a:      slices.Values([]any{1, 2, 3}),
			expect: []any{1, 2, 3},
		},
		{
			name:   "values string",
			a:      slices.Values([]any{"a", "b", "c"}),
			expect: []any{"a", "b", "c"},
		},
		{
			name:   "values rune",
			a:      slices.Values([]any{'a', 'b', 'c'}),
			expect: []any{'a', 'b', 'c'},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := iter.Values(test.a)
			assert.Equal(t, test.expect, got)
		})
	}
}

func TestValues2(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		a       stdIter.Seq2[any, any]
		expectA []any
		expectB []any
	}{
		{
			name:    "values2 int",
			a:       iter.Zip([]any{1, 2, 3}, []any{4, 5, 6}),
			expectA: []any{1, 2, 3},
			expectB: []any{4, 5, 6},
		},
		{
			name:    "values2 string",
			a:       iter.Zip([]any{"a", "b", "c"}, []any{"d", "e", "f"}),
			expectA: []any{"a", "b", "c"},
			expectB: []any{"d", "e", "f"},
		},
		{
			name:    "values2 rune",
			a:       iter.Zip([]any{'a', 'b', 'c'}, []any{'d', 'e', 'f'}),
			expectA: []any{'a', 'b', 'c'},
			expectB: []any{'d', 'e', 'f'},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			gotA, gotB := iter.Values2(test.a)
			assert.Equal(t, test.expectA, gotA)
			assert.Equal(t, test.expectB, gotB)
		})
	}
}

func TestZip(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      []any
		b      []any
		expect [][]any
		length int
	}{
		{
			name:   "zip two slices of int",
			a:      []any{4, 5, 6},
			b:      []any{7, 8, 9},
			expect: [][]any{{4, 7}, {5, 8}, {6, 9}},
			length: 3,
		},
		{
			name:   "zip one slices of string, and one slice of int",
			a:      []any{"a", "b", "c"},
			b:      []any{1, 2, 3},
			expect: [][]any{{"a", 1}, {"b", 2}, {"c", 3}},
			length: 3,
		},
		{
			name:   "zip two slices of int with different length len(a) < len(b)",
			a:      []any{4, 5, 6},
			b:      []any{7, 8},
			expect: [][]any{{4, 7}, {5, 8}},
			length: 2,
		},
		{
			name:   "zip two slices of int with different length len(a) > len(b)",
			a:      []any{7, 8},
			b:      []any{4, 5, 6},
			expect: [][]any{{7, 4}, {8, 5}},
			length: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			i := 0

			for a, b := range iter.Zip(test.a, test.b) {
				assert.Equal(t, test.expect[i][0], a)
				assert.Equal(t, test.expect[i][1], b)

				i++
			}

			assert.Equal(t, test.length, i)
		})
	}
}

func TestZipLongest(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      []any
		b      []any
		fill   any
		expect [][]any
		length int
	}{
		{
			name:   "ziplongest two slices of int",
			a:      []any{4, 5, 6},
			b:      []any{7, 8, 9},
			fill:   42,
			expect: [][]any{{4, 7}, {5, 8}, {6, 9}},
			length: 3,
		},
		{
			name:   "ziplongest one slices of string, and one slice of int",
			a:      []any{"a", "b", "c"},
			b:      []any{1, 2, 3},
			fill:   "0",
			expect: [][]any{{"a", 1}, {"b", 2}, {"c", 3}},
			length: 3,
		},
		{
			name:   "ziplongest two slices of int with different length len(a) < len(b)",
			a:      []any{4, 5, 6},
			b:      []any{7, 8},
			fill:   42,
			expect: [][]any{{4, 7}, {5, 8}, {6, 42}},
			length: 3,
		},
		{
			name:   "ziplongest two slices of int with different length len(a) > len(b)",
			a:      []any{7, 8},
			b:      []any{4, 5, 6},
			fill:   42,
			expect: [][]any{{7, 4}, {8, 5}, {42, 6}},
			length: 3,
		},
		{
			name:   "ziplongest with fill of different type",
			a:      []any{7, 8},
			b:      []any{4, 5, 6},
			fill:   "42",
			expect: [][]any{{7, 4}, {8, 5}, {"42", 6}},
			length: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			i := 0

			for a, b := range iter.ZipLongest(test.a, test.b, test.fill) {
				assert.Equal(t, test.expect[i][0], a)
				assert.Equal(t, test.expect[i][1], b)

				i++
			}

			assert.Equal(t, test.length, i)
		})
	}
}

func TestAccumulate(t *testing.T) {
	t.Parallel()

	// Cannot do table tests because any does not satisfy cmp.Ordered

	t.Run("accumulate int", func(t *testing.T) {
		t.Parallel()

		a := []int{1, 2, 3, 4, 5}
		expect := []int{1, 3, 6, 10, 15}
		i := 0

		for s := range iter.Accumulate(a) {
			assert.Equal(t, expect[i], s)

			i++
		}

		assert.Equal(t, len(expect), i)
	})

	t.Run("accumulate float", func(t *testing.T) {
		t.Parallel()

		a := []float32{1, 2, 3, 4, 5}
		expect := []float32{1, 3, 6, 10, 15}
		i := 0

		for s := range iter.Accumulate(a) {
			assert.InEpsilon(t, expect[i], s, 0.01)

			i++
		}

		assert.Equal(t, len(expect), i)
	})

	t.Run("accumulate string", func(t *testing.T) {
		t.Parallel()

		a := []string{"a", "b", "c", "d", "e"}
		expect := []string{"a", "ab", "abc", "abcd", "abcde"}
		i := 0

		for s := range iter.Accumulate(a) {
			assert.Equal(t, expect[i], s)

			i++
		}

		assert.Equal(t, len(expect), i)
	})

	t.Run("accumulate slices", func(t *testing.T) {
		t.Parallel()

		a := []rune{'a', 'b', 'c', 'd', 'e', 'f'}
		expect := []int32{97, 195, 294, 394, 495, 597}
		i := 0

		for s := range iter.Accumulate(a) {
			assert.Equal(t, expect[i], s)

			i++
		}

		assert.Equal(t, len(expect), i)
	})

	t.Run("accumulate empty slice", func(t *testing.T) {
		t.Parallel()

		a := []rune{}
		expect := []int32{}
		i := 0

		for s := range iter.Accumulate(a) {
			assert.Equal(t, expect[i], s)

			i++
		}

		assert.Equal(t, len(expect), i)
	})
}

func TestChain(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		seqs   [][]any
		expect []any
		length int
	}{
		{
			name:   "chain two slices of int",
			seqs:   [][]any{{1, 2, 3}, {4, 5, 6}},
			expect: []any{1, 2, 3, 4, 5, 6},
			length: 6,
		},
		{
			name:   "chain two slices of string",
			seqs:   [][]any{{"a", "b", "c"}, {"d", "e", "f"}},
			expect: []any{"a", "b", "c", "d", "e", "f"},
			length: 6,
		},
		{
			name:   "chain three slices of int with one empty slice",
			seqs:   [][]any{{1, 2, 3}, {4, 5, 6}, {}},
			expect: []any{1, 2, 3, 4, 5, 6},
			length: 6,
		},
		{
			name:   "chain two slices of int with one empty slice",
			seqs:   [][]any{{}, {4, 5, 6}},
			expect: []any{4, 5, 6},
			length: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			i := 0

			for s := range iter.Chain(test.seqs...) {
				assert.Equal(t, test.expect[i], s)

				i++
			}

			assert.Equal(t, test.length, i)
		})
	}
}

func TestChainSeq(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		seqs   []stdIter.Seq[any]
		expect []any
		length int
	}{
		{
			name:   "chain two slices of int",
			seqs:   []stdIter.Seq[any]{slices.Values([]any{1, 2, 3}), slices.Values([]any{4, 5, 6})},
			expect: []any{1, 2, 3, 4, 5, 6},
			length: 6,
		},
		{
			name:   "chain two slices of string",
			seqs:   []stdIter.Seq[any]{slices.Values([]any{"a", "b", "c"}), slices.Values([]any{"d", "e", "f"})},
			expect: []any{"a", "b", "c", "d", "e", "f"},
			length: 6,
		},
		{
			name:   "chain three slices of int with one empty slice",
			seqs:   []stdIter.Seq[any]{slices.Values([]any{1, 2, 3}), slices.Values([]any{4, 5, 6}), slices.Values([]any{})},
			expect: []any{1, 2, 3, 4, 5, 6},
			length: 6,
		},
		{
			name:   "chain two slices of int with one empty slice",
			seqs:   []stdIter.Seq[any]{slices.Values([]any{}), slices.Values([]any{4, 5, 6})},
			expect: []any{4, 5, 6},
			length: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			i := 0

			for s := range iter.ChainSeq(test.seqs...) {
				t.Log(s)
				assert.Equal(t, test.expect[i], s)

				i++
			}

			assert.Equal(t, test.length, i)
		})
	}
}

func TestCompress(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		data      []any
		selectors []bool
		expect    []any
		length    int
	}{
		{
			name:      "compress with all true",
			data:      []any{1, 2, 3, 4, 5},
			selectors: []bool{true, true, true, true, true},
			expect:    []any{1, 2, 3, 4, 5},
			length:    5,
		},
		{
			name:      "compress with all false",
			data:      []any{1, 2, 3, 4, 5},
			selectors: []bool{false, false, false, false, false},
			expect:    []any{},
			length:    0,
		},
		{
			name:      "compress with some true",
			data:      []any{1, 2, 3, 4, 5},
			selectors: []bool{true, false, true, false, true},
			expect:    []any{1, 3, 5},
			length:    3,
		},
		{
			name:      "compress with some false",
			data:      []any{1, 2, 3, 4, 5},
			selectors: []bool{false, true, false, true, false},
			expect:    []any{2, 4},
			length:    2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			i := 0

			for s := range iter.Compress(test.data, test.selectors) {
				t.Log(s)
				assert.Equal(t, test.expect[i], s)

				i++
			}

			assert.Equal(t, test.length, i)
		})
	}
}

func TestDropWhile(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		pred   func(any) bool
		a      []any
		expect []any
		length int
	}{
		{
			name: "drop while int",
			pred: func(i any) bool {
				return i.(int) < 3
			},
			a:      []any{1, 2, 3, 4, 5},
			expect: []any{3, 4, 5},
			length: 3,
		},
		{
			name: "drop while string",
			pred: func(s any) bool {
				return s.(string) != "c"
			},
			a:      []any{"a", "b", "c", "d", "e", "c"},
			expect: []any{"c", "d", "e", "c"},
			length: 4,
		},
		{
			name: "drop while rune",
			pred: func(r any) bool {
				return r.(rune) != 'c'
			},
			a:      []any{'a', 'b', 'c', 'd', 'e'},
			expect: []any{'c', 'd', 'e'},
			length: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			i := 0

			for s := range iter.DropWhile(test.pred, test.a) {
				t.Logf("s: %v", s)
				assert.Equal(t, test.expect[i], s)

				i++
			}

			assert.Equal(t, test.length, i)
		})
	}
}

func TestFilter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		pred   func(any) bool
		a      []any
		expect []any
		length int
	}{
		{
			name: "filter int",
			pred: func(i any) bool {
				return i.(int) < 3
			},
			a:      []any{1, 2, 3, 4, 5},
			expect: []any{1, 2},
			length: 2,
		},
		{
			name: "filter string",
			pred: func(s any) bool {
				return s.(string) != "c"
			},
			a:      []any{"a", "b", "c", "d", "e", "c"},
			expect: []any{"a", "b", "d", "e"},
			length: 4,
		},
		{
			name: "filter rune",
			pred: func(r any) bool {
				return r.(rune) != 'c'
			},
			a:      []any{'a', 'b', 'c', 'd', 'e'},
			expect: []any{'a', 'b', 'd', 'e'},
			length: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			i := 0

			for s := range iter.Filter(test.pred, test.a) {
				t.Logf("s: %v", s)
				assert.Equal(t, test.expect[i], s)

				i++
			}

			assert.Equal(t, test.length, i)
		})
	}
}

func TestFilterFalse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		pred   func(any) bool
		a      []any
		expect []any
		length int
	}{
		{
			name: "filterfalse int",
			pred: func(i any) bool {
				return i.(int) < 3
			},
			a:      []any{1, 2, 3, 4, 5},
			expect: []any{3, 4, 5},
			length: 3,
		},
		{
			name: "filterfalse string",
			pred: func(s any) bool {
				return s.(string) != "c"
			},
			a:      []any{"a", "b", "c", "d", "e", "c"},
			expect: []any{"c", "c"},
			length: 2,
		},
		{
			name: "filterfalse rune",
			pred: func(r any) bool {
				return r.(rune) != 'c'
			},
			a:      []any{'a', 'b', 'c', 'd', 'e'},
			expect: []any{'c'},
			length: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			i := 0

			for s := range iter.FilterFalse(test.pred, test.a) {
				t.Logf("s: %v", s)
				assert.Equal(t, test.expect[i], s)

				i++
			}

			assert.Equal(t, test.length, i)
		})
	}
}

func TestGroupBy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		key    func(any) any
		a      []any
		expect map[any][]any
	}{
		{
			name: "groupby int",
			key: func(i any) any {
				return i.(int) % 2
			},
			a: []any{1, 2, 3, 4, 5},
			expect: map[any][]any{
				0: {2, 4},
				1: {1, 3, 5},
			},
		},
		// {
		// name: "groupby string",
		// key: func(s any) any {
		// return s.(string)[0]
		// },
		// a: []any{"apple", "banana", "avocado", "cherry", "blueberry"},
		// expect: map[any][]any{
		// 'a': {"apple", "avocado"},
		// 'b': {"banana", "blueberry"},
		// 'c': {"cherry"},
		// },
		// },
		{
			name: "groupby rune",
			key: func(r any) any {
				return r.(rune)
			},
			a: []any{'a', 'b', 'c', 'd', 'e'},
			expect: map[any][]any{
				'a': {'a'},
				'b': {'b'},
				'c': {'c'},
				'd': {'d'},
				'e': {'e'},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			t.Logf("expect: %v", test.expect)

			for key, group := range iter.GroupBy(test.key, test.a) {
				t.Logf("key: %v, group: %v, expect: %v", key, iter.Values(group), test.expect)

				expected, ok := test.expect[key]
				require.True(t, ok)

				assert.ElementsMatch(t, expected, iter.Values(group))
			}
		})
	}
}

func TestMap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		fn     func(any) any
		a      []any
		expect []any
		length int
	}{
		{
			name: "map int",
			fn: func(i any) any {
				return i.(int) * 2
			},
			a:      []any{1, 2, 3, 4, 5},
			expect: []any{2, 4, 6, 8, 10},
			length: 5,
		},
		{
			name: "map string",
			fn: func(s any) any {
				return s.(string) + "!"
			},
			a:      []any{"a", "b", "c", "d", "e"},
			expect: []any{"a!", "b!", "c!", "d!", "e!"},
			length: 5,
		},
		{
			name: "map rune",
			fn: func(r any) any {
				return r.(rune) + 1
			},
			a:      []any{'a', 'b', 'c', 'd', 'e'},
			expect: []any{'b', 'c', 'd', 'e', 'f'},
			length: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			i := 0

			for s := range iter.Map(test.fn, test.a) {
				t.Logf("s: %v", s)
				assert.Equal(t, test.expect[i], s)

				i++
			}

			assert.Equal(t, test.length, i)
		})
	}
}

func TestMap2(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		fn     func(any, any) any
		a      []any
		b      []any
		expect []any
		length int
	}{
		{
			name: "map2 int",
			fn: func(a, b any) any {
				return a.(int) + b.(int)
			},
			a:      []any{1, 2, 3, 4, 5},
			b:      []any{2, 3, 4, 5, 6},
			expect: []any{3, 5, 7, 9, 11},
			length: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			i := 0

			for s := range iter.Map2(test.fn, test.a, test.b) {
				t.Logf("s: %v", s)
				assert.Equal(t, test.expect[i], s)

				i++
			}

			assert.Equal(t, test.length, i)
		})
	}
}

func TestTakeWhile(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		pred   func(any) bool
		a      []any
		expect []any
		length int
	}{
		{
			name: "take while int",
			pred: func(i any) bool {
				return i.(int) < 3
			},
			a:      []any{1, 2, 3, 4, 5},
			expect: []any{1, 2},
			length: 2,
		},
		{
			name: "take while string",
			pred: func(s any) bool {
				return s.(string) != "c"
			},
			a:      []any{"a", "b", "c", "d", "e", "c"},
			expect: []any{"a", "b"},
			length: 2,
		},
		{
			name: "take while rune",
			pred: func(r any) bool {
				return r.(rune) != 'c'
			},
			a:      []any{'a', 'b', 'c', 'd', 'e'},
			expect: []any{'a', 'b'},
			length: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			i := 0

			for s := range iter.TakeWhile(test.pred, test.a) {
				t.Logf("s: %v", s)
				assert.Equal(t, test.expect[i], s)

				i++
			}

			assert.Equal(t, test.length, i)
		})
	}
}
