package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/* Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow». */

func main() {
	fmt.Print("Введите строку: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	strFields := strings.Fields(str)
	newStr := make([]string, 0, len(strFields))

	for i := len(strFields) - 1; i >= 0; i-- {
		newStr = append(newStr, strFields[i])
	}

	fmt.Printf("Новая строка: %s\n", strings.Join(newStr, " "))
}
