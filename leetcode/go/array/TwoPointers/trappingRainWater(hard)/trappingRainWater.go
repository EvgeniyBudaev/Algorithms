package main

import "fmt"

/* https://leetcode.com/problems/trapping-rain-water/description/

Учитывая n неотрицательных целых чисел, представляющих карту высот, где ширина каждой полосы равна 1, вычислите,
сколько воды она может удержать после дождя.

Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Пояснение: Приведенная выше карта высот (черное сечение) представлена массивом [0,1,0,2,1,0,1,3,2,1,2,1].
В данном случае задерживается 6 единиц дождевой воды (синяя секция).
*/

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height)) // 6
}

func trap(height []int) int {
	left := 0
	leftMaxValue := height[left]
	right := len(height) - 1
	rightMaxValue := height[right]
	sum := 0

	for left < right {
		if leftMaxValue <= rightMaxValue {
			sum += leftMaxValue - height[left]
			left++
			leftMaxValue = max(leftMaxValue, height[left])
		} else {
			sum += rightMaxValue - height[right]
			right--
			rightMaxValue = max(rightMaxValue, height[right])
		}
	}

	return sum
}
