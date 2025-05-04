package main

import (
	"fmt"
	"math"
	"sort"
)

// Fruit представляет структуру фрукта
type Fruit struct {
	Size   float64
	Weight float64
	Type   string
}

// FruitDistance содержит фрукт и расстояние до него
type FruitDistance struct {
	Fruit    Fruit
	Distance float64
}

// Данные о фруктах
var fruitsData = []Fruit{
	{5, 150, "orange"},
	{6, 180, "grapefruit"},
	// Можно добавить больше данных о фруктах здесь
}

// euclideanDistance вычисляет евклидово расстояние между двумя фруктами
func euclideanDistance(a, b Fruit) float64 {
	return math.Sqrt(math.Pow(a.Size-b.Size, 2) + math.Pow(a.Weight-b.Weight, 2))
}

// findKNearestNeighbors находит k ближайших соседей
func findKNearestNeighbors(newPoint Fruit, k int) []Fruit {
	var distances []FruitDistance

	for _, fruit := range fruitsData {
		dist := euclideanDistance(newPoint, fruit)
		distances = append(distances, FruitDistance{fruit, dist})
	}

	// Сортировка по возрастанию расстояния
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})

	// Выбираем первые k элементов
	var neighbors []Fruit
	for i := 0; i < k && i < len(distances); i++ {
		neighbors = append(neighbors, distances[i].Fruit)
	}

	return neighbors
}

// classifyFruit классифицирует новый фрукт
func classifyFruit(newPoint Fruit, k int) string {
	neighbors := findKNearestNeighbors(newPoint, k)
	orangeCount := 0
	grapefruitCount := 0

	for _, fruit := range neighbors {
		switch fruit.Type {
		case "orange":
			orangeCount++
		case "grapefruit":
			grapefruitCount++
		}
	}

	if orangeCount > grapefruitCount {
		return "orange"
	}
	return "grapefruit"
}

// K-ближайших соседей для классификации фруктов
func main() {
	newFruit := Fruit{7, 200, ""}
	classification := classifyFruit(newFruit, 3)
	fmt.Printf("The new fruit is classified as: %s\n", classification) // grapefruit
}
