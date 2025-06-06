package main

import "fmt"

/* 239. Sliding Window Maximum
https://leetcode.com/problems/sliding-window-maximum/description/
https://leetcode.com/problems/minimum-window-substring/solutions/3061723/js-fastest-easy-commented/

Вам дан массив целых чисел nums, есть скользящее окно размера k, которое перемещается из самого левого края массива в
самый правый. В окне вы можете увидеть только k чисел. Каждый раз скользящее окно перемещается вправо на одну позицию.
Верните максимальное число в скользящем окне.

Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]
Explanation:
Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

Input: nums = [1], k = 1
Output: [1]
*/

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(nums, 3)) // [3,3,5,5,6,7]
}

// maxSlidingWindow - функция для нахождения максимального значения в каждом окне.
// time: O(n), space: O(k), где n - длина входного массива, k - размер окна.
func maxSlidingWindow(nums []int, k int) []int {
	var result []int // Инициализируем массив для хранения максимальных значений
	var deque []int  // Инициализируем дек (двустороннюю очередь) для хранения индексов элементов в текущем окне

	// Перебираем входной массив
	for i := 0; i < len(nums); i++ {
		// Удаляем из дека индексы элементов, которые меньше или равны текущему элементу
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}
		// Удаляем индексы элементов из дека, находящихся за пределами текущего окна
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}
		deque = append(deque, i) // Добавляем текущий индекс в дек
		// Если текущий индекс достиг размера окна или больше, добавляем максимальный элемент из дека в массив результатов
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result // Возвращаем массив, содержащий максимальный элемент в каждом скользящем окне
}
