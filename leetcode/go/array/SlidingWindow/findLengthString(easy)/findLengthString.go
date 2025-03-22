package main

import "fmt"

/*
Вам дана двоичная строка s (строка, содержащая только «0» и «1»). Вы можете выбрать до одного «0» и изменить его на «1».
Какова длина самой длинной подстроки, содержащей только «1»?

Например, если задано s = «1101100111», ответ будет равен 5. Если вы выполните переворот по индексу 2,
строка станет 1111100111.

Input: nums = "1101100111"
Output: 5
*/

func main() {
	fmt.Println(findLengthString("1101100111"))
}

func findLengthString(s string) int {
	left, zeroCount, ans := 0, 0, 0

	for right := 0; right < len(s); right++ {
		if s[right] == '0' {
			zeroCount++
		}

		// Если количество нулей больше 1, сдвигаем left
		for zeroCount > 1 {
			if s[left] == '0' {
				zeroCount--
			}
			left++
		}

		// Обновляем максимальную длину
		currentLength := right - left + 1
		if currentLength > ans {
			ans = currentLength
		}
	}

	return ans
}
