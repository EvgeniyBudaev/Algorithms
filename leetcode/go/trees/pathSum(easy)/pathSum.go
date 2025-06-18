package main

import "fmt"

/* 112. Path Sum
https://leetcode.com/problems/path-sum/description/

При наличии корня бинарного дерева и целочисленной targetSum вернуть true, если дерево имеет путь от корня к листьям,
такой, что сложение всех значений вдоль пути равно targetSum.
Лист — это узел без дочерних элементов.

Example 1:
Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
Output: true
Пояснение: Показан путь от корня к листьям с целевой суммой.
*/

func main() {
	root := createTree([]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1})
	targetSum := 22
	fmt.Println(hasPathSum(root, targetSum)) // true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// hasPathSum проверяет, есть ли в дереве путь от корня до листа, сумма чисел которого равна заданной сумме.
// time: O(n), где n - количество узлов в дереве
// space: O(h), где h - высота дерева. В худшем случае h = n.
func hasPathSum(root *TreeNode, targetSum int) bool {
	return checkPathSum(root, 0, targetSum)
}

// checkPathSum проверяет наличие пути от заданного узла до листа, сумма чисел которого равна заданной сумме.
// time: O(n), где n - количество узлов в дереве
// space: O(h), где h - высота дерева. В худшем случае h = n.
func checkPathSum(node *TreeNode, currentSum, targetSum int) bool {
	// Если узел пустой, возвращаем false
	if node == nil {
		return false
	}

	// currentSum - префиксная сумма.
	// Добавляем значение текущего узла к сумме
	currentSum += node.Val

	// Если это лист и сумма совпадает
	if isLeaf(node) && currentSum == targetSum {
		return true
	}

	// Проверяем левое и правое поддеревья
	return checkPathSum(node.Left, currentSum, targetSum) ||
		checkPathSum(node.Right, currentSum, targetSum)
}

// isLeaf проверяет, является ли узел листом.
func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
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
