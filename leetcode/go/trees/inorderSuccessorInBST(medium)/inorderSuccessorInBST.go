package main

import (
	"fmt"
	"strconv"
	"strings"
)

/* 285. Inorder Successor in BST
https://leetcode.ca/2016-09-10-285-Inorder-Successor-in-BST/

Если задан корень бинарного дерева поиска и узел p в нем, вернуть упорядоченного преемника этого узла в BST.
Если у заданного узла нет упорядоченного преемника в дереве, вернуть null.

Узел-преемник узла p — это узел с наименьшим ключом, большим, чем p.val.

Input: root = [2,1,3], p = 1
Output: 2
Объяснение: следующий узел 1 по порядку — 2. Обратите внимание, что и p, и возвращаемое значение имеют тип TreeNode.
*/

func main() {
	root := createTree([]interface{}{2, 1, 3})
	p := createTree([]interface{}{1})
	fmt.Println(printTree(inorderSuccessor(root, p))) // 2
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// inorderSuccessor возвращает упорядоченного преемника узла p в бинарном дереве поиска.
// time: O(n), где n - количество узлов в дереве.
// space: O(1)
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var ans *TreeNode // Переменная для хранения упорядоченного преемника

	for root != nil { // Обход дерева в порядке обхода в глубину
		// Если значение текущего узла больше значения узла p, то следующий узел по порядку может быть в левом поддереве.
		if root.Val > p.Val {
			ans = root       // Обновляем упорядоченного преемника до текущего узла
			root = root.Left // Переходим в левое поддерево
			// Если значение текущего узла меньше или равно значению узла p, то следующий узел по порядку может быть в правом поддереве
		} else {
			root = root.Right // Переходим в правое поддерево
		}
	}

	return ans // Возвращаем упорядоченного преемника
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
