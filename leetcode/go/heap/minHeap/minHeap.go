package main

import "fmt"

// MinHeap - min heap
type MinHeap struct {
	arr []int // Массив для хранения элементов кучи
}

func NewMinHeap() *MinHeap {
	return &MinHeap{arr: make([]int, 0)}
}

// Push - добавить элемент в кучу.
// time: O(log(n)), space: O(1)
func (h *MinHeap) Push(x int) {
	h.arr = append(h.arr, x)  // Добавляем элемент в конец массива
	h.shiftUp(len(h.arr) - 1) // Просеиваем элемент вверх
}

// PopTop - удалить корневой элемент из кучи.
// time: O(log(n)), space: O(1)
func (h *MinHeap) PopTop() {
	if len(h.arr) == 0 { // Если куча пуста, то ничего не делаем
		return
	}
	h.arr[0], h.arr[len(h.arr)-1] = h.arr[len(h.arr)-1], h.arr[0] // Меняем местами корневой элемент и последний
	h.arr = h.arr[:len(h.arr)-1]                                  // Удаляем последний элемент
	h.shiftDown(0)                                                // Просеиваем корневой элемент вниз
}

// Top - вернуть корневой элемент из кучи.
// time: O(1), space: O(1)
func (h *MinHeap) Top() int {
	if len(h.arr) == 0 { // Если куча пуста, то паникуем
		panic("heap is empty")
	}
	return h.arr[0] // Возвращаем корневой элемент
}

// Empty - проверить, пуста ли куча.
// time: O(1), space: O(1)
func (h *MinHeap) Empty() bool {
	return len(h.arr) == 0 // Если размер массива равен нулю, то куча пуста
}

// shiftDown - просеивание вниз.
// time: O(log(n)), space: O(1)
func (h *MinHeap) shiftDown(i int) {
	leftChildIdx := i*2 + 1  // Индекс левого потомка
	rightChildIdx := i*2 + 2 // Индекс правого потомка

	if leftChildIdx >= len(h.arr) { // Если левого потомка нет, то просеивание вниз закончено
		return
	}

	nextIdx := leftChildIdx           // Индекс следующего элемента, который будем просеивать вниз
	nextMinVal := h.arr[leftChildIdx] // Значение следующего элемента, который будем просеивать вниз

	if rightChildIdx < len(h.arr) && h.arr[rightChildIdx] < nextMinVal { // Если правый потомок существует и меньше левого, то выбираем его
		nextIdx = rightChildIdx           // Индекс следующего элемента, который будем просеивать вниз
		nextMinVal = h.arr[rightChildIdx] // Значение следующего элемента, который будем просеивать вниз
	}

	currentVal := h.arr[i] // Значение текущего элемента

	if nextMinVal < currentVal { // Если значение следующего элемента меньше значения текущего, то меняем их местами и просеиваем следующий элемент вниз
		h.arr[i], h.arr[nextIdx] = h.arr[nextIdx], h.arr[i] // Меняем местами значения текущего и следующего элементов
		h.shiftDown(nextIdx)                                // Просеиваем следующий элемент вниз
	}
}

// shiftUp - просеивание вверх.
// time: O(log(n)), space: O(1)
func (h *MinHeap) shiftUp(i int) {
	if i == 0 { // Если текущий элемент - корень, то просеивание вверх закончено
		return
	}
	parentIdx := (i - 1) / 2         // Индекс родителя текущего элемента
	if h.arr[parentIdx] > h.arr[i] { // Если значение родителя больше значения текущего, то меняем их местами и просеиваем текущий элемент вверх
		h.arr[parentIdx], h.arr[i] = h.arr[i], h.arr[parentIdx] // Меняем местами значения родителя и текущего элементов
		h.shiftUp(parentIdx)                                    // Просеиваем текущий элемент вверх
	}
}

func main() {
	h := NewMinHeap()
	h.Push(3)
	h.Push(1)
	h.Push(4)
	h.Push(1)
	h.Push(5)

	fmt.Println("Top:", h.Top()) // 1
	h.PopTop()
	fmt.Println("Top after pop:", h.Top()) // 1
	h.PopTop()
	fmt.Println("Top after pop:", h.Top()) // 3
}
