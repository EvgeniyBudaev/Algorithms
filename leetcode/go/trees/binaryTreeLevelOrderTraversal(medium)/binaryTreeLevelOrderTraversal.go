package main

import "fmt"

/* 102. Binary Tree Level Order Traversal
https://leetcode.com/problems/binary-tree-level-order-traversal/description/

Учитывая корень двоичного дерева, вернуть порядок обхода значений его узлов (т. е. слева направо, уровень за уровнем).

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]

Example 2:
Input: root = [1]
Output: [[1]]
*/

func main() {
	root := createTree([]interface{}{3, 9, 20, nil, nil, 15, 7})
	fmt.Println(levelOrder(root)) // [[3],[9,20],[15,7]]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// levelOrder возвращает список значений узлов дерева в порядке обхода уровнями.
// time: O(n), space: O(n), где n - количество узлов в дереве
func levelOrder(root *TreeNode) [][]int {
	levels := make([][]int, 0)
	preOrder(root, 0, &levels)
	return levels
}

// preOrder рекурсивно обходит дерево и заполняет уровни.
// time: O(n), space: O(n)
func preOrder(node *TreeNode, level int, levels *[][]int) *[][]int {
	// Если узел равен nil, возвращаем слайс уровней
	if node == nil {
		return levels
	}

	// Если текущего уровня еще нет в слайсе, добавляем новый уровень
	if level == len(*levels) {
		*levels = append(*levels, []int{})
	}

	// Делаем PreOrder обход дерева и добавляем вершину на текущий уровень в конец списка.
	(*levels)[level] = append((*levels)[level], node.Val)
	// Рекурсивно обходим левое и правое поддеревья
	preOrder(node.Left, level+1, levels)
	preOrder(node.Right, level+1, levels)

	return levels
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
