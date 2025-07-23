package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println(callLongFunction(ctx)) // 42 <nil> или 0 timeout
}

// callLongFunction - функция, которая вызывает longFunction.
// Если longFunction завершается не успев, то она должна отменить выполнение.
func callLongFunction(ctx context.Context) (int, error) {
	result := make(chan int)

	go func() {
		defer close(result)
		result <- longFunction()
	}()

	for {
		select {
		case v := <-result:
			return v, nil
		case <-ctx.Done():
			return 0, errors.New("timeout")
		}
	}
}

// longFunction - функция, которая работает долго.
func longFunction() int {
	time.Sleep(time.Duration(rand.N[int64](5)) * time.Second)
	return 42
}
