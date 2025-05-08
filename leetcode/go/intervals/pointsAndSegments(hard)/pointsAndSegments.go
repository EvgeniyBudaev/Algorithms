package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

/* Точки и отрезки

Дано n отрезков на числовой прямой и m точек на этой же прямой. Для каждой из данных точек определите,
скольким отрезкам они принадлежат. Точка x считается принадлежащей отрезку с концами a и b, если выполняется двойное
неравенство min (a, b) ≤ x ≤ max (a, b) .

Входные данные
Первая строка содержит два целых числа n (1≤ n ≤ 10 в степени 5 ) – число отрезков и m (1≤ m ≤10 в степени 5 ) – число точек.
В следующих n строках по два целых числи a i и b i – координаты концов соответствующего отрезка(например, может быть
пара 5 2). В последней строке m целых чисел – координаты точек. Все числа по модулю не превосходят 10 9

Выходные данные
В выходной файл выведите m чисел – для каждой точки количество отрезков, в которых она содержится.

Example 1:
Input:
3 2
0 5
-3 2
7 10
1 6

Output: 2 0
*/

// Point - структура для хранения точек и отрезков.
type Point struct {
	coord int // координата
	typ   int // тип точки: +1 для отрезка, 0 для точки
	index int // индекс точки
}

// time: O(n+m*log(n+m)), space: O(n+m)
func main() {
	reader := bufio.NewReader(os.Stdin)  // Создаем новый объект bufio.Reader
	writer := bufio.NewWriter(os.Stdout) // Создаем новый объект bufio.Writer
	defer writer.Flush()                 // Отложенное закрытие writer

	var n, m int              // Количество отрезков и точек
	fmt.Fscan(reader, &n, &m) // Считываем количество отрезков и точек

	points := make([]Point, 0, 2*n+m) // Создаем слайс для хранения точек и отрезков

	// Read segments
	for i := 0; i < n; i++ {
		var a, b int              // Концы отрезка
		fmt.Fscan(reader, &a, &b) // Считываем концы отрезка
		// Сортируем концы отрезка
		if a > b {
			a, b = b, a // Меняем местами
		}
		points = append(points, Point{a, 1, -1})  // Добавляем начало отрезка
		points = append(points, Point{b, -1, -1}) // Добавляем конец отрезка
	}

	// Читаем точки
	inPoints := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &inPoints[i])                   // Считываем точку
		points = append(points, Point{inPoints[i], 0, i}) // Добавляем точку
	}

	// Сотируем точки и отрезки
	sort.Slice(points, func(i, j int) bool {
		// Если координаты равны, то сортируем по типу точки
		if points[i].coord == points[j].coord {
			// Order: +1 (start), 0 (query), -1 (end)
			return points[i].typ > points[j].typ
		}
		// Order: +1 (start), 0 (query), -1 (end)
		return points[i].coord < points[j].coord
	})

	result := make([]int, m) // Создаем слайс для хранения результатов
	current := 0             // Текущее количество отрезков

	// Обрабатываем точки и отрезки
	for _, p := range points {
		// Если точка, то увеличиваем текущее количество отрезков
		if p.typ == 0 {
			result[p.index] = current // Записываем результат
			// Если отрезок, то увеличиваем текущее количество отрезков
		} else {
			current += p.typ
		}
	}

	// Печатаем результаты
	for i, val := range result {
		// Если не первая точка, то печатаем пробел
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		// Печатаем результат
		fmt.Fprint(writer, val)
	}
	// Печатаем перевод строки
	fmt.Fprintln(writer)
}
