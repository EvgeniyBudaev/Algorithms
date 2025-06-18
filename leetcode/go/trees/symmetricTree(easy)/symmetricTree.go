package main

import "fmt"

/* 101. Symmetric Tree
https://leetcode.com/problems/symmetric-tree/description/

Дан корень двоичного дерева, проверьте, является ли он зеркальным отражением самого себя
(т.е. симметричен ли относительно своего центра).

Example 1:
Input: root = [1,2,2,3,4,4,3]
Output: true

Example 2:
Input: root = [1,2,2,null,3,null,3]
Output: false
*/

// Проверка симметричности бинарного дерева.
func main() {
	root1 := createTree([]interface{}{1, 2, 2, 3, 4, 4, 3})
	fmt.Println(isSymmetric(root1)) // true

	root2 := createTree([]interface{}{1, 2, 2, nil, 3, nil, 3})
	fmt.Println(isSymmetric(root2)) // false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSymmetric проверяет, является ли двоичное дерево симметричным относительно своего центра.
// time: O(n), где n - количество узлов в дереве, space: O(h), где h - высота дерева.
func isSymmetric(root *TreeNode) bool {
	// Проверяем, является ли дерево пустым
	if root == nil {
		return true
	}
	// Проверяем симметричность двух поддеревьев
	return check(root.Left, root.Right)
}

// check рекурсивно проверяет симметричность двух поддеревьев.
// time: O(n), где n - количество узлов в дереве, space: O(h), где h - высота дерева.
func check(left, right *TreeNode) bool {
	// Проверяем, являются ли оба поддерева пустыми
	if left == nil || right == nil {
		return left == nil && right == nil
	}
	// Проверяем, равны ли значения корневых узлов
	if left.Val != right.Val {
		return false
	}
	return check(left.Left, right.Right) && check(left.Right, right.Left)
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
