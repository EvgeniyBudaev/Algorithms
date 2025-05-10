package main

import (
	"fmt"
	"strconv"
	"strings"
)

/* 108. Convert Sorted Array to Binary Search Tree
https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/

Дан целочисленный массив nums, элементы которого отсортированы в порядке возрастания, преобразовать его в
сбалансированное по высоте двоичное дерево поиска.

Input: nums = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Пояснение: [0,-10,5,null,-3,null,9] также принимается
*/

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	fmt.Println(printTree(sortedArrayToBST(nums))) // [0,-3,9,-10,null,5]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// sortedArrayToBST - рекурсивный алгоритм, который берет средний элемент массива и делит его на два подмассива, которые будут
// являться левым и правым поддеревом. После чего рекурсивно вызывает функцию для каждого подмассива.
// time: O(n), space: O(n)
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	middle := len(nums) / 2                      // Получаем индекс среднего элемента
	left := append([]int{}, nums[0:middle]...)   // Создаем левый подмассив
	right := append([]int{}, nums[middle+1:]...) // Создаем правый подмассив

	// Создаем новый узел с средним элементом и рекурсивно вызываем функцию для каждого подмассива
	return &TreeNode{
		Val:   nums[middle],
		Left:  sortedArrayToBST(left),
		Right: sortedArrayToBST(right),
	}
}

// printTree - функция для вывода дерева в формате [val,left,right]
// time: O(n), space: O(n)
func printTree(root *TreeNode) string {
	if root == nil { // Если дерево пустое, возвращаем пустой срез
		return "[]"
	}

	var result []string        // Срез для хранения значений узлов
	queue := []*TreeNode{root} // Используем очередь для BFS

	for len(queue) > 0 { // BFS
		node := queue[0]  // Берем первый узел из очереди
		queue = queue[1:] // Удаляем его из очереди

		if node == nil { // Если узел пустой, добавляем null в результат
			result = append(result, "null") // добавляем null в результат
			continue                        // Пропускаем остальные шаги
		}

		result = append(result, strconv.Itoa(node.Val)) // Добавляем значение узла в результат
		queue = append(queue, node.Left)                // Добавляем левого потомка в очередь
		queue = append(queue, node.Right)               // Добавляем правого потомка в очередь
	}

	// Удаляем trailing nulls
	i := len(result) - 1                // Инициализируем индекс последнего ненулевого элемента
	for i >= 0 && result[i] == "null" { // Если элемент равен null, уменьшаем индекс
		i-- // уменьшаем индекс
	}
	result = result[:i+1] // Обрезаем ненужные элементы

	return "[" + strings.Join(result, ",") + "]" // Возвращаем результат в виде строки
}
