package main

import "fmt"

/* 303. Range Sum Query - Immutable
https://leetcode.com/problems/range-sum-query-immutable/description/

Учитывая целочисленный массив nums, обработайте несколько запросов следующего типа:

Вычислить сумму элементов чисел между индексами слева и справа включительно, где слева <= справа.
Реализуйте класс NumArray:

NumArray(int[] nums) Инициализирует объект целочисленным массивом nums.
int sumRange(int left, int right) Возвращает сумму элементов nums между индексами слева и справа включительно
(т. е. nums[left] + nums[left + 1] + ... + nums[right]).

Input
["NumArray", "sumRange", "sumRange", "sumRange"]
[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
Output
[null, 1, -1, -3]

Объяснение
NumArray numArray = новый NumArray([-2, 0, 3, -5, 2, -1]);
numArray.sumRange(0, 2); // возвращаем (-2) + 0 + 3 = 1
numArray.sumRange(2, 5); // возвращаем 3 + (-5) + 2 + (-1) = -1
numArray.sumRange(0, 5); // возвращаем (-2) + 0 + 3 + (-5) + 2 + (-1) = -3
*/

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums) // Создаем объект NumArray с массивом чисел

	fmt.Println(obj.SumRange(0, 2)) // 1 (-2 + 0 + 3)
}

// NumArray - структура для хранения массива чисел и эффективного вычисления суммы элементов в диапазоне.
type NumArray struct {
	nums      []int // Исходный массив чисел
	prefixSum []int // Массив префиксных сумм для оптимизации
}

// Constructor создает новый объект NumArray и инициализирует массив префиксных сумм.
// time: O(n) - линейное время, так как мы проходим по всем элементам массива чисел
// space: O(n) - линейное пространство, так как мы храним массив префиксных сумм
func Constructor(nums []int) NumArray {
	prefix := make([]int, len(nums)+1) // Создаем массив префиксных сумм
	prefix[0] = 0                      // Инициализируем первое значение в 0

	// Заполняем массив префиксных сумм
	for i := 0; i < len(nums); i++ {
		// Суммируем текущее число с предыдущей суммой префикса
		prefix[i+1] = prefix[i] + nums[i]
	}
	// prefix [0, -2, -2, 1, -4, -2, -3]

	return NumArray{
		nums:      nums,
		prefixSum: prefix,
	}
}

// SumRange вычисляет сумму элементов массива между индексами left и right включительно.
// time: O(1) - постоянное время, так как мы просто обращаемся к массиву префиксных сумм
// space: O(n) - линейное пространство, так как мы храним массив префиксных сумм
func (this *NumArray) SumRange(left int, right int) int {
	// Используем массив префиксных сумм для быстрого вычисления
	return this.prefixSum[right+1] - this.prefixSum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
