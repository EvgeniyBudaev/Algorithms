package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Паттерн Fan-Out — создаёт большее количество горутин на определенном этапе, что увеличивает его пропускную способность.
Если один из этапов конвейера занимает больше времени и ресурсов, чем другие, работа всех остальных этапов будет
заблокирована. Чтобы избежать этого, можно использовать паттерн Fan-Out.

Паттерн Fan-In - объединяет несколько результатов в один канал. Этот процесс также называют мультиплексированием.
*/
func main() {
	// ваши данные в слайсе
	input := []int{1, 2, 3}

	// канал для сигнала к выходу из горутины
	doneCh := make(chan struct{})
	// при завершении программы закрываем канал doneCh, чтобы все горутины тоже завершились
	defer close(doneCh)

	// получаем канал с данными с помощью генератора
	inputCh := generator(doneCh, input)

	// получаем слайс каналов из 10 рабочих add
	channels := fanOut(doneCh, inputCh)

	// а теперь объединяем десять каналов в один
	addResultCh := fanIn(doneCh, channels...)

	// передаём тот один канал в следующий этап обработки
	resultCh := multiply(doneCh, addResultCh)

	// выводим результаты расчетов из канала
	for res := range resultCh {
		fmt.Println(res)
	}
}

// fanIn объединяет несколько каналов resultChs в один.
func fanIn(doneCh chan struct{}, resultChs ...chan int) chan int {
	// конечный выходной канал в который отправляем данные из всех каналов из слайса, назовём его результирующим
	finalCh := make(chan int)

	// понадобится для ожидания всех горутин
	var wg sync.WaitGroup

	// перебираем все входящие каналы
	for _, ch := range resultChs {
		// в горутину передавать переменную цикла нельзя, поэтому делаем так
		chClosure := ch

		// инкрементируем счётчик горутин, которые нужно подождать
		wg.Add(1)

		go func() {
			// откладываем сообщение о том, что горутина завершилась
			defer wg.Done()

			// получаем данные из канала
			for data := range chClosure {
				select {
				// выходим из горутины, если канал закрылся
				case <-doneCh:
					return
				// если не закрылся, отправляем данные в конечный выходной канал
				case finalCh <- data:
				}
			}
		}()
	}

	go func() {
		// ждём завершения всех горутин
		wg.Wait()
		// когда все горутины завершились, закрываем результирующий канал
		close(finalCh)
	}()

	// возвращаем результирующий канал
	return finalCh
}

// fanOut принимает канал данных, порождает 10 горутин
func fanOut(doneCh chan struct{}, inputCh chan int) []chan int {
	// количество горутин add
	numWorkers := 10
	// каналы, в которые отправляются результаты
	channels := make([]chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		// получаем канал из горутины add
		addResultCh := add(doneCh, inputCh)
		// отправляем его в слайс каналов
		channels[i] = addResultCh
	}

	// возвращаем слайс каналов
	return channels
}

// add принимает на вход сигнальный канал для прекращения работы и канал с входными данными для работы,
// а возвращает канал, в который будет отправляться результат вычислений.
// На фоне будет запущена горутина, выполняющая вычисления до момента закрытия doneCh.
func add(doneCh chan struct{}, inputCh chan int) chan int {
	// канал с результатом
	addRes := make(chan int)

	// горутина, в которой добавляем к значению из inputCh единицу и отправляем результат в addRes
	go func() {
		// закрываем канал, когда горутина завершается
		defer close(addRes)

		// берём из канала inputCh значения, которые надо изменить
		for data := range inputCh {
			// замедлим вычисление, как будто функция add требует больше вычислительных ресурсов
			time.Sleep(time.Second)
			result := data + 1

			select {
			// если канал doneCh закрылся, выходим из горутины
			case <-doneCh:
				return
			// если doneCh не закрыт, отправляем результат вычисления в канал результата
			case addRes <- result:
			}
		}
	}()

	// возвращаем канал для результатов вычислений
	return addRes
}

// multiply принимает на вход сигнальный канал для прекращения работы и канал с входными данными для работы,
// а возвращает канал, в который будет отправляться результат вычислений.
// На фоне будет запущена горутина, выполняющая вычисления до момента закрытия doneCh.
func multiply(doneCh chan struct{}, inputCh chan int) chan int {
	// канал с результатом
	multiplyRes := make(chan int)

	// горутина, в которой значение из inputCh умножаем на 2 и возвращаем в канал multiplyRes
	go func() {
		// закрываем канал multipleRes, когда горутина завершается
		defer close(multiplyRes)

		// берем из канала inputCh значения, которые надо изменить
		for data := range inputCh {
			// изменяем данные
			result := data * 2

			select {
			// если канал doneCh закрылся, выходим из горутины
			case <-doneCh:
				return
			// если doneCh не закрыт, отправляем результат вычисления в канал результата
			case multiplyRes <- result:
			}
		}
	}()

	// возвращаем канал для результатов вычислений
	return multiplyRes
}

// generator возвращает канал с данными
func generator(doneCh chan struct{}, input []int) chan int {
	// канал, в который будем отправлять данные из слайса
	inputCh := make(chan int)

	// горутина, в которой отправляем в канал inputCh данные
	go func() {
		// как отправители закрываем канал, когда всё отправим
		defer close(inputCh)

		// перебираем все данные в слайсе
		for _, data := range input {
			select {
			// если doneCh закрыт, сразу выходим из горутины
			case <-doneCh:
				return
			// если doneCh не закрыт, кидаем в канал inputCh данные data
			case inputCh <- data:
			}
		}
	}()

	// возвращаем канал для данных
	return inputCh
}
