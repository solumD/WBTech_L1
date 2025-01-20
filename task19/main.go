package main

import (
	"fmt"
	"strings"
)

/* Разработать программу, которая переворачивает подаваемую
на ход строку (например: «главрыба — абырвалг»). Символы
могут быть unicode. */

func ReverseString(str string) string {
	var sb strings.Builder
	strInRunes := []rune(str) // используем руны, чтобы работать со всеми символами unicode

	for i := len(strInRunes) - 1; i >= 0; i-- {
		sb.WriteRune(strInRunes[i])
	}

	return sb.String()
}

func main() {
	var str string
	fmt.Print("Введите строку: ")
	fmt.Scan(&str)
	fmt.Printf("Перевернутая строка: %s\n", ReverseString(str))

	/*str1 := "главрыба"
	str2 := "mainfish"

	fmt.Printf("Строка:  %s | В перевернутом виде: %s\n", str1, ReverseString(str1))
	fmt.Printf("Строка:  %s | В перевернутом виде: %s\n", str2, ReverseString(str2))*/
}
