package main

import "fmt"

/* Дана переменная int64. Разработать программу
которая устанавливает i-й бит в 1 или 0. */

func main() {
	// считываем значения
	fmt.Printf("Number: ")
	var n int64
	fmt.Scan(&n)
	fmt.Printf("Position: ")
	var p int
	fmt.Scan(&p)
	fmt.Printf("Value (0 or 1): ")
	var v int
	fmt.Scan(&v)

	// меняем бит с помощью маски
	if v == 1 {
		n = n | (1 << p)
	} else if v == 0 {
		n = n & ^(1 << p)
	} else {
		fmt.Println("Invalid val")
	}

	fmt.Printf("Result %b: %d\n", n, n)
}
