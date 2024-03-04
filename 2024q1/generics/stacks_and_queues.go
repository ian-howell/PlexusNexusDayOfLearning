package generics

import (
	"errors"
	"fmt"
	"strings"
)

type Stack[T any] []T

func (s *Stack[T]) Push(val T) {
	*s = append(*s, val)
}

func (s *Stack[T]) Pop() (T, error) {
	var rv T
	if len(*s) == 0 {
		return rv, errors.New("stack is empty")
	}
	rv = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return rv, nil
}

func (s Stack[T]) Len() int {
	return len(s)
}

func (s Stack[T]) String() string {
	sb := strings.Builder{}
	sb.WriteString("[ ")
	for _, val := range s {
		sb.WriteString(fmt.Sprintf("%v ", val))
	}
	sb.WriteString("]")
	return sb.String()
}

type Queue[T any] []T

func (q *Queue[T]) Push(val T) {
	*q = append([]T{val}, (*q)...)
}

func (q *Queue[T]) Pop() (T, error) {
	var rv T
	if len(*q) == 0 {
		return rv, errors.New("stack is empty")
	}
	rv = (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return rv, nil
}

func (q Queue[T]) Len() int {
	return len(q)
}

func (q Queue[T]) String() string {
	sb := strings.Builder{}
	sb.WriteString("[ ")
	for _, val := range q {
		sb.WriteString(fmt.Sprintf("%v ", val))
	}
	sb.WriteString("]")
	return sb.String()
}
