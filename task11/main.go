package main

import "fmt"

/* Реализовать пересечение двух неупорядоченных множеств. */

func main() {
	// реализовано с учетом того, что в множестве
	// значения как правило не повторяются

	arr1 := []int{3, 7, 2, 9, 5, 15, 4, 11}
	arr2 := []int{12, 4, 8, 15, 6, 7}

	crossing := []int{}
	hm := make(map[int]struct{})

	// сохраняем значения множества 1
	for _, v := range arr1 {
		hm[v] = struct{}{}
	}

	// проходим по значениям множества 2 и, если какое-то уже встречалось,
	// добавляем его в пересечение
	for _, v := range arr2 {
		if _, exist := hm[v]; exist {
			crossing = append(crossing, v)
		}
	}

	fmt.Println(crossing)
}
