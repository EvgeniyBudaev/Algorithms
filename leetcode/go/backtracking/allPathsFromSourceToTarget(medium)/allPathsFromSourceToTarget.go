package main

import "fmt"

/* https://leetcode.com/problems/all-paths-from-source-to-target/description/

Дан направленный ациклический граф (DAG) из n узлов, помеченных от 0 до n-1, найдите все возможные пути от узла 0 до
узла n-1 и верните их в любом порядке.

Граф задается следующим образом: граф[i] — это список всех узлов, которые вы можете посетить из узла i
(т. е. существует направленное ребро от узла i к узлу граф[i][j]).

Input: graph = [[1,2],[3],[3],[]]
Output: [[0,1,3],[0,2,3]]
Пояснение: Есть два пути: 0 -> 1 -> 3 и 0 -> 2 -> 3.

Input: graph = [[4,3,1],[3,2,4],[3],[4],[]]
Output: [[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]
*/

func main() {
	// Пример графа в виде списка смежности
	// Узел 0 ведет к узлам 1 и 2
	// Узел 1 ведет к узлу 3
	// Узел 2 ведет к узлу 3
	// Узел 3 не имеет исходящих ребер (конечный узел)
	graph := [][]int{
		{1, 2},
		{3},
		{3},
		{},
	}
	fmt.Println(allPathsSourceTarget(graph)) // [[0,1,3],[0,2,3]]
}

// allPathsSourceTarget находит все пути от начального узла (0) до конечного (n-1) в графе
func allPathsSourceTarget(graph [][]int) [][]int {
	var result [][]int // Здесь будем хранить все найденные пути
	n := len(graph)    // Количество узлов в графе
	path := []int{0}   // Начинаем путь с узла 0

	// Внутренняя рекурсивная функция для обхода в глубину (DFS)
	var dfs func(int, []int)
	dfs = func(node int, currentPath []int) {
		// Если достигли целевого узла, добавляем текущий путь в результат
		if node == n-1 {
			// Важно создать копию пути, чтобы избежать изменений потом
			temp := make([]int, len(currentPath))
			copy(temp, currentPath)
			result = append(result, temp)
			return
		}

		// Рекурсивно посещаем всех соседей текущего узла
		for _, neighbor := range graph[node] {
			// Добавляем соседа в путь и продолжаем поиск
			newPath := append([]int{}, currentPath...)
			newPath = append(newPath, neighbor)
			dfs(neighbor, newPath)
		}
	}

	dfs(0, path) // Запускаем DFS с начального узла
	return result
}
