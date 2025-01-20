package main

import "fmt"

/* Удалить i-ый элемент из слайса. */

func DeleteIth(arr []int, idx int) ([]int, error) {
	if idx >= len(arr) || idx < 0 {
		return arr, fmt.Errorf("invalid index")
	}

	// удаляем через соединение срезов
	arr = append(arr[:idx], arr[idx+1:]...)
	return arr, nil
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr)
	arr, err := DeleteIth(arr, 3)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(arr)
}
