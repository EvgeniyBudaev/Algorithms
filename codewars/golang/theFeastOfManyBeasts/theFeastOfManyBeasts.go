package main

import (
	"fmt"
)

/* The Feast of Many Beasts
https://www.codewars.com/kata/5aa736a455f906981800360d/train/go

Все животные пируют! Каждое животное приносит одно блюдо. Есть только одно правило: название блюда должно начинаться и
заканчиваться теми же буквами, что и название животного. Например, the great blue heron is bringing garlic naan and the
chickadee is bringing chocolate cake.

Напишите функцию feast, которая принимает имя животного и блюдо в качестве аргументов и возвращает true или false,
чтобы указать, разрешено ли животному принести блюдо на пир.

Предположим, что beast и dish всегда являются строками в нижнем регистре, и каждое состоит как минимум из двух букв.
beast и dish могут содержать дефисы и пробелы, но они не будут появляться в начале или конце строки. Они не будут
содержать цифр.
*/

func main() {
	fmt.Println(Feast("great blue heron", "garlic naan")) // true
	fmt.Println(Feast("brown bear", "bear claw"))         // false
}

// Feast проверяет, можно ли животному принести блюдо на пир.
// time: O(1), space: O(1)
func Feast(beast string, dish string) bool {
	return beast[0] == dish[0] && beast[len(beast)-1] == dish[len(dish)-1]
}
