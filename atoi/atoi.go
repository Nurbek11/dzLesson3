package dz3Lesson

import (
	"strconv"
)

func Atoi(s string) (int, error) {
	a, err := strconv.Atoi(s)
	return a, err
}
