package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

/* К каким негативным последствиям может привести
данный фрагмент кода, и как это исправить? Приведите
 корректный пример реализации. */

/* var justString string // глобальная переменная, не очень хорошая практика

// непонятное название, из которого нельзя понял суть функции
func someFunc() { // функция ничего не вовращает
  v := createHugeString(1 << 10) // функция не объявлена
  justString = v[:100] // длина строки может быть меньше 100
}

func main() {
  someFunc() // вызов функции, которая ничего не возвращает - бесмысленно
} */

const alphabet = "abcdefghijklmnopqrstuvwxyz" // алфавит

func CreateString(n int) string {
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder // строку собираем через strings builder для оптимизации
	for i := 0; i < n; i++ {
		idx := rand.Intn(len(alphabet))
		sb.WriteByte(alphabet[idx]) // получаем случ. элемент
	}

	return sb.String()
}

func TrimRight(str string, rBorder int) (string, error) {
	if len(str) < rBorder {
		return str, fmt.Errorf("right border is greater than string's length")
	}

	return str[:rBorder], nil
}

func main() {
	str := CreateString(1 << 10)

	str, err := TrimRight(str, 100)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(str)
}
