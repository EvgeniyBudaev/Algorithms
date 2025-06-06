package main

import "fmt"

/* 268. Missing Number
https://leetcode.com/problems/missing-number/description/

Учитывая массив nums, содержащий n различных чисел в диапазоне [0, n], верните единственное число в диапазоне,
которое отсутствует в массиве.

Input: nums = [3,0,1]
Output: 2
Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the
range since it does not appear in nums.

Input: nums = [0,1]
Output: 2
Explanation: n = 2 since there are 2 numbers, so all numbers are in the range [0,2]. 2 is the missing number in the
range since it does not appear in nums.

Input: nums = [9,6,4,2,3,5,7,0,1]
Output: 8
Explanation: n = 9 since there are 9 numbers, so all numbers are in the range [0,9]. 8 is the missing number in the
range since it does not appear in nums.
*/

func main() {
	nums := []int{3, 0, 1}
	fmt.Println(missingNumber(nums)) // 2
}

// missingNumber находит пропущенное число в последовательности [0, n]
// time: O(n), space: O(1)
func missingNumber(nums []int) int {
	n := len(nums)                 // Длина массива
	expectedSum := n * (n + 1) / 2 // Сумма чисел от 0 до n
	actualSum := 0                 // Сумма чисел в массиве
	// Суммируем числа в массиве
	for _, num := range nums {
		actualSum += num // Суммируем число с текущей суммой
	}
	// Вычитаем сумму чисел в массиве из ожидаемой суммы чисел от 0 до n
	return expectedSum - actualSum
}

// missingNumber находит пропущенное число в последовательности [0, n]
// time complexity: O(n) - проходим по всем элементам массива
// space complexity: O(n) - использует дополнительное пространство для хранения множества чисел
//func missingNumber(nums []int) int {
//	// Создаем множество (map) для хранения чисел из массива
//	numSet := make(map[int]bool)
//
//	// Заполняем множество числами из входного массива
//	for _, num := range nums {
//		numSet[num] = true
//	}
//
//	// Проверяем каждое число от 0 до длины массива
//	for i := 0; i < len(nums); i++ {
//		if !numSet[i] { // Если числа нет в множестве
//			return i // Это пропущенное число
//		}
//	}
//
//	// Если все числа от 0 до n-1 присутствуют, значит пропущено число n
//	return len(nums)
//}
