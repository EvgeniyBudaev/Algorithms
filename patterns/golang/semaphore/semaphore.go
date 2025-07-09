package main

import (
	"log"
	"sync"
	"time"
)

/*
Паттерн Семафор — при создании множества горутин для обработки запросов может возникнуть проблема с общим доступом к
ресурсам, таким как база данных.
Паттерн Семафор поможет избежать отправки неограниченного количества запросов к базе данных.
*/

// Semaphore структура семафора
type Semaphore struct {
	semaCh chan struct{}
}

// NewSemaphore создает семафор с буферизованным каналом емкостью maxReq
func NewSemaphore(maxReq int) *Semaphore {
	return &Semaphore{
		semaCh: make(chan struct{}, maxReq),
	}
}

// когда горутина запускается, отправляем пустую структуру в канал semaCh
func (s *Semaphore) Acquire() {
	s.semaCh <- struct{}{}
}

// когда горутина завершается, из канала semaCh убирается пустая структура
func (s *Semaphore) Release() {
	<-s.semaCh
}

func main() {
	// чтобы дождаться всех горутин
	var wg sync.WaitGroup

	// создаём семафор емкостью 2: он будет пропускать только 2 горутины
	semaphore := NewSemaphore(2)

	// создаем 10 горутин
	for idx := 0; idx < 10; idx++ {
		wg.Add(1)

		// горутина в которую помещаем её порядковый номер
		go func(taskID int) {
			// отправляем в канал семафора пустую структуру
			semaphore.Acquire()

			// откладываем уменьшение счетчика в WaitGroup, когда завершится горутина
			defer wg.Done()

			// забираем из канала семафора пустую структуру, дав возможность запуститься другим горутинам
			defer semaphore.Release()

			log.Println("Запущен рабочий %d", taskID)

			time.Sleep(1 * time.Second)
		}(idx)
	}

	// ждём завершения всех горутин
	wg.Wait()
}
