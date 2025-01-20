package main

import (
	"fmt"
	"strings"
)

/* Разработать программу, которая проверяет, что все символы
в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой. Например:
abcd — true, abCdefAaf — false, aabcd — false */

func CheckUniq(str string) bool {
	set := make(map[rune]struct{}) // используем сет, чтобы сохранять встреченные символы

	for _, v := range strings.ToLower(str) { // приводим к нижнему регистру, чтобы ф-ия была регистронезависимой
		if _, exist := set[v]; exist {
			return false
		}

		set[v] = struct{}{} // отмечаем встреченный символ
	}

	return true
}

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"

	fmt.Printf("СТрока: %s | Результат: %t\n", s1, CheckUniq(s1)) // true
	fmt.Printf("СТрока: %s | Результат: %t\n", s2, CheckUniq(s2)) // false
	fmt.Printf("СТрока: %s | Результат: %t\n", s3, CheckUniq(s3)) // false
}
