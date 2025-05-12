package main

import (
	"container/heap"
	"fmt"
)

/* 215. Kth Largest Element in an Array
https://leetcode.com/problems/kth-largest-element-in-an-array/description/

Дан целочисленный массив nums и целое число k, вернуть k-й наибольший элемент в массиве.
Обратите внимание, что это k-й наибольший элемент в отсортированном порядке, а не k-й отдельный элемент.
Можете ли вы решить эту задачу без сортировки?

Input: nums = [3,2,1,5,6,4], k = 2
Output: 5
*/

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(findKthLargest(nums, k)) // 5
}

// MinHeap используется для хранения элементов, чтобы найти k-й наибольший элемент.
type MinHeap struct {
	array []int // используем массив для хранения элементов кучи
}

// Len - возвращает длину массива.
// time: O(1), space: O(1)
func (m *MinHeap) Len() int {
	return len(m.array)
}

// Less - возвращает true, если элемент с индексом i меньше элемента с индексом j.
// time: O(1), space: O(1)
func (m *MinHeap) Less(i, j int) bool {
	return m.array[i] < m.array[j]
}

// Swap - меняет местами элементы с индексами i и j.
// time: O(1), space: O(1)
func (m *MinHeap) Swap(i, j int) {
	m.array[i], m.array[j] = m.array[j], m.array[i]
}

// Push - добавляет элемент в кучу.
// time: O(1), space: O(1)
func (m *MinHeap) Push(num any) {
	m.array = append(m.array, num.(int))
}

// Pop - удаляет элемент из кучи.
// time: O(1), space: O(1)
func (m *MinHeap) Pop() any {
	l := len(m.array)
	last := m.array[l-1]
	m.array = m.array[:l-1]
	return last
}

// findKthLargest - находит k-й наибольший элемент в массиве.
// time: O(n*log(k)), space: O(k)
func findKthLargest(nums []int, k int) int {
	h := &MinHeap{} // создаем кучу
	heap.Init(h)    // инициализируем кучу
	for _, num := range nums {
		heap.Push(h, num) // добавляем элемент в кучу
		if h.Len() > k {  // если количество элементов в куче больше k, удаляем наименьший элемент
			_ = heap.Pop(h) // удаляем наименьший элемент
		}
	}
	return heap.Pop(h).(int) // возвращаем наибольший элемент в куче
}

// findKthLargest - находит k-й наибольший элемент в массиве.
// time: O(n*log(n)), space: O(n)
//func findKthLargest(nums []int, k int) int {
//	sort.Ints(nums)
//	return nums[len(nums)-k]
//}
