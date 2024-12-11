package commons

import (
	"fmt"
	"maps"
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

func (s Set[T]) Add(e T) Set[T] {
	s[e] = struct{}{}
	return s
}

func (s Set[T]) Contains(e T) bool {
	_, ok := s[e]
	return ok
}

func (s Set[T]) Remove(e T) Set[T] {
	delete(s, e)
	return s
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) Clone() Set[T] {
	return maps.Clone(s)
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	for v := range other {
		s.Add(v)
	}

	return s
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

func Map[T any, U any](s []T, f func(T) U) []U {
	new := make([]U, len(s))
	for i, v := range s {
		new[i] = f(v)
	}
	return new
}

func FlatMap[T any, U any](s []T, f func(T) []U) []U {
	new := make([]U, 0)
	for _, v := range s {
		new = append(new, f(v)...)
	}
	return new
}

func Foldl[T, U any](seed T, xs []U, f func(T, U) T) T {
	acc := seed
	for _, x := range xs {
		acc = f(acc, x)
	}
	return acc
}

func Sum(xs []int) int {
	return Foldl(0, xs, func(acc, x int) int { return acc + x })
}

/* Does not work with recursive functions unless created inline with the
 * variable declared before.
 * e.g.:
 *    var fib func(int) int
 *    fib = c.Memoize(func(n int) int {
 *        code that uses fib()
 *    })
 */
func Memoize[T comparable, U any](f func(T) U) func(T) U {
	cache := make(map[T]U)
	return func(t T) U {
		if v, ok := cache[t]; ok {
			return v
		}
		cache[t] = f(t)
		return cache[t]
	}
}
