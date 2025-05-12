package main

import (
	"fmt"
	"strconv"
	"strings"
)

/* 116. Populating Next Right Pointers in Each Node
https://leetcode.com/problems/populating-next-right-pointers-in-each-node/description/

Вам дано идеальное двоичное дерево, в котором все листья находятся на одном уровне, а у каждого родителя есть два
потомка. Двоичное дерево имеет следующее определение:

struct Node {
int val;
Node *left;
Node *right;
Node *next;
}

Заполните каждый указатель next, чтобы он указывал на его следующий правый узел. Если следующего правого узла нет,
указатель next должен быть установлен в NULL.
Изначально все указатели next установлены в NULL.

Input: root = [1,2,3,4,5,6,7]
Output: [1,#,2,3,#,4,5,6,7,#]
Пояснение: Учитывая приведенное выше идеальное двоичное дерево (рисунок A), ваша функция должна заполнить каждый
следующий указатель так, чтобы он указывал на его следующий правый узел, как на рисунке B. Сериализованный вывод
располагается в порядке уровней, соединенных указателями next, при этом «#» обозначает конец каждого уровня.
*/

func main() {
	root := createTree([]interface{}{1, 2, 3, 4, 5, 6, 7})
	connected := connect(root)
	fmt.Println(printTreeWithNext(connected)) // [1,#,2,3,#,4,5,6,7,#]
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// connect связывает указатели next узлов в идеальном двоичном дереве.
// time: O(n), space: O(1)
func connect(root *Node) *Node {
	if root == nil { // Если дерево пустое, возвращаем его
		return nil
	}

	// Начинаем с корня
	leftmost := root

	// Пока есть следующий уровень
	for leftmost.Left != nil {
		head := leftmost

		// Проходим по текущему уровню и связываем узлы следующего уровня
		for head != nil {
			// Связываем левого потомка с правым
			head.Left.Next = head.Right

			// Связываем правого потомка с левым потомком следующего узла (если есть)
			if head.Next != nil {
				head.Right.Next = head.Next.Left
			}

			// Переходим к следующему узлу в текущем уровне
			head = head.Next
		}

		// Переходим на следующий уровень
		leftmost = leftmost.Left
	}

	return root // Возвращаем корневой узел
}

// createTree создает идеальное бинарное дерево из массива значений.
func createTree(values []interface{}) *Node {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &Node{Val: values[0].(int)}
	queue := []*Node{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		// Левый потомок
		if i < len(values) && values[i] != nil {
			node.Left = &Node{Val: values[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		// Правый потомок
		if i < len(values) && values[i] != nil {
			node.Right = &Node{Val: values[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// printTreeWithNext выводит дерево в формате [val,nextVal,#,...]
func printTreeWithNext(root *Node) string {
	if root == nil {
		return "[]"
	}

	var result []string
	queue := []*Node{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// Добавляем значение текущего узла
			result = append(result, strconv.Itoa(node.Val))

			// Добавляем значение next (или #)
			if node.Next != nil {
				result = append(result, strconv.Itoa(node.Next.Val))
			} else {
				result = append(result, "#")
			}

			// Добавляем потомков в очередь
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	// Убираем лишние # в конце (опционально)
	for len(result) > 0 && result[len(result)-1] == "#" {
		result = result[:len(result)-1]
	}

	return "[" + strings.Join(result, ",") + "]"
}
