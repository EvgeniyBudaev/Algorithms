package main

import "fmt"

/* https://www.lintcode.com/problem/3662/description

В этом вопросе вам нужно создать класс HitCounter.
В этом классе имеются следующие функции:

HitCounter(): конструктор без аргументов
void hit(int timestamp): указывает, что касание происходит в указанное время.
int getHits(int timestamp): возвращает общее количество обращений за 300 секунд до указанного времени.
Где временная метка находится в секундах.

Input:
["HitCounter", "hit", "hit", "hit", "getHits", "hit", "getHits", "getHits"]
[[], [1], [2], [3], [4], [300], [300], [301]]
Output: [null, null, null, null, 3, null, 4, 3]

Explanation
HitCounter hitCounter = new HitCounter();
hitCounter.hit(1);       // hit at timestamp 1.
hitCounter.hit(2);       // hit at timestamp 2.
hitCounter.hit(3);       // hit at timestamp 3.
hitCounter.getHits(4);   // get hits at timestamp 4, return 3.
hitCounter.hit(300);     // hit at timestamp 300.
hitCounter.getHits(300); // get hits at timestamp 300, return 4.
hitCounter.getHits(301); // get hits at timestamp 301, return 3.
*/

type HitCounter struct {
	ts []int
}

func NewHitCounter() *HitCounter {
	return &HitCounter{ts: make([]int, 0)}
}

func (hc *HitCounter) Hit(timestamp int) {
	hc.ts = append(hc.ts, timestamp)
}

func (hc *HitCounter) GetHits(timestamp int) int {
	// Функция для поиска первого индекса, который больше или равен заданному значению
	search := func(x int) int {
		l, r := 0, len(hc.ts)
		for l < r {
			mid := (l + r) / 2
			if hc.ts[mid] >= x {
				r = mid
			} else {
				l = mid + 1
			}
		}
		return l
	}

	// Вычисляем границу временного окна
	boundary := timestamp - 300 + 1
	if boundary < 0 {
		boundary = 0
	}

	// Находим количество хитов за последние 300 секунд
	return len(hc.ts) - search(boundary)
}

func main() {
	hitCounter := NewHitCounter()
	hitCounter.Hit(1)
	hitCounter.Hit(2)
	hitCounter.Hit(3)
	fmt.Println(hitCounter.GetHits(4)) // Output: 3
	hitCounter.Hit(300)
	fmt.Println(hitCounter.GetHits(300)) // Output: 4
	fmt.Println(hitCounter.GetHits(301)) // Output: 3
}
