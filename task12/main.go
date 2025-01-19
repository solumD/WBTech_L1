package main

import "fmt"

/* Имеется последовательность строк - (cat, cat, dog, cat, tree)
создать для нее собственное множество. */

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем множества по каждому уникальному слову (hashSet).
	// Значения не важны, по этому в качестве типа указываем
	// пустую структуру, она ничего не весит по памяти.
	set := make(map[string]struct{})

	for _, v := range arr {
		set[v] = struct{}{}
	}

	fmt.Println(set)
}
