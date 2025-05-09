package main

import "fmt"

/* 42. Trapping Rain Water
https://leetcode.com/problems/trapping-rain-water/description/

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

// trap вычисляет количество воды, которое может быть залита в массиве высот после дождя.
func trap(height []int) int {
	left := 0
	leftMaxValue := height[left] // максимальная высота слева
	right := len(height) - 1
	rightMaxValue := height[right] // максимальная высота справа
	sum := 0                       // общее количество воды, которое может быть залита

	for left < right {
		// если максимальная высота слева меньше максимальной высоты справа, двигаем левую границу вправо
		if leftMaxValue <= rightMaxValue {
			sum += leftMaxValue - height[left] // вычисляем количество воды, которое может быть залита в текущей ячейке
			left++
			leftMaxValue = max(leftMaxValue, height[left]) // обновляем максимальную высоту слева
		} else {
			sum += rightMaxValue - height[right] // вычисляем количество воды, которое может быть залита в текущей ячейке
			right--
			rightMaxValue = max(rightMaxValue, height[right]) // обновляем максимальную высоту справа
		}
	}

	return sum
}
