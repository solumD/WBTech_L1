package main

import (
	"context"
	"fmt"
	"time"
)

/* Реализовать все возможные способы остановки выполнения горутины. */

func Writer1(ctx context.Context) chan int64 {
	out := make(chan int64)

	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Closed Writer1 by context")
				return
			case out <- time.Now().Unix():
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	return out
}

func Reader1(in chan int64) {

	for v := range in {
		fmt.Println(v)
	}

}

func Writer2(cancel chan struct{}) chan int64 {
	out := make(chan int64)

	go func() {
		defer close(out)
		for {
			select {
			case <-cancel: // завершаем горутину при чтении пустой структуры
				fmt.Println("Closed Writer2 by cancel-chan")
				return
			case out <- time.Now().Unix():
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	return out
}

func Reader2(in chan int64) {
	for v := range in {
		fmt.Println(v)
	}

}

func main() {
	// через отмену контекста/контекст с таймаутом
	const n = 5

	ctx, cancel := context.WithTimeout(context.Background(), n*time.Second)
	defer cancel()

	out1 := Writer1(ctx)
	Reader1(out1)

	// через канал отмены
	cancelChan := make(chan struct{})
	go func() {
		time.Sleep(n * time.Second)
		cancelChan <- struct{}{} // отправляем пустую структуру в канал отмены
	}()

	out2 := Writer2(cancelChan)
	Reader2(out2)
}
