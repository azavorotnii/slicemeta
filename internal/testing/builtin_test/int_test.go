package builtin_test

//go:generate slicemeta -type int -less operator
import (
	"math/rand"
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/azavorotnii/slicemeta/internal/testing/builtin_test/intutil"
)

func even(v int) bool {
	return v%2 == 0
}
func odd(v int) bool {
	return !even(v)
}
func twoDigits(v int) bool {
	return v >= 10 && v < 100
}

func TestContains(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	assert.True(t, intutil.Contains(input, 8))
	assert.False(t, intutil.Contains(input, 10))

	assert.True(t, intutil.ContainsAny(input, -1, 10, 5))
	assert.False(t, intutil.ContainsAny(input, -1, 10))

	assert.True(t, intutil.ContainsFunc(input, even))
	assert.False(t, intutil.ContainsFunc(input, twoDigits))
}

func TestCount(t *testing.T) {
	input := []int{5, 4, 3, 2, 4, 3, 4, 2, 0}

	assert.Equal(t, 3, intutil.Count(input, 4))
	assert.Equal(t, 0, intutil.Count(input, 1))

	assert.Equal(t, 6, intutil.CountAny(input, 8, 6, 4, 2, 0))
	assert.Equal(t, 3, intutil.CountAny(input, 9, 7, 5, 3, 1))

	assert.Equal(t, 6, intutil.CountFunc(input, even))
	assert.Equal(t, 3, intutil.CountFunc(input, odd))
}

func TestEqual(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}

	sameValue := []int{1, 2, 3, 4, 5, 6, 7}
	sameLength := []int{1, 2, 3, 4, 5, 6, 6}
	lessValues := []int{0, 1, 2}
	moreValues := append(input, 8, 9)

	assert.True(t, intutil.Equal(input, sameValue))
	assert.False(t, intutil.Equal(input, sameLength))
	assert.False(t, intutil.Equal(input, lessValues))
	assert.False(t, intutil.Equal(input, moreValues))
}

func TestFilter(t *testing.T) {
	input := []int{123, 78, 56, 45, 9, 4, 20}

	assert.Equal(t, []int{78, 56, 4, 20}, intutil.Filter(input, even))
	assert.Equal(t, []int{123, 45, 9}, intutil.Filter(input, odd))
	assert.Equal(t, []int{78, 56, 45, 20}, intutil.Filter(input, twoDigits))
}

func TestIndex(t *testing.T) {
	input := []int{5, 4, 3, 2, 4, 3, 4, 2, 0}

	assert.Equal(t, 0, intutil.Index(input, 5))
	assert.Equal(t, 1, intutil.Index(input, 4))
	assert.Equal(t, 8, intutil.Index(input, 0))
	assert.Equal(t, -1, intutil.Index(input, 1))

	assert.Equal(t, 0, intutil.IndexAny(input, 5, 4, 0, 1))
	assert.Equal(t, 1, intutil.IndexAny(input, 2, 3, 4))
	assert.Equal(t, 8, intutil.IndexAny(input, 1, 0, 6))
	assert.Equal(t, -1, intutil.IndexAny(input, 1, 7, 6))

	assert.Equal(t, 1, intutil.IndexFunc(input, even))
	assert.Equal(t, 0, intutil.IndexFunc(input, odd))
	assert.Equal(t, -1, intutil.IndexFunc(input, twoDigits))
}

func TestLastIndex(t *testing.T) {
	input := []int{5, 4, 3, 2, 4, 3, 4, 2, 0}

	assert.Equal(t, 0, intutil.LastIndex(input, 5))
	assert.Equal(t, 6, intutil.LastIndex(input, 4))
	assert.Equal(t, 8, intutil.LastIndex(input, 0))
	assert.Equal(t, -1, intutil.LastIndex(input, 1))

	assert.Equal(t, 8, intutil.LastIndexAny(input, 5, 4, 0, 1))
	assert.Equal(t, 7, intutil.LastIndexAny(input, 2, 3, 4))
	assert.Equal(t, 8, intutil.LastIndexAny(input, 1, 0, 6))
	assert.Equal(t, -1, intutil.LastIndexAny(input, 1, 7, 6))

	assert.Equal(t, 8, intutil.LastIndexFunc(input, even))
	assert.Equal(t, 5, intutil.LastIndexFunc(input, odd))
	assert.Equal(t, -1, intutil.LastIndexFunc(input, twoDigits))
}

func TestMap(t *testing.T) {
	input := []int{5, 4, 3, 2, 4, 3, 4, 2, 0}

	sqr := func(v int) int {
		return v * v
	}
	assert.Equal(t, []int{25, 16, 9, 4, 16, 9, 16, 4, 0}, intutil.Map(input, sqr))
}

func TestReduce(t *testing.T) {
	input := []int{5, 4, 3, 2, 4, 3, 4, 2}

	mul := func(l, r int) int {
		return l * r
	}
	assert.Equal(t, 11520, intutil.Reduce(input, mul))
}

func TestShuffle(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	assert.True(t, sort.IsSorted(sort.IntSlice(input)))
	intutil.Shuffle(input)
	assert.False(t, sort.IsSorted(sort.IntSlice(input)))
}

func TestSort(t *testing.T) {
	input := intutil.IntSlice([]int{5, 4, 3, 2, 4, 3, 4, 2, 0})

	assert.False(t, sort.IsSorted(input))
	sort.Sort(input)
	assert.True(t, sort.IsSorted(input))
}

func BenchmarkEqual(b *testing.B) {
	const size = 100000
	left := rand.Perm(size)
	right := make([]int, size)
	copy(right, left)

	b.Run("DeepEqual", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflect.DeepEqual(left, right)
		}
	})

	b.Run("intutil", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			intutil.Equal(left, right)
		}
	})
}
