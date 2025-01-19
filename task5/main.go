package main

import (
	"context"
	"fmt"
	"time"
)

/* Разработать программу, которая будет последовательно
отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться. */

func Writer(ctx context.Context) chan int64 {
	out := make(chan int64)

	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case out <- time.Now().Unix():
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	return out
}

func Reader(in chan int64) {

	for v := range in {
		fmt.Println(v)
	}

}

func main() {
	const n = 5

	// через контекст ограничиваем время работы программы до n секунд
	ctx, cancel := context.WithTimeout(context.Background(), n*time.Second)
	defer cancel()

	out := Writer(ctx)
	Reader(out)
}
