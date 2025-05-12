package main

import "fmt"

/* 103. Binary Tree Zigzag Level Order Traversal
https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/description/

Учитывая корень двоичного дерева, вернуть зигзагообразный порядок обхода значений его узлов (т. е. слева направо, затем
справа налево для следующего уровня и попеременно).

Input: root = [3,9,20,null,null,15,7]
Output: [[3],[20,9],[15,7]]
*/

func main() {
	root := createTree([]interface{}{3, 9, 20, nil, nil, 15, 7})
	fmt.Println(zigzagLevelOrder(root)) // [[3],[20,9],[15,7]]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// zigzagLevelOrder выполняет зигзагообразный обход дерева и возвращает результат.
// time: O(n), где n - количество узлов в дереве
// space: O(n), где n - количество узлов в дереве
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := [][]int{} // Инициализируем пустой результат
	if root == nil {    // Если корень равен nil, возвращаем пустой результат
		return result
	}

	queue := []*TreeNode{root} // Создаем очередь и добавляем корень
	for len(queue) > 0 {       // Пока очередь не пуста
		levelSize := len(queue)          // Получаем количество узлов на текущем уровне
		levelValues := []int{}           // Создаем слайс для значений на текущем уровне
		for i := 0; i < levelSize; i++ { // Обрабатываем все узлы на текущем уровне
			pop := queue[0]   // Берем первый элемент из очереди
			queue = queue[1:] // Удаляем первый элемент из очереди

			if pop.Left != nil { // Если узел имеет левого потомка, добавляем его в очередь
				queue = append(queue, pop.Left) // Добавляем левого потомка в очередь
			}
			if pop.Right != nil { // Если узел имеет правого потомка, добавляем его в очередь
				queue = append(queue, pop.Right) // Добавляем правого потомка в очередь
			}
			if len(result)%2 == 0 { // Если текущий уровень четный, добавляем значения в слайс слева направо
				levelValues = append(levelValues, pop.Val) // Добавляем значение в слайс
			} else { // Если текущий уровень нечетный, добавляем значения в слайс справа налево
				levelValues = append([]int{pop.Val}, levelValues...) // Добавляем значение в начало слайса
			}
		}
		result = append(result, levelValues) // Добавляем слайс значений текущего уровня в результат
	}

	return result // Возвращаем результат
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
