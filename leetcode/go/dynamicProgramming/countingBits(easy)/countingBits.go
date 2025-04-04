package main

import "fmt"

/* https://leetcode.com/problems/counting-bits/description/

Учитывая целое число n, верните массив ans длины n + 1 такой, что для каждого i (0 <= i <= n) ans[i] — это количество
единиц в двоичном представлении i.

Input: n = 2
Output: [0,1,1]
Explanation:
0 --> 0
1 --> 1
2 --> 10
*/

func main() {
	fmt.Println(countBits(2)) // [0 1 1]
}

// countBits возвращает массив, где каждый элемент содержит количество 1 в двоичном представлении числа
func countBits(n int) []int {
	// Инициализируем массив результатов с нулями
	dp := make([]int, n+1)

	// offset - текущая степень двойки
	offset := 1

	for i := 1; i <= n; i++ {
		// Если достигли новой степени двойки, обновляем offset
		if offset*2 == i {
			offset = i
		}
		// Количество единиц = 1 (для нового старшего бита) + количество в оставшейся части
		dp[i] = 1 + dp[i-offset]
	}

	return dp
}
