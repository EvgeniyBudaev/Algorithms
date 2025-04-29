package main

import (
	"fmt"
)

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(nums, 9)) // 4
}

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) / 2
		guess := nums[mid]
		if guess == target {
			return mid
		}
		if guess > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
