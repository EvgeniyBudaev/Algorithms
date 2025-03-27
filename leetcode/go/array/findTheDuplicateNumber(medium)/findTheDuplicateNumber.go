package main

import "fmt"

/* https://leetcode.com/problems/find-the-duplicate-number/description/

Дан массив целых чисел nums, содержащий n + 1 целых чисел, где каждое целое число находится в диапазоне [1, n]
включительно.
В nums есть только одно повторяющееся число, верните это повторяющееся число.
Вы должны решить проблему, не изменяя числа массива и используя только постоянное дополнительное пространство.

Input: nums = [1,3,4,2,2]
Output: 2

Input: nums = [3,1,3,4,2]
Output: 3
*/

func main() {
	arr := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(arr)) // 2
}

func findDuplicate(nums []int) int {
	visited := make([]bool, len(nums))

	for _, num := range nums {
		if !visited[num] {
			visited[num] = true
		} else {
			return num
		}
	}

	return -1
}
