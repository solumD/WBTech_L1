package main

import (
	"fmt"
	"sync"
)

/* Написать программу, которая конкурентно рассчитает значение квадратов
чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout. */

func Squares(nums []int) {
	wg := &sync.WaitGroup{}

	for _, v := range nums {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(v * v)
		}()
	}

	wg.Wait() // дожидаемся завершения горутин
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	Squares(nums)
}
