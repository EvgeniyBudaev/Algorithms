package main

/* 346. Moving Average from Data Stream

Учитывая поток целых чисел и размер окна, вычислите скользящее среднее всех целых чисел в скользящем окне.

* For example,
*
* MovingAverage m = new MovingAverage(3);
* m.next(1) = 1
* m.next(10) = (1 + 10) / 2
* m.next(3) = (1 + 10 + 3) / 3
* m.next(5) = (10 + 3 + 5) / 3

Input: ["MovingAverage", "next", "next", "next", "next"]
[[3], [1], [10], [3], [5]]
Output: [null, 1.0, 5.5, 4.66667, 6.0]
*/

func main() {
	movingAverage := Constructor(3)

	println(movingAverage.Next(1))  // 1.0
	println(movingAverage.Next(10)) // 5.5
	println(movingAverage.Next(3))  // 4.666666666666667
	println(movingAverage.Next(5))  // 6.0
}

type MovingAverage struct {
	size  int
	queue []int
	sum   float64
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		size:  size,
		queue: make([]int, 0, size),
		sum:   0.0,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.queue) >= this.size {
		removedVal := this.queue[0]
		this.queue = this.queue[1:]
		this.sum -= float64(removedVal)
	}

	this.queue = append(this.queue, val)
	this.sum += float64(val)

	return this.sum / float64(len(this.queue))
}
