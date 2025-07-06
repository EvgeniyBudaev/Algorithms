package main

import (
	"log"
	"time"
)

func main() {
	// Слайс данных
	input := []int{1, 2, 3, 4, 5, 6}

	// Отправляем данные в функцию handler с явной отменой
	handler(input)

	time.Sleep(time.Second)
}

// handler получает данные из слайса
func handler(input []int) {
	// канал для явной отмены
	doneCh := make(chan struct{})

	// когда выходим из handler — сразу закрываем канал doneCh
	defer close(doneCh)

	// Получаем канал с данными из генератора
	inputCh := generator(doneCh, input)

	// забираем данные из канала
	for data := range inputCh {
		// если в данных 3 — выходим из handler
		if data == 3 {
			log.Println("Прекращаем обработку данных из канала")
			return
		}
		log.Println(data)
	}
	log.Println("Данные во входном канале закончились")
}

// generator отправляет данные из слайса в канал, а потом его возвращает.
func generator(doneCh chan struct{}, input []int) <-chan int {
	// Канал, в который будем отправлять данные из слайса
	inputCh := make(chan int)

	// Через отдельную горутину генератор отправляет данные в канал
	go func() {
		// Закрываем канал по завершению горутины
		defer close(inputCh)

		// Перебираем данные в слайсе
		for _, data := range input {
			select {
			// Если канал doneCh закрылся - сразу выходим из горутины
			case <-doneCh:
				log.Println("Останавливаем генератор")
				return
			// Отправляем данные в канал inputCh
			case inputCh <- data:
			}
		}
	}()

	// Возвращаем канал с данными
	return inputCh
}
