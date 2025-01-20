package main

import (
	"fmt"
	"time"
)

/* Реализовать собственную функцию sleep. */

func MySleep(d time.Duration) {
	start := time.Now()
	// бесконечный цикл, пока time.Since меньше d
	for time.Since(start) < d {
	}
}

func main() {
	MySleep(5 * time.Second)
	fmt.Println("5 seconds passed")
}
