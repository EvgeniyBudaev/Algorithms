package main

import "fmt"

/* 145. Binary Tree Postorder Traversal
https://leetcode.com/problems/binary-tree-postorder-traversal/description/

Для заданного корня бинарного дерева вернуть обратный обход значений его узлов.

Example 1:
Input: root = [1,null,2,3]
Output: [3,2,1]

Example 2:
Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]
Output: [4,6,7,5,2,9,8,3,1]
*/

func main() {
	root1 := createTree([]interface{}{1, nil, 2, 3})
	fmt.Println(postorderTraversal(root1)) // [3,2,1]

	root2 := createTree([]interface{}{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9})
	fmt.Println(postorderTraversal(root2)) // [4,6,7,5,2,9,8,3,1]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// postorderTraversal возвращает обратный обход значений узлов бинарного дерева.
// time complexity: O(n), где n - количество узлов в дереве
// space complexity: O(h), где h - высота дерева
func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	traverse(root, &result)
	return result
}

// traverse рекурсивно проходит по дереву и добавляет значения узлов в ответ.
func traverse(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	// Post order - проходим по дереву в порядке: левый потомок, правый потомок, корень.
	traverse(node.Left, result)         // Добавляем значение левого потомка
	traverse(node.Right, result)        // Добавляем значение правого потомка
	*result = append(*result, node.Val) // Добавляем значение текущего узла в ответ
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
