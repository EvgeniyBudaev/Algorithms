package main

import "fmt"

/* 739. Daily Temperatures
https://leetcode.com/problems/daily-temperatures/description/

Учитывая, что массив целых чисел представляет собой ежедневную температуру, верните ответ массива, такой,
что ответ [i] — это количество дней, которое вам нужно подождать после i-го дня, чтобы получить более высокую
температуру. Если нет будущего дня, для которого это возможно, вместо этого сохраните ответ [i] == 0.

Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]

Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]

Input: temperatures = [30,60,90]
Output: [1,1,0]
*/

func main() {
	temps := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temps)) // [1,1,4,2,1,1,0,0]
}

// dailyTemperatures — функция, которая возвращает количество дней, которое вам нужно подождать после i-го дня,
// чтобы получить более высокую температуру.
func dailyTemperatures(temperatures []int) []int {
	var stack []int // Стек для хранения индексов дней с пониженной температурой
	result := make([]int, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]   // Запоминаем индекс последнего элемента в стеке
			stack = stack[:len(stack)-1] // Удаляем последний элемент
			result[idx] = i - idx        // Записываем количество дней
		}
		stack = append(stack, i)
	}

	return result
}
