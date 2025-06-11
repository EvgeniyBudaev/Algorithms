package main

import (
	"fmt"
)

/* Consecutive strings
https://www.codewars.com/kata/56a5d994ac971f1ac500003e/train/go

Вам дан массив(список) strarr строк и целое число k. Ваша задача — вернуть первую самую длинную строку,
состоящую из k последовательных строк, взятых в массиве.

strarr = ["tree", "foling", "trashy", "blue", "abcdef", "uvwxyz"], k = 2

Объединяем последовательные строки strarr по 2, получаем:

treefoling   (length 10)  concatenation of strarr[0] and strarr[1]
folingtrashy ("      12)  concatenation of strarr[1] and strarr[2]
trashyblue   ("      10)  concatenation of strarr[2] and strarr[3]
blueabcdef   ("      10)  concatenation of strarr[3] and strarr[4]
abcdefuvwxyz ("      12)  concatenation of strarr[4] and strarr[5]

Две строки самые длинные: «folingtrashy» и «abcdefuvwxyz».
Первая пришла «folingtrashy», поэтому
longest_consec(strarr, 2) должна вернуть «folingtrashy».

Таким же образом:
longest_consec(["zone", "abigail", "theta", "form", "libe", "zas", "theta", "abigail"], 2) --> "abigailtheta"
*/

func main() {
	strarr := []string{"tree", "foling", "trashy", "blue", "abcdef", "uvwxyz"}
	fmt.Println(LongestConsec(strarr, 2)) // "folingtrashy"
}

// LongestConsec возвращает самую длинную строку из k последовательных строк в массиве strarr.
func LongestConsec(strarr []string, k int) string {
	n := len(strarr) // Длина массива

	if n == 0 || k > n || k <= 0 {
		return ""
	}

	longest := "" // Самая длинная строка
	maxLen := 0   // Длина самой длинной строки

	for i := 0; i < n-k; i++ { // Для каждого i-го элемента
		current := ""            // Текущая строка
		for j := 0; j < k; j++ { // Для каждого k-го элемента
			current += strarr[i+j] // Добавляем в текущую строку
		}
		if len(current) > maxLen { // Если длина текущей строки больше длины самой длинной строки
			maxLen = len(current) // Текущая длина становится самой длинной
			longest = current     // Текущая строка становится самой длинной
		}
	}

	return longest // Возвращаем самую длинную строку
}
