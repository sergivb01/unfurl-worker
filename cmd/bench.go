package main

import (
	"fmt"
	"time"
)

func BenchmarkFunction(start time.Time, name string){
	fmt.Printf("[%s] took %s to finish!\n", name, time.Since(start))
}
