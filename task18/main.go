package main

import (
	"fmt"
	"sync"
)

/* Реализовать структуру-счетчик, которая будет инкрементироваться
в конкурентной среде. По завершению программа должна выводить
итоговое значение счетчика. */

type Counter struct {
	mu    *sync.Mutex
	count int
}

func NewCounter() *Counter {
	return &Counter{mu: &sync.Mutex{}, count: 0}
}

func (c *Counter) Increment() {
	// мьютексом делаем закрытие на кокурентную запись
	c.mu.Lock()
	defer c.mu.Unlock()

	c.count++
}

func main() {
	counter := NewCounter()
	wg := &sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()

	fmt.Printf("Итоговое значение счетчика: %d\n", counter.count) // 30000
}
