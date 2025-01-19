package main

import "fmt"

func main() {
	measures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	hm := make(map[int][]float64)

	for _, v := range measures {
		key := int(v) / 10 * 10      // рассчитываем шаг
		hm[key] = append(hm[key], v) // добавляем измерение в мапу по ключу (шагу)
	}

	// вывод подмножеств
	for k, v := range hm {
		fmt.Printf("%d: %v\n", k, v)
	}
}
