package main

import "fmt"

/* Реализовать быструю сортировку массива
(quicksort) встроенными методами языка. */

func Quicksort(arr []int) []int {
	return qs(arr, 0, len(arr)-1)
}

func qs(arr []int, low int, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = qs(arr, low, p-1)
		arr = qs(arr, p+1, high)
	}

	return arr
}

// определяем партитию
func partition(arr []int, low, high int) ([]int, int) {
	mid := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < mid {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]

	return arr, i
}

func main() {
	arr := []int{15, 20, 5, 7, 45, 3, 100, 6}
	fmt.Println(arr)
	Quicksort(arr)
	fmt.Println(arr)
}
