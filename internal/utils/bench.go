package utils

import (
	"fmt"
	"time"
)

func StartBench() time.Time {
	return time.Now()
}

func BenchmarkFunction(start time.Time, name string) {
	fmt.Printf("[%s] took %s to finish!\n", name, time.Since(start))
}
