package main

import "fmt"

/* If you can't sleep, just count sheep!!
https://www.codewars.com/kata/5b077ebdaf15be5c7f000077/train/go

Если вы не можете спать, просто посчитайте овец!!

Задача:
Для заданного неотрицательного целого числа, например 3, вернуть строку с бормотанием: "1 sheep...2 sheep...3 sheep...".
Ввод всегда будет допустимым, т. е. никаких отрицательных целых чисел.
*/

func main() {
	fmt.Println(countSheep(3)) // "1 sheep...2 sheep...3 sheep..."
	fmt.Println(countSheep(0)) // ""
}

// countSheep возвращает строку "1 sheep...2 sheep...3 sheep..." для заданного числа.
// time: O(n), space: O(n)
func countSheep(num int) string {
	result := ""

	for i := 0; i < num; i++ {
		result += fmt.Sprintf("%d sheep...", i+1)
	}

	return result
}
