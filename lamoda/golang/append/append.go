package main

import "fmt"

// foo принимает копию слайса.
// time: O(1), space: O(1)
func foo(src []int) {
	src = append(src, 5)
}

func main() {
	arr := []int{1, 2, 3}
	src := arr[:1] // src = [1], len=1, cap=3 (берёт базовый массив `arr`)

	foo(src) // передаётся копия слайса `src`

	fmt.Println("src:", src) // [1] (len=1, cap=3)
	fmt.Println("arr:", arr) // [1 5 3] (изменился, т.к. `append` записал 5 в массив)
}
