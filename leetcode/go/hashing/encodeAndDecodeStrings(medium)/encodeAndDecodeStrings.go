package main

import (
	"fmt"
	"strconv"
)

/* 659 · Encode and Decode Strings

Разработайте алгоритм для кодирования списка строк в одну строку. Закодированная строка затем декодируется обратно в
исходный список строк.
Пожалуйста, реализуйте кодирование и декодирование

Input: ["neet","code","love","you"]
Output:["neet","code","love","you"]

Input: ["we","say",":","yes"]
Output: ["we","say",":","yes"]
*/

func main() {
	encoded := encode([]string{"Hello", "World"})
	fmt.Println("encode:", encoded) // "00000101Hello00000101World"

	decoded := decode(encoded)
	fmt.Println("decode:", decoded) // ["Hello", "World"]
}

// Кодирует список строк в одну строку.
// Длина каждой строки сохраняется как префикс фиксированного размера перед фактическим содержимым строки.
func encode(strs []string) string {
	var sb []byte
	for _, str := range strs { /* Time O(N) */
		code := getCode(str) /* Time O(1) */ // для "Hello" -> 00000101
		sb = append(sb, code...)
		sb = append(sb, str...)
	}
	return string(sb) /* Time O(N) | Ignore Auxiliary Space O(N) */
}

// Gets the 8-byte binary representation of string length
func getCode(str string) string {
	length := len(str)
	// Преобразование длины в двоичную систему: .toString(2) - метод преобразует число (в данном случае длину строки) в
	// строку, представляющую это число в двоичной системе счисления
	binary := strconv.FormatInt(int64(length), 2)
	// Добавление ведущих нулей: .padStart(8, '0') - метод добавляет символы ('0' в данном случае) в начало строки до
	// тех пор, пока ее общая длина не достигнет указанного значения (8). Это делается для обеспечения фиксированной
	// ширины результата, что особенно полезно при работе с битовыми данными или форматировании данных.
	return fmt.Sprintf("%08s", binary) // Pad with leading zeros to make 8 chars
}

// Декодирует одну строку в список строк, считывая префикс длины фиксированного размера
// и затем считываем соответствующее количество символов.
func decode(str string) []string {
	var output []string
	left := 0
	for left < len(str) {
		right := left + 8
		if right > len(str) {
			break
		}

		countString := str[left:right] // 00000101
		length, err := strconv.ParseInt(countString, 2, 64)
		if err != nil {
			panic("Invalid encoded string format")
		}

		rightEnd := right + int(length)
		if rightEnd > len(str) {
			panic("Invalid encoded string length")
		}

		decoding := str[right:rightEnd]
		output = append(output, decoding)

		left = rightEnd
	}
	return output
}
