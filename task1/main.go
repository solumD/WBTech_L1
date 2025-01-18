package main

import "fmt"

// Структура Human
type Human struct {
	Age     int
	Name    string
	Surname string
	Gender  string
}

func (h *Human) PrintFullName() {
	fmt.Printf("Name: %s, Surname: %s\n", h.Name, h.Surname)
}

func (h *Human) PrintAgeAndGender() {
	fmt.Printf("Age: %d, Gender: %s\n", h.Age, h.Gender)
}

// Структура Action
type Action struct {
	Human // встраиваем структуру Human в Action
}

func (a *Action) Walk() {
	fmt.Printf("Human %s %s with age %d and gender %s is walking\n", a.Name, a.Surname, a.Age, a.Gender)
}

func main() {
	h := Human{
		Name:    "Dmitry",
		Surname: "Kononov",
		Age:     19,
		Gender:  "male",
	}

	a := Action{h}

	// методы структуры Human
	a.PrintFullName()
	a.PrintAgeAndGender()

	fmt.Println()

	// методы структуры Action
	a.Walk()
}
