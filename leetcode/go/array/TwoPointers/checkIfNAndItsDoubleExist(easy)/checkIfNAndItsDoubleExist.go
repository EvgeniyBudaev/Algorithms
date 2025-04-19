package main

import "fmt"

/* 1346. Check If N and Its Double Exist
https://leetcode.com/problems/check-if-n-and-its-double-exist/description/

Учитывая массив целых чисел, проверьте, существуют ли два индекса i и j такие, что:
i != j
0 <= i, j < arr.length
arr[i] == 2 * arr[j]

Input: arr = [10,2,5,3]
Output: true
Explanation: For i = 0 and j = 2, arr[i] == 10 == 2 * 5 == 2 * arr[j]

Input: arr = [3,1,7,11]
Output: false
Пояснение: Не существует i и j, удовлетворяющих этим условиям.
*/

func main() {
	fmt.Println(checkIfExist([]int{10, 2, 5, 3})) // true
}

// checkIfExist проверяет, существует ли в массиве arr два различных индекса i и j,
// таких что один элемент равен удвоенному другому (т.е. arr[i] == 2 * arr[j] или arr[j] == 2 * arr[i]).
func checkIfExist(arr []int) bool {
	left, right := 0, 1

	for left < len(arr)-1 {
		if arr[left] == arr[right]*2 || arr[right] == arr[left]*2 {
			return true
		} else if right == len(arr)-1 {
			left++
			right = left + 1
		} else {
			right++
		}
	}

	return false
}
