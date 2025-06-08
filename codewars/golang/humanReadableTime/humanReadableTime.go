package main

import "fmt"

/* Human Readable Time

Напишите функцию, которая принимает в качестве входных данных неотрицательное целое число (секунды) и возвращает время в
удобном для восприятия формате (HH:MM:SS)

HH = hours, padded to 2 digits, range: 00 - 99
MM = minutes, padded to 2 digits, range: 00 - 59
SS = seconds, padded to 2 digits, range: 00 - 59
Максимальное время никогда не превышает 359999 (99:59:59)
*/

func main() {
	fmt.Println(HumanReadableTime(359999)) // 99:59:59
}

// HumanReadableTime возвращает время в удобном для восприятия формате (HH:MM:SS)
// time: O(1), space: O(1)
func HumanReadableTime(seconds int) string {
	hours := seconds / 3600            // Часы // 99
	remainingSeconds := seconds % 3600 // Оставшиеся секунды // 3599
	minutes := remainingSeconds / 60   // Минуты // 59
	secs := remainingSeconds % 60      // Секунды // 59

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, secs)
}
