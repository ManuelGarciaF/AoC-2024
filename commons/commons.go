package commons

import "strconv"

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
