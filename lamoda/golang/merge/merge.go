package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// merge объединяет несколько входных каналов в один выходной.
// time: O(n), space: O(n), где n - количество каналов
func merge(ch ...<-chan int) <-chan int {
	out := make(chan int) // Выходной канал
	var wg sync.WaitGroup

	// Для каждого входного канала запускаем горутину, которая читает из него
	for _, c := range ch {
		wg.Add(1)
		go func(input <-chan int) {
			defer wg.Done()
			for val := range input {
				out <- val
			}
		}(c)
	}

	// Запускаем горутину для закрытия выходного канала после завершения всех читателей
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// source генерирует случайные числа.
// time: O(n), space: O(n), где n - количество каналов
func source(sourceFunc func(int) int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < 3; i++ {
			ch <- sourceFunc(i)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		}
	}()

	return ch
}

func main() {
	rand.Seed(time.Now().UnixMilli()) // Инициализация генератора случайных чисел

	// Генерируем два канала
	in1 := source(func(_ int) int {
		return rand.Int()
	})

	in2 := source(func(i int) int {
		return i
	})

	// Объединяем два канала
	out := merge(in1, in2)

	// Печатаем значения из выходного канала
	for val := range out {
		fmt.Println("Value: ", val)
	}
}
