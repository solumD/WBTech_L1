package main

import (
	"fmt"
	"math"
	"math/big"
)

/* Разработать программу, которая перемножает, делит,
складывает, вычитает две числовых переменных a и b,
значение которых > 2^20. */

func main() {
	a, b := big.NewFloat(math.Pow(2, 30)), big.NewFloat(math.Pow(2, 25))
	fmt.Printf("Число 1: %f | Число 2: %f\n", a, b)
	fmt.Println("Результат умножения -", new(big.Float).Mul(a, b))
	fmt.Println("Результат деления -", new(big.Float).Quo(a, b))
	fmt.Println("Результат сложения -", new(big.Float).Add(a, b))
	fmt.Println("Результат вычитания -", new(big.Float).Sub(a, b))
}
