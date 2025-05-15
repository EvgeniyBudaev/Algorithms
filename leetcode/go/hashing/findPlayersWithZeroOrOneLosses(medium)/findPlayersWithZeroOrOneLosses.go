package main

import (
	"fmt"
	"sort"
)

/* 2225. Find Players With Zero or One Losses
https://leetcode.com/problems/find-players-with-zero-or-one-losses/description/

Вам дан целочисленный массив matchs, где match[i] = [winneri,loseri] указывает, что игрок-победитель победил
игрока-проигравшего в матче.
Вернуть ответ в виде списка размером 2, где:

ответ[0] — список всех игроков, которые не проиграли ни одного матча.
ответ[1] — список всех игроков, проигравших ровно один матч.
Значения в двух списках должны возвращаться в порядке возрастания.

Примечание:
Следует учитывать только игроков, сыгравших хотя бы один матч.
Тестовые случаи будут сгенерированы таким образом, чтобы никакие два совпадения не имели одинаковый результат.

Input: matches = [[1,3],[2,3],[3,6],[5,6],[5,7],[4,5],[4,8],[4,9],[10,4],[10,9]]
Output: [[1,2,10],[4,5,7,8]]
Explanation:
Players 1, 2, and 10 have not lost any matches.
Players 4, 5, 7, and 8 each have lost one match.
Players 3, 6, and 9 each have lost two matches.
Thus, answer[0] = [1,2,10] and answer[1] = [4,5,7,8].

Input: matches = [[2,3],[1,3],[5,4],[6,4]]
Output: [[1,2,5,6],[]]
Explanation:
Players 1, 2, 5, and 6 have not lost any matches.
Players 3 and 4 each have lost two matches.
Thus, answer[0] = [1,2,5,6] and answer[1] = [].
*/

func main() {
	matches := [][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}
	fmt.Println(findWinners(matches)) // [[1,2,10],[4,5,7,8]]
}

// findWinners - функция для нахождения игроков, которые не проиграли ни одного матча или с одним проигрышем.
// time: O(n), space: O(n), где n - количество матчей
func findWinners(matches [][]int) [][]int {
	playerLosses := make(map[int]int) // Карта для подсчета проигрышей для каждого игрока

	// Подсчет проигрышей для каждого игрока
	for _, match := range matches { // Проходим по матчам
		winner := match[0] // Победитель
		loser := match[1]  // Проигравший

		// Инициализируем победителя, если его нет (только если мы хотим отслеживать всех игроков)
		if _, exists := playerLosses[winner]; !exists {
			playerLosses[winner] = 0 // Инициализируем победителя с нулевым количеством проигрышей
		}

		// Увеличить количество проигрышей проигравшего
		playerLosses[loser]++
	}

	// Подготавливаем фрагменты ответов
	answer := make([][]int, 2) // 2 фрагмента ответа
	answer[0] = make([]int, 0) // Игроки с 0 поражениями
	answer[1] = make([]int, 0) // Игроки с 1 поражением

	// Классифицируем игроков
	for player, losses := range playerLosses { // Проходим по карте с подсчетом проигрышей
		switch losses { // Классифицируем игроков по количеству проигрышей
		case 0:
			answer[0] = append(answer[0], player) // Добавляем игрока с 0 проигрышами
		case 1:
			answer[1] = append(answer[1], player) // Добавляем игрока с 1 проигрышем
		}
	}

	// Сортируем результаты
	sort.Ints(answer[0]) // Сортируем игроков с 0 проигрышами
	sort.Ints(answer[1]) // Сортируем игроков с 1 проигрышем

	return answer // Возвращаем ответ
}
