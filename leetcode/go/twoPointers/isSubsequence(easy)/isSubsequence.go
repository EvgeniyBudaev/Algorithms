package main

import "fmt"

/*
Учитывая две строки s и t, верните true, если s является подпоследовательностью t, или false в противном случае.

Подпоследовательность строки — это последовательность символов, которую можно получить путем удаления некоторых
(или ни одного) символов из исходной строки, сохраняя при этом относительный порядок оставшихся символов.
Например, «ace» является подпоследовательностью «abcde», а «aec» — нет.
*/

func main() {
	fmt.Println(isSubsequence("ace", "abcde")) // true
	fmt.Println(isSubsequence("aec", "abcde")) // false
}

// isSubsequence проверяет, является ли строка s подпоследовательностью строки t.
// time: O(n), space: O(1)
func isSubsequence(s, t string) bool {
	left, right := 0, 0

	for left < len(s) && right < len(t) { // пока не достигнут конец одной из строк
		if s[left] == t[right] { // если символы совпадают
			left++
		}
		right++
	}

	return left == len(s) // если все символы из s были найдены в t, то возвращаем true
}
