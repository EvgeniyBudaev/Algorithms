package main

import (
	"fmt"
	"math"
)

/* 98. Validate Binary Search Tree
https://leetcode.com/problems/validate-binary-search-tree/description/

Учитывая корень двоичного дерева, определите, является ли оно допустимым двоичным деревом поиска (BST).
Допустимое BST определяется следующим образом:

Левое поддерево узла содержит только узлы с ключами, меньшими ключа узла.
Правильное поддерево узла содержит только узлы с ключами, большими ключа узла.
И левое, и правое поддеревья также должны быть двоичными деревьями поиска.

Input: root = [2,1,3]
Output: true

Input: root = [5,1,4,null,null,3,6]
Output: false
Пояснение: Значение корневого узла равно 5, но значение его правого дочернего узла равно 4.
*/

func main() {
	root := createTree([]interface{}{2, 1, 3})
	fmt.Println(isValidBST(root)) // true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isValidBST проверяет, является ли дерево корректным BST.
// time: O(n), space: O(n)
func isValidBST(root *TreeNode) bool {
	minVale := math.MinInt64 // Минимальное значение, которое может иметь узел
	isValid := true          // Флаг, указывающий, является ли дерево корректным BST

	var visit func(node *TreeNode) // Объявляем функцию visit вне тела функции isValidBST

	// Создаем функцию visit, которая будет обходить дерево и проверять условие BST
	visit = func(node *TreeNode) {
		if node == nil { // Если узел равен nil, то выходим из функции
			return
		}
		visit(node.Left)         // Обходим левое поддерево
		if node.Val <= minVale { // Если значение узла меньше или равно минимальному значению, то дерево не является корректным BST
			isValid = false // Устанавливаем флаг isValid в false
		}
		minVale = node.Val // Обновляем минимальное значение
		visit(node.Right)  // Обходим правое поддерево
	}

	visit(root)    // Вызываем функцию visit с корневым узлом дерева
	return isValid // Возвращаем флаг isValid
}

// createTree создает бинарное дерево из массива значений.
func createTree(values []interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: values[0].(int)} // Создаем корневой узел с первым значением в массиве
	queue := []*TreeNode{root}              // Создаем очередь для обхода дерева по уровням
	i := 1                                  // Индекс для обхода массива значений

	for len(queue) > 0 && i < len(values) {
		node := queue[0]  // Берем первый элемент из очереди
		queue = queue[1:] // Удаляем первый элемент из очереди

		// Левый потомок
		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: values[i].(int)} // Создаем левый потомок и добавляем его в очередь
			queue = append(queue, node.Left)            // Добавляем левого потомка в очередь
		}
		i++

		// Правый потомок
		// Если правый потомок не равен nil, то создаем его и добавляем его в очередь
		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: values[i].(int)} // Создаем правого потомка и добавляем его в очередь
			queue = append(queue, node.Right)            // Добавляем правого потомка в очередь
		}
		i++
	}

	// Возвращаем указатель на корневой узел
	return root
}
