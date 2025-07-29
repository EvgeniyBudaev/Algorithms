package main

import (
	"fmt"
	"math"
)

/* Maximum Length Difference
https://www.codewars.com/kata/5663f5305102699bad000056/train/go

Вам даны два массива строк a1 и a2. Каждая строка состоит из букв от a до z. Пусть x — любая строка из первого массива,
а y — любая строка из второго.

Найти max(abs(length(x) − length(y)))

Если a1 и/или a2 пустые, вернуть -1 во всех языках, кроме Haskell (F#), где вы вернёте Nothing (None).

Пример:
a1 = ["hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"]
a2 = ["cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"]
mxdiflg(a1, a2) --> 13
*/

func main() {
	a1 := []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
	a2 := []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}
	fmt.Println(MxDifLg(a1, a2)) // 13
}

// MxDifLg возвращает максимальное значение abs(length(x) − length(y)).
// time: O(n^2), space: O(1), где n - длина
func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}

	minA1, maxA1 := findMinMaxLengths(a1)
	minA2, maxA2 := findMinMaxLengths(a2)

	diff1 := math.Abs(float64(maxA1 - minA2))
	diff2 := math.Abs(float64(minA1 - maxA2))

	return int(math.Max(diff1, diff2))
}

// findMinMaxLengths возвращает минимальную и максимальную длину строки в массиве.
// time: O(n), space: O(1), где n - длина
func findMinMaxLengths(arr []string) (int, int) {
	if len(arr) == 0 {
		return 0, 0
	}

	minLen := len(arr[0])
	maxLen := len(arr[0])

	for _, s := range arr {
		length := len(s)
		if length < minLen {
			minLen = length
		}
		if length > maxLen {
			maxLen = length
		}
	}

	return minLen, maxLen
}
