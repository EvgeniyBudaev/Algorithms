package main

import (
	"fmt"
	"math"
)

func main() {
	graph := make(map[string]map[string]int)
	graph["start"] = map[string]int{}
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2

	graph["a"] = map[string]int{}
	graph["a"]["finish"] = 1

	graph["b"] = map[string]int{}
	graph["b"]["a"] = 3
	graph["b"]["finish"] = 5

	graph["finish"] = map[string]int{}

	costs, parents := findShortestPath(graph, "start", "finish")
	fmt.Println(costs, parents) // map[a:5 b:2 finish:6] map[a:b b:start finish:a]
}

// findShortestPath находит кратчайший путь с помощью алгоритма Дейкстры.
func findShortestPath(graph map[string]map[string]int, startNode string, finishNode string) (map[string]int, map[string]string) {
	costs := make(map[string]int)     // Список весов
	costs[finishNode] = math.MaxInt32 // Значение по умолчанию для бесконечности

	parents := make(map[string]string) // Список родителей
	parents[finishNode] = ""           // Значение по умолчанию для конечной вершины

	processed := make(map[string]bool) // Список обработанных вершин

	// Инициализация весов и родителей
	for node, cost := range graph[startNode] {
		costs[node] = cost        // Задаем вес
		parents[node] = startNode // Задаем родителя
	}
	// costs map[a:6, b:2, finish:2147483647]
	// parents map[a:start, b:start, finish:]

	lowestCostNode := findLowestCostNode(costs, processed) // Находим вершину с наименьшим весом // b
	for lowestCostNode != "" {
		// Расчет весов для соседей
		for node, cost := range graph[lowestCostNode] {
			newCost := costs[lowestCostNode] + cost // Расчет нового веса
			// Если новый вес меньше текущего, значит, нужно обновить вес
			if newCost < costs[node] {
				// Устанавливаем новый вес для этого узла
				costs[node] = newCost          // Обновляем вес
				parents[node] = lowestCostNode // Обновляем родителя
			}
			// costs map[a:5, b:2, finish:2147483647] -> costs map[a:5, b:2, finish:7] -> costs map[a:5, b:2, finish:6]
			// parents map[a:b, b:start, finish:] -> parents map[a:b, b:start, finish:b] -> parents map[a:b, b:start, finish:a]
		}

		processed[lowestCostNode] = true                      // Отметить вершину как обработанную
		lowestCostNode = findLowestCostNode(costs, processed) // Находим следующую вершину с наименьшим весом
	}

	// Восстановление пути
	return costs, parents
}

// findLowestCostNode находит вершину с наименьшей стоимостью.
func findLowestCostNode(costs map[string]int, processed map[string]bool) string {
	lowestCost := math.MaxInt32 // Наименьшая стоимость
	lowestCostNode := ""        // Вершина с наименьшей стоимостью
	for node, cost := range costs {
		// Если это самая низкая стоимость на данный момент и она еще не обработана...
		if _, inProcessed := processed[node]; cost < lowestCost && !inProcessed {
			lowestCost = cost     // ... установите его как новую вершину с наименьшим весом.
			lowestCostNode = node // Вершина с наименьшим весом
		}
	}

	// Вернуть вершину с наименьшим весом
	return lowestCostNode
}
