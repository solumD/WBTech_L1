package main

/* Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные
из канала и выводят в stdout. Необходима возможность выбора количества
воркеров при старте. Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров. */

import (
	"context"
	"flag"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Worker(id int, data chan string, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case val, ok := <-data:
			if !ok {
				fmt.Println("stopping worker")
				return
			}

			fmt.Printf("worker id: %d | recieved data: %s\n", id, val)
		}
	}
}

func main() {
	// флаг, в котором указывается кол-во воркеров, по умолчанию 3
	workersCount := flag.Int("workers", 3, "sets number of workers")
	flag.Parse()

	// канал, в который отправляются данные
	data := make(chan string)
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	wg := sync.WaitGroup{}
	for i := 0; i < *workersCount; i++ {
		wg.Add(1)
		// запускаем воркер
		go func() {
			defer wg.Done()
			Worker(i+1, data, ctx)
		}()
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				// в случае остановки программы закрываем канал
				close(data)
				return
			case data <- time.Now().String():
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// дожидаемся завершения воркеров
	wg.Wait()

	fmt.Println("work is finished")
}
