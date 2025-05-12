package main

import (
	"fmt"
	"strconv"
	"strings"
)

/* 105. Construct Binary Tree from Preorder and Inorder Traversal
https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/

Даны два целочисленных массива preorder и inorder, где preorder — это предварительный обход двоичного дерева,
а inorder — симметричный обход того же дерева. Построить и вернуть двоичное дерево.

Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
*/

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	fmt.Println(printTree(buildTree(preorder, inorder))) // [3,9,20,null,null,15,7]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// buildTree строит бинарное дерево на основе предварительного и симметричного обхода.
// time: O(n), где n - количество узлов в дереве
// space: O(n), где n - количество узлов в дереве
func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderMap := make(map[int]int) // Создаем карту для быстрого поиска индекса элемента в inorder
	for i, val := range inorder {   // Заполняем карту значениями из inorder и соответствующими индексами
		inorderMap[val] = i // Ключом является значение элемента, а значением - его индекс в inorder
	}

	var i int // Индекс текущего элемента в preorder

	var helper func(j, k int) *TreeNode // Вспомогательная функция для рекурсивного построения дерева
	helper = func(j, k int) *TreeNode { // Вспомогательная функция для рекурсивного построения дерева
		if j > k { // Если индексы пересекаются, то дерево закончилось
			return nil
		}

		nodeVal := preorder[i]          // Значение текущего элемента в preorder
		i++                             // Переходим к следующему элементу в preorder
		node := &TreeNode{Val: nodeVal} // Создаем узел с текущим значением
		idx := inorderMap[nodeVal]      // Находим индекс текущего элемента в inorder
		node.Left = helper(j, idx-1)    // Рекурсивно строим левое поддерево
		node.Right = helper(idx+1, k)   // Рекурсивно строим правое поддерево
		return node                     // Возвращаем созданный узел
	}

	return helper(0, len(inorder)-1) // Вызываем вспомогательную функцию с начальными индексами
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
