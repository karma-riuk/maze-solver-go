package utils

import (
	"log"
	"testing"

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

func AssertEqual[T comparable](t *testing.T, got T, want T, msg string, args ...any) {
	args = append(args, got, want)
	if got != want {
		t.Fatalf(msg+"\nGot: %v, Want: %v", args...)
	}
}
