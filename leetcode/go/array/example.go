package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(nums)) // 3
}

func findMaxConsecutiveOnes(nums []int) int {
	left, zeroCount, ans := 0, 0, 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}
		for zeroCount > 0 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
