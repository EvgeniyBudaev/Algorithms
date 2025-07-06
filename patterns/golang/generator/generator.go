package main

import "fmt"

func main() {
	// Данные в слайсе, которые будем отправлять
	input := []int{1, 2, 3, 4, 5, 6}

	// Получаем канал с данными из генератора
	inputCh := generator(input)

	// Отправляем данные потребителю через канал inputCh
	consumer(inputCh)
}

// generator создает канал и сразу возвращает его.
func generator(input []int) <-chan int {
	inputCh := make(chan int)

	// Через отдельную горутину генератор отправляет данные в канал
	go func() {
		// Закрываем канал по завершению горутины — это отправитель
		defer close(inputCh)

		// Перебираем данные в слайсе
		for _, data := range input {
			// Отправляем данные в канал inputCh
			inputCh <- data
		}
	}()

	// Возвращаем канал inputCh
	return inputCh
}

// consumer проходит через канал и одновременно обрабатывает данные из него (выводит на экран)
func consumer(input <-chan int) {
	for data := range input {
		fmt.Println(data)
	}
}
