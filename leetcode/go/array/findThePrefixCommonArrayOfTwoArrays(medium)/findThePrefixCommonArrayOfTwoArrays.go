package main

import "fmt"

/* https://leetcode.com/problems/find-the-prefix-common-array-of-two-arrays/description/

Вам даны две целочисленные перестановки A и B с нулевым индексом длины n.
Общий префиксный массив A и B — это массив C, такой, что C[i] равен количеству чисел, которые присутствуют в индексе i
или перед ним как в A, так и в B.
Верните общий массив префиксов A и B.
Последовательность из n целых чисел называется перестановкой, если она содержит все целые числа от 1 до n ровно один раз.

Input: A = [1,3,2,4], B = [3,1,2,4]
Output: [0,2,3,4]
Объяснение: При i = 0: нет общих чисел, поэтому C[0] = 0.
При i = 1: 1 и 3 являются общими в A и B, поэтому C[1] = 2.
При i = 2: 1, 2 и 3 являются общими в A и B, поэтому C[2] = 3.
При i = 3: 1, 2, 3 и 4 являются общими в A и B, поэтому C[3] = 4.
*/

func main() {
	A := []int{1, 3, 2, 4}
	B := []int{3, 1, 2, 4}
	fmt.Println(findThePrefixCommonArray(A, B)) // [0, 2, 3, 4]
}

func findThePrefixCommonArray(A []int, B []int) []int {
	res := []int{}
	set := make(map[int]bool)
	count := 0

	for i := 0; i < len(A); i++ {
		if set[A[i]] {
			count++
		}
		if set[B[i]] {
			count++
		}
		if A[i] == B[i] {
			count++
		}
		set[A[i]] = true
		set[B[i]] = true
		res = append(res, count)
	}

	return res
}
