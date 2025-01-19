package main

import (
	"fmt"
)

/* Разработать конвейер чисел. Даны два канала: в первый
пишутся числа (x) из массива, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout. */

func Squares(in chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			out <- v * v
		}
	}()

	return out
}

func PrintSquares(in chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	in := make(chan int)
	go func() {
		defer close(in)
		for _, n := range nums {
			in <- n
		}
	}()

	out := Squares(in)
	PrintSquares(out)
}
