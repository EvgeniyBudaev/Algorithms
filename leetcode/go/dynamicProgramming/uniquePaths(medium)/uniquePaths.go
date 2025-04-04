package main

import "fmt"

/* https://leetcode.com/problems/unique-paths/description/

На сетке m x n находится робот. Робот изначально расположен в верхнем левом углу (т. е. сетка[0][0]). Робот пытается
переместиться в правый нижний угол (т. е. сетку[m - 1][n - 1]). В любой момент времени робот может двигаться только вниз
или вправо.

Учитывая два целых числа m и n, верните количество возможных уникальных путей, по которым робот может добраться до
правого нижнего угла.

Тестовые случаи генерируются так, что ответ будет меньше или равен 2 * 109.

Input: m = 3, n = 7
Output: 28

Input: m = 3, n = 2
Output: 3
Пояснение: Из верхнего левого угла можно попасть в правый нижний угол тремя способами:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down
*/

func main() {
	fmt.Println(uniquePaths(3, 7)) // 28
}

// uniquePaths вычисляет количество уникальных путей из верхнего левого угла
func uniquePaths(m int, n int) int {
	// Создаем временный массив для хранения промежуточных результатов
	aboveRow := make([]int, n)

	// Инициализируем первую строку единицами
	for i := range aboveRow {
		aboveRow[i] = 1
	}

	// Заполняем последующие строки
	for row := 1; row < m; row++ {
		currentRow := make([]int, n)
		currentRow[0] = 1 // Первый элемент в строке всегда 1

		for col := 1; col < n; col++ {
			// Количество путей равно сумме путей сверху и слева
			currentRow[col] = currentRow[col-1] + aboveRow[col]
		}

		aboveRow = currentRow // Обновляем aboveRow для следующей итерации
	}

	return aboveRow[n-1]
}
