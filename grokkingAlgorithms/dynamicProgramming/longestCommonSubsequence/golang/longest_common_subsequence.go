package main

import "fmt"

// Самая длинная общая подстрока
func main() {
	fmt.Println(subsequence("fish", "fosh")) // 3
	fmt.Println(subsequence("fort", "fosh")) // 2
}

// createMatrix - функция для создания матрицы размера rows на cols
func createMatrix(rows, cols int) [][]int {
	cell := make([][]int, rows) // Создание массива, который будет хранить ссылки на строки

	for i := range cell {
		cell[i] = make([]int, cols)
	}

	return cell
}

// subsequence возвращает максимальную длину последовательности в строках.
func subsequence(a, b string) int {
	cell := createMatrix(len(a)+1, len(b)+1) // Создание матрицы размером len(a)+1 на len(b)+1

	for i := 1; i <= len(a); i++ { // Заполнение матрицы
		for j := 1; j <= len(b); j++ { // Проверка совпадения символов
			// Если символы совпадают, увеличиваем значение ячейки на 1
			if a[i-1] == b[j-1] {
				cell[i][j] = cell[i-1][j-1] + 1
				// Если символы не совпадают, выбираем максимальное значение из предыдущих ячеек
			} else {
				// Выбираем максимальное значение из ячеек выше и слева от текущей ячейки
				cell[i][j] = cell[i-1][j]
				// Если ячейка выше больше, чем ячейка слева, то присваиваем ячейке текущей ячейку выше
				if cell[i][j] < cell[i][j-1] {
					cell[i][j] = cell[i][j-1]
				}
			}
		}
	}

	// Возвращаем значение ячейки в правом нижнем углу, которая содержит максимальную длину последовательности
	return cell[len(a)][len(b)]
}
