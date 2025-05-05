package main

import "fmt"

/* 100. Same Tree
https://leetcode.com/problems/same-tree/description/

Имея корни двух бинарных деревьев p и q, напишите функцию для проверки их идентичности.
Два бинарных дерева считаются одинаковыми, если они структурно идентичны, а узлы имеют одинаковое значение.

Example 1:
Input: p = [1,2,3], q = [1,2,3]
Output: true

Example 2:
Input: p = [1,2], q = [1,null,2]
Output: false
*/

// Проверка идентичности бинарных деревьев.
func main() {
	p := createTree([]interface{}{1, 2, 3})
	q := createTree([]interface{}{1, 2, 3})
	fmt.Println(isSameTree(p, q)) // true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSameTree проверяет, являются ли два дерева идентичными.
// time complexity: O(n), где n - количество узлов в дереве
// space complexity: O(h), где h - высота дерева. В худшем случае h = n.
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Если одно из деревьев nil, проверяем оба
	if p == nil || q == nil {
		return p == nil && q == nil
	}
	// Проверяем значения текущих узлов
	if p.Val != q.Val {
		return false
	}
	// Рекурсивно проверяем левые и правые поддеревья
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
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
