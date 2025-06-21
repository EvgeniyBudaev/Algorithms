package main

import (
	"fmt"
)

/* L1: Set Alarm
https://www.codewars.com/kata/568dcc3c7f12767a62000038/train/go

Напишите функцию с именем setAlarm/set_alarm/set-alarm/setalarm (в зависимости от языка), которая получает два параметра.
Первый параметр, busy, имеет значение true, когда вы работаете, а второй параметр, vacation, имеет значение true, когда
вы в отпуске.

Функция должна возвращать значение true, если вы работаете и не в отпуске (потому что это обстоятельства, при которых
вам нужно установить будильник). В противном случае она должна возвращать значение false. Примеры:

employed | vacation
true     | true     => false
true     | false    => true
false    | true     => false
false    | false    => false
*/

func main() {
	fmt.Println(SetAlarm(true, false)) // true
	fmt.Println(SetAlarm(false, true)) // false
}

// SetAlarm возвращает true, если работаем и не в отпуске.
// time: O(1), space: O(1)
func SetAlarm(employed, vacation bool) bool {
	if employed && !vacation {
		return true
	}
	return false
}
