package main

import (
	"fmt"
	"reflect"
)

// через форматирование
func GetTypeFmt(variable interface{}) {
	fmt.Printf("Тип переменной: %T\n", variable)
}

// через оператор .type и reflect
func GetType(variable interface{}) {
	switch variable.(type) {
	case int:
		fmt.Println("Тип переменной: int")
	case string:
		fmt.Println("Тип переменной: string")
	case bool:
		fmt.Println("Тип переменной: bool")
	default:
		// получаем тип элемента и сраниваем
		if reflect.TypeOf(variable).Kind() == reflect.Chan {
			fmt.Println("Тип переменной: channel")
		} else {
			fmt.Println("Неизвестный тип")
		}
	}
}

func main() {
	num := 10
	str := "10"
	ok := true
	someChan := make(chan int)

	GetType(num)
	GetType(str)
	GetType(ok)
	GetType(someChan)

	fmt.Println()

	GetTypeFmt(num)
	GetTypeFmt(str)
	GetTypeFmt(ok)
	GetTypeFmt(someChan)

	fmt.Println()

	arr := []int{}
	GetType(arr)
	GetTypeFmt(arr)
}
