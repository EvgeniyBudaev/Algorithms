package main

import "fmt"

func main() {
	sweetness := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(maximizeSweetness(sweetness, 5)) // 6
}

func maximizeSweetness(sweetness []int, k int) int {
	left := 0
	right := 0
	for _, num := range sweetness {
		right += num
	}

	canSplit := func(minimumSweetness int) bool {
		currentSweetness := 0
		partsCount := 0
		for _, value := range sweetness {
			currentSweetness += value
			if currentSweetness >= minimumSweetness {
				currentSweetness = 0
				partsCount++
			}
		}
		return partsCount > k
	}

	for left < right {
		mid := (left + right + 1) / 2
		if canSplit(mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return left
}
