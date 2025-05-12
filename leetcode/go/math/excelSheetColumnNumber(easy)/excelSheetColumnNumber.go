package main

import (
	"fmt"
)

/* 171. Excel Sheet Column Number
https://leetcode.com/problems/excel-sheet-column-number/description/

Если задана строка columnTitle, представляющая заголовок столбца, как он отображается на листе Excel,
вернуть соответствующий ему номер столбца.

Input: columnTitle = "A"
Output: 1
*/

func main() {
	fmt.Println(titleToNumber("A")) // 1
}

// titleToNumber - принимает строку и возвращает ее номер в excel.
// time: O(n), space: O(1)
func titleToNumber(columnTitle string) int {
	final := 0 // Инициализируем переменную для хранения итогового значения

	// Перебираем все символы
	for _, v := range columnTitle {
		// Берем остаток текущего символа и добавляем 1, потому что номер столбца в Excel начинается с 1
		// 65 это значение 'A' в таблице ASCII
		num := int(v%65) + 1
		// Умножаем на 26 и прибавляем текущий символ
		// 26 это количество букв в английском алфавите
		final = (final * 26) + num
	}

	return final // Возвращаем итоговое значение
}
