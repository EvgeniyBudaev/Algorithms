package main

import (
	"fmt"
)

/* Opposites Attract
https://www.codewars.com/kata/555086d53eac039a2a000083/train/go

Тимми и Сара думают, что влюблены друг в друга, но в их районе они узнают об этом только после того, как сорвут по
цветку. Если у одного цветка чётное количество лепестков, а у другого — нечётное, это значит, что они влюблены.

Напишите функцию, которая будет принимать количество лепестков каждого цветка и возвращать значение true, если они
влюблены, и false, если нет.
*/

func main() {
	fmt.Println(LoveFunc(1, 4)) // true
	fmt.Println(LoveFunc(2, 2)) // false
}

// LoveFunc принимает количество лепестков каждого цветка и возвращает значение true, если они влюблены.
// time: O(1), space: O(1)
func LoveFunc(flower1, flower2 int) bool {
	return flower1%2 != flower2%2
}
