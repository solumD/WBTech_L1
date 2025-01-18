package main

/* Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(22+32+42….) с использованием конкурентных вычислений. */

import (
	"fmt"
	"sync"
)

func Squares(nums []int) chan int {
	wg := &sync.WaitGroup{}
	squaresOut := make(chan int)

	for _, v := range nums {
		wg.Add(1)
		go func() {
			defer wg.Done()
			squaresOut <- v * v
		}()
	}

	go func() {
		defer close(squaresOut) // закрываем канал
		wg.Wait()               // дожидаемся завершения горутин
	}()

	return squaresOut
}

func SquaresSum(squaresIn chan int) int {
	sum := 0

	// читаем значения квадратов из канала, пока он не будет закрыт
	for v := range squaresIn {
		sum += v
	}

	return sum
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	squares := Squares(nums)
	sum := SquaresSum(squares)

	fmt.Printf("Сумма квадратов: %d\n", sum)
}
