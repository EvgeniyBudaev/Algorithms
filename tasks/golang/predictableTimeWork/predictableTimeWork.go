package main

import (
	"math/rand"
	"time"
)

// randomTimeWork функция, которая работает неопределенно долго (до 100 секунд)
func randomTimeWork() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Second)
}

// написать обертку для этой функции, которая будет прерывать выполнение, если функция работает дольше 3 секунд
func predictableTimeWork() {
	ch := make(chan struct{})

	go func() {
		randomTimeWork()
		close(ch)
	}()

	select {
	case <-ch:
	case <-time.After(3 * time.Second):
	}
}

func main() {
	predictableTimeWork()
}
