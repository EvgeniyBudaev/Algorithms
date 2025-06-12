package main

import (
	"fmt"
)

/* Function 1 - hello world
https://www.codewars.com/kata/523b4ff7adca849afe000035/train/go

Создайте простую функцию под названием greet, которая возвращает самое известное «hello world!».

Очки стиля
Конечно, это так просто, как только можно. Но насколько вы умны, чтобы создать самое креативное «hello world»,
которое только можете придумать? Какое решение «hello world» вы хотели бы показать своим друзьям?
*/

func main() {
	fmt.Println(greet()) // "hello world!"
}

// greet возвращает строку "hello world!".
// time: O(1), space: O(1)
func greet() string {
	return "hello world!"
}
