package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* https://leetcode.com/problems/insert-delete-getrandom-o1/description/

Реализуйте класс RandomizedSet:

RandomizedSet() Инициализирует объект RandomizedSet.
bool Insert(int val) Вставляет элемент val в набор, если он отсутствует. Возвращает true, если элемент отсутствует, в
противном случае — false.
bool Remove(int val) Удаляет элемент val из набора, если он присутствует. Возвращает true, если элемент присутствовал, и
false в противном случае.
int getRandom() Возвращает случайный элемент из текущего набора элементов (гарантируется, что при вызове этого метода
существует хотя бы один элемент). Каждый элемент должен иметь одинаковую вероятность возврата.
Вы должны реализовать функции класса так, чтобы каждая функция работала со средней временной сложностью O(1).

Input
["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
[[], [1], [2], [2], [], [1], [2], []]
Output: [null, true, false, true, 2, true, false, 2]

Объяснение
RandomizedSet randomizedSet = новый RandomizedSet();
randomizedSet.insert(1); // Вставляет 1 в набор. Возвращает true, поскольку 1 была вставлена успешно.
randomizedSet.remove(2); // Возвращает false, поскольку числа 2 не существует в наборе.
randomizedSet.insert(2); // Вставляет 2 в набор, возвращает true. Набор теперь содержит [1,2].
randomizedSet.getRandom(); // getRandom() должен возвращать либо 1, либо 2 случайным образом.
randomizedSet.remove(1); // Удаляет 1 из набора, возвращает true. Набор теперь содержит [2].
randomizedSet.insert(2); // 2 уже было в наборе, поэтому возвращаем false.
randomizedSet.getRandom(); // Поскольку 2 — единственное число в наборе, getRandom() всегда будет возвращать 2.
*/

func main() {
	rs := Constructor()
	fmt.Println(rs.Insert(1))   // true
	fmt.Println(rs.Remove(2))   // false
	fmt.Println(rs.Insert(2))   // true
	fmt.Println(rs.GetRandom()) // 1 или 2
	fmt.Println(rs.Remove(1))   // true
	fmt.Println(rs.Insert(2))   // false
	fmt.Println(rs.GetRandom()) // 2
}

type RandomizedSet struct {
	container map[int]int // Хранит значение и его индекс в slice
	values    []int       // Хранит значения для получения случайного элемента
}

func Constructor() RandomizedSet {
	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел
	return RandomizedSet{
		container: make(map[int]int),
		values:    []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.container[val]; exists {
		return false // Элемент уже существует
	}
	this.container[val] = len(this.values) // Добавляем значение в map с индексом в slice
	this.values = append(this.values, val) // Добавляем значение в slice
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, exists := this.container[val]
	if !exists {
		return false // Элемент не существует
	}

	// Удаляем элемент из slice, заменяя его последним элементом
	lastElement := this.values[len(this.values)-1]
	this.values[index] = lastElement
	this.container[lastElement] = index

	// Удаляем последний элемент из slice и удаляем значение из map
	this.values = this.values[:len(this.values)-1]
	delete(this.container, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.values)) // Генерация случайного индекса
	return this.values[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
