package generics

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {

	t.Run("double-ints", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		want := []int{2, 4, 6, 8, 10}
		fn := func(x int) int { return x * 2 }
		assert.Equal(t, Map(fn, input), want)
	})

	t.Run("double-floats", func(t *testing.T) {
		input := []float64{1, 2, 3, 4, 5}
		want := []float64{2, 4, 6, 8, 10}
		fn := func(x float64) float64 { return x * 2 }
		// This may fail? I'm not sure how to compare floats using assert
		assert.Equal(t, Map(fn, input), want)
	})

	t.Run("capitalize-strings", func(t *testing.T) {
		input := []string{"Hello", "World", "How", "Are", "You"}
		want := []string{"HELLO", "WORLD", "HOW", "ARE", "YOU"}
		fn := func(x string) string { return strings.ToUpper(x) }
		assert.Equal(t, Map(fn, input), want)
	})

	t.Run("custom-types", func(t *testing.T) {
		type MyType struct {
			Value int
		}
		input := []MyType{{1}, {2}, {3}, {4}, {5}}
		want := []MyType{{2}, {4}, {6}, {8}, {10}}
		fn := func(x MyType) MyType { return MyType{x.Value * 2} }
		assert.Equal(t, Map(fn, input), want)
	})

}

func TestFilter(t *testing.T) {

	t.Run("evens", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		want := []int{2, 4}
		fn := func(x int) bool { return x%2 == 0 }
		assert.Equal(t, Filter(fn, input), want)
	})

	t.Run("short-strings", func(t *testing.T) {
		input := []string{"Hello", "World", "How", "Are", "You"}
		want := []string{"How", "Are", "You"}
		fn := func(x string) bool { return len(x) <= 3 }
		assert.Equal(t, Filter(fn, input), want)
	})

	t.Run("custom-types", func(t *testing.T) {
		type MyType struct {
			Value int
		}
		input := []MyType{{1}, {2}, {3}, {4}, {5}}
		want := []MyType{{2}, {4}}
		fn := func(x MyType) bool { return x.Value%2 == 0 }
		assert.Equal(t, Filter(fn, input), want)
	})

}

func TestReduce(t *testing.T) {

	t.Run("sum-ints", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		want := 15
		fn := func(acc, x int) int { return acc + x }
		assert.Equal(t, Reduce(fn, input, 0), want)
	})

	t.Run("sum-floats", func(t *testing.T) {
		input := []float64{1, 2, 3, 4, 5}
		want := 15.0
		fn := func(acc, x float64) float64 { return acc + x }
		// This may fail? I'm not sure how to compare floats using assert
		assert.Equal(t, Reduce(fn, input, 0.0), want)
	})

	t.Run("concat-strings", func(t *testing.T) {
		input := []string{"Hello", "World", "How", "Are", "You"}
		want := "HelloWorldHowAreYou"
		fn := func(acc, x string) string { return acc + x }
		assert.Equal(t, Reduce(fn, input, ""), want)
	})

}
