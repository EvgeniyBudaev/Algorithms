package main

import (
	"fmt"
	"math/rand"
)

/* 384. Shuffle an Array
https://leetcode.com/problems/shuffle-an-array/description/

Имея целочисленный массив nums, разработайте алгоритм для случайного перемешивания массива. Все перестановки массива должны быть одинаково вероятны в результате перемешивания.

Реализуйте класс Solution:

Solution(int[] nums) Инициализирует объект целочисленным массивом nums.
int[] reset() Сбрасывает массив в исходную конфигурацию и возвращает его.
int[] shuffle() Возвращает случайное перемешивание массива.

Input
["Solution", "shuffle", "reset", "shuffle"]
[[[1, 2, 3]], [], [], []]
Output
[null, [3, 1, 2], [1, 2, 3], [1, 3, 2]]
*/

func main() {
	// Инициализация
	nums := []int{1, 2, 3}
	solution := Constructor(nums)
	// Тестирование
	fmt.Println("Original:", nums)
	fmt.Println("Shuffled:", solution.Shuffle())
	fmt.Println("Reset:", solution.Reset())
}

type Solution struct {
	orig []int // исходный массив
	copy []int // копия исходного массива
}

// Constructor - инициализирует объект целочисленным массивом nums.
// time: O(n), space: O(n)
func Constructor(nums []int) Solution {
	copyArr := make([]int, len(nums)) // копия исходного массива
	copy(copyArr, nums)               // копируем исходный массив
	return Solution{                  // возвращаем новый объект класса Solution с исходным массивом и копией исходного массива
		orig: nums,    // исходный массив
		copy: copyArr, // копия исходного массива
	}
}

// Reset - сбрасывает массив в исходную конфигурацию и возвращает его.
// time: O(n), space: O(1)
func (this *Solution) Reset() []int {
	return this.orig // возвращаем исходный массив
}

// Shuffle - возвращает случайное перемешивание массива.
// time: O(n), space: O(n)
func (this *Solution) Shuffle() []int {
	rand.Shuffle(len(this.orig), func(i, j int) { // перемешиваем копию исходного массива
		this.copy[i], this.copy[j] = this.copy[j], this.copy[i] // меняем элементы местами
	})

	return this.copy // возвращаем перемешанный массив
}
