package generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {

	t.Run("ints", func(t *testing.T) {
		s := Stack[int]{}
		for i := 1; i <= 5; i++ {
			s.Push(i)
		}

		assert.Equal(t, 5, s.Len())

		expectedStr := "[ 1 2 3 4 5 ]"
		assert.Equal(t, expectedStr, s.String())

		got, err := s.Pop()
		require.NoError(t, err)
		assert.Equal(t, 4, s.Len())
		assert.Equal(t, 5, got)

		for s.Len() != 0 {
			_, err = s.Pop()
			require.NoError(t, err)
		}

		_, err = s.Pop()
		assert.Error(t, err)

	})

	t.Run("strings", func(t *testing.T) {
		s := Stack[string]{}
		for _, str := range []string{"a", "b", "c", "d", "e"} {
			s.Push(str)
		}

		assert.Equal(t, 5, s.Len())

		expectedStr := "[ a b c d e ]"
		assert.Equal(t, expectedStr, s.String())

		got, err := s.Pop()
		require.NoError(t, err)
		assert.Equal(t, 4, s.Len())
		assert.Equal(t, "e", got)

		for s.Len() != 0 {
			_, err = s.Pop()
			require.NoError(t, err)
		}

		_, err = s.Pop()
		assert.Error(t, err)

	})

}

func TestQueue(t *testing.T) {

	t.Run("ints", func(t *testing.T) {
		q := Queue[int]{}
		for i := 1; i <= 5; i++ {
			q.Push(i)
		}

		assert.Equal(t, 5, q.Len())

		expectedStr := "[ 5 4 3 2 1 ]"
		assert.Equal(t, expectedStr, q.String())

		got, err := q.Pop()
		require.NoError(t, err)
		assert.Equal(t, 4, q.Len())
		assert.Equal(t, 1, got)

		for q.Len() != 0 {
			_, err = q.Pop()
			require.NoError(t, err)
		}

		_, err = q.Pop()
		assert.Error(t, err)

	})

	t.Run("strings", func(t *testing.T) {
		q := Queue[string]{}
		for _, str := range []string{"a", "b", "c", "d", "e"} {
			q.Push(str)
		}

		assert.Equal(t, 5, q.Len())

		expectedStr := "[ e d c b a ]"
		assert.Equal(t, expectedStr, q.String())

		got, err := q.Pop()
		require.NoError(t, err)
		assert.Equal(t, 4, q.Len())
		assert.Equal(t, "a", got)

		for q.Len() != 0 {
			_, err = q.Pop()
			require.NoError(t, err)
		}

		_, err = q.Pop()
		assert.Error(t, err)

	})

}
