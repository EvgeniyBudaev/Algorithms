package main

import "fmt"

/* 199. Binary Tree Right Side View
https://leetcode.com/problems/binary-tree-right-side-view/description/

Имея корень двоичного дерева, представьте, что вы стоите с правой стороны от него, и верните значения узлов,
которые вы видите, упорядоченные сверху вниз.

Example 1:
Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]
*/

func main() {
	root := createTree([]interface{}{1, 2, 3, nil, 5, nil, 4})
	fmt.Println(rightSideView(root)) // [1,3,4]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// rightSideView возвращает список значений узлов, которые вы видите с правой стороны, упорядоченные сверху вниз.
// time complexity: O(n), где n - количество узлов в дереве
// space complexity: O(n)
func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)
	preOrder(root, 0, &result)
	return result
}

// preOrder рекурсивно обходит дерево и заполняет уровни
func preOrder(node *TreeNode, level int, list *[]int) *[]int {
	// Если узел равен nil, возвращаем слайс уровней
	if node == nil {
		return list
	}

	// Делаем PreOrder и на каждом уровне запоминаем последнюю вершину для уровня
	// Если текущего уровня еще нет в слайсе, добавляем новый уровень
	if level == len(*list) {
		*list = append(*list, 0)
	}

	// Запоминаем последнюю вершину для уровня
	(*list)[level] = node.Val
	// Рекурсивно обходим левое и правое поддеревья
	preOrder(node.Left, level+1, list)
	preOrder(node.Right, level+1, list)

	return list
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
