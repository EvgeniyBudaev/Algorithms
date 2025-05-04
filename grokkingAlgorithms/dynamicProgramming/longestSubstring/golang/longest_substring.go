package main

import "fmt"

// Самая длинная общая подстрока
func main() {
	fmt.Println(substring("vista", "hish")) // "is"
	fmt.Println(substring("fish", "hish"))  // "ish"
}

// createMatrix - функция для создания матрицы размера rows на cols
func createMatrix(rows, cols int) [][]int {
	cell := make([][]int, rows)
	for i := range cell {
		cell[i] = make([]int, cols)
	}

	return cell
}

// substring - функция для поиска самой длинной общей подстроки в строках
func substring(a, b string) string {
	lcs := 0                                 // наибольшая общая подстрока
	lastSubIndex := 0                        // индекс последней подстроки
	cell := createMatrix(len(a)+1, len(b)+1) // матрица размером len(a)+1 на len(b)+1

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			// если символы совпадают, увеличиваем значение ячейки на 1
			if a[i-1] == b[j-1] {
				// если символы не совпадают, выбираем максимальное значение из предыдущих ячеек
				cell[i][j] = cell[i-1][j-1] + 1
				// Если ячейка выше больше, чем ячейка слева, то присваиваем ячейке текущей ячейку выше
				if cell[i][j] > lcs {
					lcs = cell[i][j] // наибольшая общая подстрока
					lastSubIndex = i // индекс последней подстроки
				}
				// Если символы не совпадают, выбираем максимальное значение из предыдущих ячеек
			} else {
				cell[i][j] = 0
			}
		}
	}

	// возвращаем подстроку
	return a[lastSubIndex-lcs : lastSubIndex]
}
