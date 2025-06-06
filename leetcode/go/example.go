package main

import (
	"fmt"
)

func main() {
	arr := []int{0, 3, 2, 1}
	fmt.Println(validMountainArray(arr)) // true
}

func validMountainArray(arr []int) bool {
	if len(arr) < 2 {
		return false
	}
	left, right := 0, len(arr)-1
	for i := 0; i < len(arr); i++ {
		if arr[left] < arr[left+1] {
			left++
		}
		if arr[right] < arr[right-1] {
			right--
		}
		if left == 0 || right == len(arr)-1 {
			return false
		}
		if left == right {
			return true
		}
	}
	return false
}
