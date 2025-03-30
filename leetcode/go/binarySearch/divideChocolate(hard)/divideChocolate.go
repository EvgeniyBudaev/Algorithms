package main

/* 1231. Divide Chocolate

У вас есть одна плитка шоколада, состоящая из нескольких кусочков. Каждый кусок имеет свою сладость, определяемую
сладостью массива.
Вы хотите поделиться шоколадом со своими K друзьями, поэтому начинаете разрезать плитку шоколада на K+1 кусочков,
используя K разрезов, каждый кусок состоит из нескольких последовательных кусков.
Проявив щедрость, вы съедите кусок с минимальной общей сладостью, а остальные кусочки раздадите друзьям.
Найдите максимальную общую сладость кусочка, которую вы можете получить, оптимально разрезав плитку шоколада.

Input: sweetness = [1,2,3,4,5,6,7,8,9], K = 5
Output: 6
Пояснение: Вы можете разделить шоколад на [1,2,3], [4,5], [6], [7], [8], [9]

Input: sweetness = [5,6,7,8,9,1,2,3,4], K = 8
Output: 1
Пояснение: Разрезать брусок на 9 частей можно только одним способом.

Input: sweetness = [1,2,2,1,2,2,1,2,2], K = 2
Output: 5
Пояснение: Вы можете разделить шоколад на [1,2,2], [1,2,2], [1,2,2]
*/

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
