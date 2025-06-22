package main

import (
	"fmt"
	"math"
)

// https://www.youtube.com/watch?v=8WCfY0Uh1iM&ab_channel=%D0%9E%D0%B1%D1%83%D1%87%D0%B5%D0%BD%D0%B8%D0%B5HTML%2CCSS%2CJavaScript%7C%D0%90%D0%BD%D0%B0%D1%82%D0%BE%D0%BB%D0%B8%D0%B9%D0%98%D0%B2%D0%B0%D1%88%D0%BE%D0%B2
// Задача: нужно вывести кратчайший путь из точки А в точку Б

type Graph map[string]map[string]int

// findLowestCostNode находит вершину с наименьшей стоимостью в графе.
// time: O(n), space: O(1), где n - количество вершин графа
func findLowestCostNode(distances map[string]int, visited map[string]bool) string {
	lowestCost := math.MaxInt32 // Инициализируем наибольшим возможным значением
	var lowestCostNode string   // Вершина с наименьшей стоимостью

	for node, cost := range distances { // Перебираем все вершины графа
		if cost < lowestCost && !visited[node] {
			lowestCost = cost
			lowestCostNode = node
		}
	}
	return lowestCostNode
}

// dijkstra алгоритм для нахождения кратчайшего пути между двумя точками в графе.
// time: O(n^2), space: O(n), где n - количество вершин графа
func dijkstra(graph Graph, start, end string) []string {
	distances := make(map[string]int) // Расстояния от начальной вершины до каждой другой вершины
	visited := make(map[string]bool)  // Посещенные вершины
	path := make(map[string]string)   // Путь от начальной вершины до каждой другой вершины

	// Инициализация расстояний
	for node := range graph { // Перебираем все вершины графа
		if node != start { // Если вершина не начальная, то расстояние до нее равно бесконечности
			distances[node] = math.MaxInt32
		} else { // Если вершина начальная, то расстояние до нее равно 0
			distances[start] = 0
		}
	}

	for !visited[end] { // Пока не посетим все вершины
		node := findLowestCostNode(distances, visited)
		if node == "" {
			break // Нет пути
		}

		neighbors := graph[node]                // Соседи текущей вершины
		for neighbor, cost := range neighbors { // Перебираем всех соседей текущей вершины
			newDistance := distances[node] + cost  // Расстояние до соседа через текущую вершину
			if newDistance < distances[neighbor] { // Если новое расстояние меньше старого, то обновляем расстояние и путь
				distances[neighbor] = newDistance // Обновляем расстояние
				path[neighbor] = node             // Обновляем путь
			}
		}

		visited[node] = true // Посетили вершину
	}

	// Восстановление пути
	shortPath := []string{}
	current := end // Текущая вершина

	for current != start { // Пока не дойдем до начальной вершины
		shortPath = append([]string{current}, shortPath...) // Добавляем текущую вершину в начало пути
		if prev, exists := path[current]; exists {          // Если существует предыдущая вершина, то продолжаем восстановление пути
			current = prev // Переходим к предыдущей вершине
		} else { // Иначе, путь не существует
			return []string{} // Нет пути
		}
	}
	shortPath = append([]string{start}, shortPath...) // Добавляем начальную вершину в начало пути

	return shortPath // Возвращаем путь
}

func main() {
	graph := Graph{
		"a": {"b": 2, "c": 1, "i": 3},
		"b": {"a": 2, "d": 3},
		"c": {"a": 1, "d": 1},
		"d": {"b": 3, "c": 1, "e": 5},
		"e": {"d": 5, "i": 2},
		"i": {"a": 3, "e": 2},
	}

	fmt.Println(dijkstra(graph, "a", "e")) // [a i e]
}
