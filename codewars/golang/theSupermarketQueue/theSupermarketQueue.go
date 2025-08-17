package main

import (
	"fmt"
	"sort"
)

/* The Supermarket Queue
https://www.codewars.com/kata/57b06f90e298a7b53d000a86/train/go

В супермаркете есть очередь к кассам самообслуживания. Ваша задача — написать функцию для расчета общего времени,
необходимого всем покупателям для оплаты!

входные данные
клиенты: массив положительных целых чисел, представляющий очередь. Каждое целое число представляет покупателя, а его
значение — время, необходимое ему для оплаты.
n: положительное целое число, количество касс самообслуживания.
выходные данные
Функция должна возвращать целое число, представляющее общее затраченное время.
*/

func main() {
	fmt.Println(QueueTime([]int{}, 1))                 // 0
	fmt.Println(QueueTime([]int{1, 2, 3, 4}, 1))       // 10
	fmt.Println(QueueTime([]int{2, 2, 3, 3, 4, 4}, 2)) // 9
	fmt.Println(QueueTime([]int{1, 2, 3, 4, 5}, 100))  // 5
}

// QueueTime находит общее затраченное время.
// time: O(n), space: O(n), n - количество покупателей
func QueueTime(customers []int, n int) int {
	queueTime := make([]int, n)

	for _, customerTime := range customers {
		queueTime[0] += customerTime
		sort.Ints(queueTime)
	}

	return queueTime[n-1]
}
