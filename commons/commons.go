package commons

import (
	"fmt"
	"strconv"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Avoid if err!=nil in non important cases
func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func MustAtoi(s string) int {
	return Must(strconv.Atoi(s))
}

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Add(e T) {
	s[e] = struct{}{}
}

func (s Set[T]) Contains(e T) bool {
	_, ok := s[e]
	return ok
}

func (s Set[T]) Remove(e T) {
	delete(s, e)
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) String() string {
	str := "{"
	for e := range s {
		str += fmt.Sprintf("%v", e) + ", "
	}
	if len(str) > 1 {
		str = str[:len(str)-2]
	}
	str += "}"
	return str
}
