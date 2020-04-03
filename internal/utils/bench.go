package utils

import (
	"time"
)

func BenchFunc(s string) (string, time.Time) {
	return s, time.Now()
}

func Track(name string, start time.Time) {
	// fmt.Printf("[%s] took %s to finish!\n", name, time.Since(start))
}
