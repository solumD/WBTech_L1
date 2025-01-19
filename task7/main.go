package main

import (
	"fmt"
	"sync"
)

/* Реализовать конкурентную запись данных в map. */

type SyncMap struct {
	vals map[int]int
	mu   *sync.Mutex
}

func New() *SyncMap {
	return &SyncMap{
		vals: map[int]int{},
		mu:   &sync.Mutex{},
	}
}

func (m *SyncMap) Set(key, val int) {
	// закрываем мьютекс для безопасной записи
	m.mu.Lock()
	defer m.mu.Unlock()

	m.vals[key] = val
	fmt.Printf("Get key: %d val: %d\n", key, val)
}

func Worker(wg *sync.WaitGroup, syncMap *SyncMap, key int) {
	defer wg.Done()
	for i := 0; i < key; i++ {
		syncMap.Set(key, i+1)
	}
}

func main() {
	sMap := New()
	wg := &sync.WaitGroup{}

	wg.Add(2)
	// запускаем два воркера для теста
	go Worker(wg, sMap, 100)
	go Worker(wg, sMap, 200)

	wg.Wait()
	fmt.Println(sMap.vals)
}
