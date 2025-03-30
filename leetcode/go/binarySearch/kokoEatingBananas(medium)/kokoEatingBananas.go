package main

import (
	"fmt"
	"math"
)

/* https://leetcode.com/problems/koko-eating-bananas/description/

Коко любит есть бананы. Есть n стопок бананов, в i-й стопке бананов находится [i]. Охранники ушли и вернутся через час.
Коко может решить, что ее скорость поедания бананов в час равна k. Каждый час она выбирает какую-то стопку бананов и
съедает k бананов из этой стопки. Если в куче меньше k бананов, она съедает их все и больше не будет есть бананов в
течение этого часа.
Коко любит есть медленно, но все равно хочет съесть все бананы до того, как вернутся охранники.
Найдите минимальное целое число k, при котором она сможет съесть все бананы за h часов.

Input: piles = [3,6,7,11], h = 8
Output: 4
*/

func main() {
	piles := []int{3, 6, 7, 11}
	fmt.Println(minEatingSpeed(piles, 8)) // 4
}

func minEatingSpeed(piles []int, h int) int {
	left := 1
	right := maxInSlice(piles)

	for left < right {
		mid := (left + right) / 2
		hourSpent := getHourSpent(mid, piles)

		if h < hourSpent {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return right
}

// Вспомогательная функция для вычисления необходимого времени
func getHourSpent(speed int, piles []int) int {
	totalHours := 0
	for _, pile := range piles {
		totalHours += int(math.Ceil(float64(pile) / float64(speed)))
	}
	return totalHours
}

// Вспомогательная функция для нахождения максимального значения в срезе
func maxInSlice(nums []int) int {
	maxNum := nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}
