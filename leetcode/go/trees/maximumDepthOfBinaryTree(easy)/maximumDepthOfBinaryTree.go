package main

import "fmt"

/* 104. Maximum Depth of Binary Tree
https://leetcode.com/problems/maximum-depth-of-binary-tree/description/

Учитывая корень бинарного дерева, вернуть его максимальную глубину.
Максимальная глубина бинарного дерева — это количество узлов на самом длинном пути от корневого узла до самого дальнего
листового узла.

Input: root = [3,9,20,null,null,15,7]
Output: 3
*/

func main() {
	root := createTree([]interface{}{3, 9, 20, nil, nil, 15, 7})
	fmt.Println(maxDepth(root)) // 3
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxDepth возвращает максимальную глубину дерева.
// time: O(n), где n - количество узлов в дереве., space: O(h), где h - высота дерева.
func maxDepth(root *TreeNode) int {
	if root == nil { // Если дерево пустое, то возвращаем 0
		return 0
	}
	leftDepth := maxDepth(root.Left)      // Рекурсивно вычисляем максимальную глубину левого поддерева
	rightDepth := maxDepth(root.Right)    // Рекурсивно вычисляем максимальную глубину правого поддерева
	return 1 + max(leftDepth, rightDepth) // Возвращаем максимальную глубину дерева, увеличенную на 1 (текущий уровень)
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
