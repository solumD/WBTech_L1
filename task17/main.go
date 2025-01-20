package main

import "fmt"

// BinarySearch определяет индекс указанного значения,
// и возвращает его, иначе возвращает -1
func BinarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + ((right - left) / 2)

		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{15, 20, 5, 7, 45, 3, 100, 6}

	idx := BinarySearch(arr, 7)
	if idx == -1 {
		fmt.Println("Число не найдено")
	} else {
		fmt.Printf("Индекс числа: %d\n", idx)
	}
}
