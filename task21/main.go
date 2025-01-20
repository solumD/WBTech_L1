package main

import "fmt"

/* Реализовать паттерн «адаптер» на любом примере. */

// Интерфейс компьютера с lighting
type Computer interface {
	InsertIntoLightning()
}

// Клиент, который выполненяет соединение с lighting
type Client struct {
}

// метод имеет какую-то полезную логику, но может взаимодействовать только
// с объектами интерфейса Computer, поэтому при необходимости нужно написать
// адаптер для других объектов
func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Клиент подключает lightning-провод к компьютеру.")
	com.InsertIntoLightning()
}

// удовлетворяет интерфейсу Computer
type Mac struct {
}

func (m *Mac) InsertIntoLightning() {
	fmt.Println("Провод lightning подключен к mac.")
}

// не удовлетворяет интерфейсу Computer
type Windows struct {
}

func (w *Windows) insertIntoUSB() {
	fmt.Println("Провод USB подключен к windows.")
}

// адаптирует тип Windows, чтобы он мог работать с клиентом
type WindowsAdapter struct {
	windowsComputer *Windows
}

func (a *WindowsAdapter) InsertIntoLightning() {
	fmt.Println("Адаптер конвертирует lightning-провод в USB-провод.")
	a.windowsComputer.insertIntoUSB()
}

func main() {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	fmt.Println()

	windowsPC := &Windows{}
	windowsAdapter := &WindowsAdapter{
		windowsComputer: windowsPC,
	}

	client.InsertLightningConnectorIntoComputer(windowsAdapter)
}
