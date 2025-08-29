package main

import (
	"fmt"
)

/* Volume of a Cuboid
https://www.codewars.com/kata/58261acb22be6e2ed800003a/train/go

Бобу нужен быстрый способ вычисления объема прямоугольного параллелепипеда по трем значениям: длине, ширине и высоте
параллелепипеда.
*/

func main() {
	fmt.Println(GetVolumeOfCuboid(1.0, 2.0, 2.0)) // 4.0
}

// GetVolumeOfCuboid возвращает объём параллелепипеда по трем значениям: длине, ширине и высоте.
// time: O(1), space: O(1)
func GetVolumeOfCuboid(length, width, height float64) float64 {
	return length * width * height
}
