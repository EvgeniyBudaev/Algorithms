package main

import "fmt"

/* 230. Kth Smallest Element in a BST
https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/

Учитывая корень двоичного дерева поиска и целое число k, вернуть k-е наименьшее значение (индексированное на 1) среди
всех значений узлов в дереве.

Input: root = [3,1,4,null,2], k = 1
Output: 1
*/

func main() {
	root := createTree([]interface{}{3, 1, 4, nil, 2})
	fmt.Println(kthSmallest(root, 1)) // 1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// kthSmallest возвращает k-е наименьшее значение в двоичном дереве поиска.
// time: O(n), где n - количество узлов в дереве.
// space: O(h), где h - высота дерева.
func kthSmallest(root *TreeNode, k int) int {
	var result int                    // Переменная для хранения результата
	var count int                     // Переменная для подсчета количества обработанных узлов
	inOrder(root, k, &count, &result) // Вызываем функцию для обхода дерева
	return result                     // Возвращаем результат
}

// inOrder выполняет обход дерева в порядке возрастания значений и находит k-е наименьшее значение.
// time: O(n), где n - количество узлов в дереве.
// space: O(h), где h - высота дерева.
func inOrder(node *TreeNode, k int, count *int, result *int) {
	if node == nil { // Если узел равен nil, то выходим из функции
		return
	}

	// Обходим левое поддерево
	inOrder(node.Left, k, count, result)

	// Обрабатываем текущий узел
	*count++         // Увеличиваем счетчик
	if *count == k { // Если счетчик равен k, то записываем значение текущего узла в результат и выходим из функции
		*result = node.Val // Записываем значение текущего узла в результат
		return
	}

	// Обходим правое поддерево
	inOrder(node.Right, k, count, result)
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
