package main

import (
	"bytes"
	"fmt"
	"math"
)

/* 166. Fraction to Recurring Decimal
https://leetcode.com/problems/fraction-to-recurring-decimal/description/

Даны два целых числа, представляющих числитель и знаменатель дроби, вернуть дробь в строковом формате.
Если дробная часть повторяется, заключите повторяющуюся часть в скобки.
Если возможны несколько ответов, верните любой из них.
Гарантируется, что длина строки ответа меньше 104 для всех заданных входных данных.

Input: numerator = 1, denominator = 2
Output: "0.5"
*/

func main() {
	fmt.Println(fractionToDecimal(1, 2)) // "0.5"
}

// fractionToDecimal Возвращает дробь в строковом формате.
// time: O(n), space: O(n)
func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 { // если числитель равен 0, то возвращаем 0
		return "0"
	}
	var res bytes.Buffer // создаем буфер для записи результата
	// если числитель и знаменатель имеют разные знаки, то добавляем знак "-" в начало результата
	if ((numerator < 0) && (denominator > 0)) || ((numerator > 0) && (denominator < 0)) {
		res.WriteString("-")
	}
	num := int(math.Abs(float64(numerator)))    // берем абсолютное значение числителя и знаменателя
	den := int(math.Abs(float64(denominator)))  // берем абсолютное значение числителя и знаменателя
	res.WriteString(fmt.Sprintf("%d", num/den)) // делим числитель на знаменатель и добавляем результат в результат
	num %= den                                  // берем остаток от деления числителя на знаменатель
	if num == 0 {                               // если остаток равен 0, то возвращаем результат
		return res.String()
	}
	res.WriteString(".")           // добавляем точку в результат
	remainder := make(map[int]int) // создаем мапу для записи остатка от деления
	remainder[num] = res.Len()     // записываем остаток от деления в мапу
	for num != 0 {                 // пока остаток от деления не равен 0
		num *= 10                                   // умножаем остаток от деления на 10
		res.WriteString(fmt.Sprintf("%d", num/den)) // делим остаток от деления на знаменатель и добавляем результат в результат
		num %= den                                  // берем остаток от деления от остатка от деления
		if index, ok := remainder[num]; ok {        // если остаток от деления уже встречался, то добавляем скобки в результат и возвращаем результат
			result := res.String()                             // создаем копию результата
			return result[:index] + "(" + result[index:] + ")" // возвращаем результат с добавленными скобками
		}
		remainder[num] = res.Len() // записываем остаток от деления в мапу
	}
	return res.String() // возвращаем результат
}
