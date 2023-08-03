package utils

import "log"

func Check(err error, msg string, args ...any) {
	if err != nil {
		log.Printf(msg, args...)
		panic(err)
	}
}
