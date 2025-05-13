package main

import (
	"fmt"
)

/* 977. Squares of a Sorted Array
https://leetcode.com/problems/squares-of-a-sorted-array/description/

Учитывая целочисленный массив nums, отсортированный в неубывающем порядке, верните массив квадратов каждого числа,
отсортированного в неубывающем порядке.

Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].

Input: nums = [-7,-3,2,3,11]
Output: [4,9,9,49,121]

Two pointers
Time complexity: O(n)
Space complexity: O(n)
*/

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10})) // [0,1,9,16,100]
}

// sortedSquares - возвращает массив квадратов каждого числа, отсортированного в неубывающем порядке.
// time: O(n), space: O(n)
//func sortedSquares(nums []int) []int {
//	for i, num := range nums {
//		nums[i] = num * num
//	}
//
//	sort.Ints(nums)
//	return nums
//}

// sortedSquares - возвращает массив квадратов каждого числа, отсортированного в неубывающем порядке.
// time: O(n), space: O(n)
func sortedSquares(nums []int) []int {
	n := len(nums)           // Длина массива
	result := make([]int, n) // Массив квадратов
	p1, p2 := 0, n-1         // Указатели на начало и конец массива

	for i := n - 1; p1 <= p2; i-- { // Проход по массиву
		if abs(nums[p1]) > abs(nums[p2]) { // Если модуль первого элемента больше модуля второго
			result[i] = nums[p1] * nums[p1] // Записываем квадрат первого элемента
			p1++
		} else {
			result[i] = nums[p2] * nums[p2]
			p2--
		}
	}

	return result
}

// abs - возвращает абсолютное значение числа.
// time: O(1), space: O(1)
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
