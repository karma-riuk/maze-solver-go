package utils

import (
	"log"

	"golang.org/x/exp/constraints"
)

func Check(err error, msg string, args ...any) {
	if err != nil {
		log.Printf(msg, args...)
		panic(err)
	}
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}
