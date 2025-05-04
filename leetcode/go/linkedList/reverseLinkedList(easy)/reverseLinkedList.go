package main

import "fmt"

/* 206. Reverse Linked List
https://leetcode.com/problems/reverse-linked-list/description/

Дан заголовок односвязного списка, перевернуть список и вернуть перевернутый список.

Example 1:
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

Example 2:
Input: head = [1,2]
Output: [2,1]

Example 3:
Input: head = []
Output: []
*/

// Реализация разворота связного списка
func main() {
	head := createList([]int{1, 2, 3, 4, 5})
	result := reverseList(head)
	printList(result) // [5, 4, 3, 2, 1]
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList переворачивает связный список и возвращает перевернутый список.
// time complexity: O(n), где n - количество узлов в списке
// space complexity: O(1)
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode // Переменная для хранения предыдущего узла
	curr := head       // Текущий узел, начинаем с головы списка

	for curr != nil {
		next := curr.Next // Сохраняем следующий узел
		curr.Next = prev  // Разворачиваем указатель
		prev = curr       // Перемещаем prev на текущий узел
		curr = next       // Перемещаем curr на следующий узел
	}

	return prev
}

// createList создает связный список из массива значений
func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head

	for _, val := range values[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return head
}

// printList печатает список в формате [v1,v2,v3]
func printList(head *ListNode) {
	fmt.Print("[")
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(",")
		}
		head = head.Next
	}
	fmt.Println("]")
}
