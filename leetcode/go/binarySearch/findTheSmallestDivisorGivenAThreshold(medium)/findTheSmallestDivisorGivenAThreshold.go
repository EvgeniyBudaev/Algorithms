package main

import (
	"fmt"
	"math"
)

/* 1283. Find the Smallest Divisor Given a Threshold
https://leetcode.com/problems/find-the-smallest-divisor-given-a-threshold/description/

Учитывая массив целых чисел nums и целочисленный threshold, мы выберем положительный целочисленный divisor, разделим на
него весь массив и просуммируем результат деления. Найдите наименьший divisor, чтобы упомянутый выше результат был
меньше или равен threshold.
Каждый результат деления округляется до ближайшего целого числа, большего или равного этому элементу.
(Например: 7/3 = 3 и 10/2 = 5).
Тестовые случаи генерируются для того, чтобы был ответ.

Input: nums = [1,2,5,9], threshold = 6
Output: 5
Пояснение: Мы можем получить сумму 17 (1+2+5+9), если делитель равен 1.
Если делитель равен 4, мы можем получить сумму 7 (1+1+2+3), а если делитель равен 5, сумма будет 5 (1+1+1+2).

Input: nums = [44,22,33,11,1], threshold = 5
Output: 44
*/

func main() {
	nums := []int{1, 2, 5, 9}
	fmt.Println(smallestDivisor(nums, 6)) // 5
}

// smallestDivisor находит наименьший делитель, который при делении всех элементов массива на него дает сумму, не превышающую порога.
func smallestDivisor(nums []int, threshold int) int {
	left := 1
	right := maxInSlice(nums) // Находим максимальное значение в массиве

	for left <= right {
		mid := (left + right) / 2
		// Если сумма делителей не превышает порога
		if sumDivisions(nums, mid) <= threshold {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

// Вспомогательная функция для вычисления суммы округленных вверх частных
func sumDivisions(nums []int, divisor int) int {
	total := 0 // Инициализируем переменную для хранения суммы частных
	for _, num := range nums {
		// Вычисляем частное и округляем вверх до ближайшего целого числа
		total += int(math.Ceil(float64(num) / float64(divisor)))
	}
	return total
}

// Вспомогательная функция для нахождения максимального значения в срезе
func maxInSlice(nums []int) int {
	maxNum := nums[0] // Инициализируем переменную для хранения максимального значения с первым элементом среза
	for _, num := range nums {
		// Если текущий элемент больше максимального значения, обновляем максимальное значение
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}
