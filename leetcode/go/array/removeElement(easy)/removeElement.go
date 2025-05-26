package main

import "fmt"

/* 27. Remove Element
https://leetcode.com/problems/remove-element/description/

Учитывая целочисленный массив nums и целочисленное значение, удалите все вхождения val в nums на месте. Порядок
элементов может быть изменен. Затем верните количество элементов в виде чисел, которые не равны val.
Учитывайте количество элементов в nums, которые не равны val be k. Чтобы вас приняли, вам необходимо сделать следующее:
Измените массив nums так, чтобы первые k элементов nums содержали элементы, не равные val. Остальные элементы nums не
важны, как и размер nums.
Вернуть К.

Input: nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,_,_]

Input: nums = [0,1,2,2,3,0,4,2], val = 2
Output: 5, nums = [0,1,4,0,3,_,_,_]
*/

func main() {
	nums := []int{3, 2, 2, 3}
	removeElement(nums, 3) // 2
}

// removeElement удаляет все вхождения val в nums.
// time: O(n), space: O(1)
func removeElement(nums []int, val int) int {
	count := 0 // Счетчик для отслеживания элементов, отличных от val

	for i := 0; i < len(nums); i++ { // Перебираем элементы массива
		if nums[i] != val { // Если текущий элемент не равен val, то переносим его в начало массива
			nums[count] = nums[i] // Перемещаем элемент в начало массива
			count++               // Увеличиваем счетчик
		}
	}

	// [2,2,2,3] -> [2,2,_,_]
	nums = nums[:count] // Обрезаем срез до фактической длины
	fmt.Println(nums)
	return count // Возвращаем фактическую длину
}
