package main

import (
	"fmt"
	"sort"
)

/* https://leetcode.com/problems/squares-of-a-sorted-array/description/

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
	r := sortedSquares([]int{-4, -1, 0, 3, 10})
	fmt.Println(r)
}

func sortedSquares(nums []int) []int {
	squares := make([]int, len(nums))

	for i, num := range nums {
		squares[i] = num * num
	}

	sort.Ints(squares)

	return squares
}
