package main

import (
	"fmt"
	"strings"
)

/* https://leetcode.com/problems/simplify-path/description/

Учитывая абсолютный путь к файловой системе в стиле Unix, который начинается с косой черты «/», преобразуйте этот путь в
упрощенный канонический путь.

В контексте файловой системы в стиле Unix одна точка '.' означает текущий каталог, двойная точка «..» означает переход
на один уровень каталога вверх, а несколько косых черт, таких как «//», интерпретируются как одна косая черта.
В этой задаче рассматривайте последовательности точек, не подпадающие под действие предыдущих правил (например, «...»),
как допустимые имена файлов или каталогов.

Упрощенный канонический путь должен соответствовать следующим правилам:

Он должен начинаться с одинарной косой черты «/».
Каталоги внутри пути должны быть разделены только одной косой чертой «/».
Он не должен заканчиваться косой чертой «/», если только это не корневой каталог.
Следует исключить любые одиночные или двойные точки, используемые для обозначения текущих или родительских каталогов.
Верните новый путь.

Input: path = "/home/"
Output: "/home"

Input: path = "/home//foo/"
Output: "/home/foo"

Input: path = "/home/user/Documents/../Pictures"
Output: "/home/user/Pictures"

Input: path = "/../"
Output: "/"

Input: path = "/.../a/../b/c/../d/./"
Output: "/.../b/d"
*/

func main() {
	fmt.Println(simplifyPath("/home/"))      // "/home"
	fmt.Println(simplifyPath("/home//foo/")) // "/home/foo"
}

func simplifyPath(path string) string {
	var stack []string
	directories := strings.Split(path, "/")

	for _, dir := range directories {
		if dir == "." || dir == "" {
			continue
		} else if dir == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1] // Удаляем последний элемент
			}
		} else {
			stack = append(stack, dir)
		}
	}

	return "/" + strings.Join(stack, "/")
}
