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

// Создаем независимую копию слайса (copy()), чтобы избежать side effect (изменения исходного массива arr).
//func main() {
//	arr := []int{1, 2, 3}
//	src := make([]int, 1, 2) // создаём новый слайс с cap=2 (чтобы append не пересоздавал массив)
//	copy(src, arr[:1])       // копируем данные из arr[:1] в src
//
//	src = foo(src) // работаем с независимой копией
//
//	fmt.Println("src:", src) // [1 5]
//	fmt.Println("arr:", arr) // [1 2 3] (не изменился)
//}

// Используем полный срез (...) для создания копии.
//func main() {
//	arr := []int{1, 2, 3}
//	src := append([]int{}, arr[:1]...) // создаём копию arr[:1]
//
//	src = foo(src) // работаем с копией
//
//	fmt.Println("src:", src) // [1 5]
//	fmt.Println("arr:", arr) // [1 2 3] (не изменился)
//}
