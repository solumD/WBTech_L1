package main

import (
	"fmt"
	"math"
)

/* Разработать программу нахождения расстояния между
двумя точками, которые представлены в виде структуры
Point с инкапсулированными параметрами x,y и конструктором. */

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func CountDistance(p1 *Point, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func main() {
	p1 := NewPoint(5.5, 7)
	p2 := NewPoint(-4.3, -8)

	fmt.Printf("Расстояние между точками: %.2f\n", CountDistance(p1, p2))
}
