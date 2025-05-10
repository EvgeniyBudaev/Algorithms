package main

import "fmt"

/* 141. Linked List Cycle
https://leetcode.com/problems/linked-list-cycle/description/

Если задан head, head связанного списка, определите, есть ли в связанном списке цикл.

Цикл в связанном списке есть, если в списке есть некоторый узел, к которому можно снова прийти, непрерывно следуя
указателю next. Внутри pos используется для обозначения индекса узла, к которому подключен указатель next в tail.
Обратите внимание, что pos не передается как параметр.

Верните true, если в связанном списке есть цикл. В противном случае верните false.0.

Input: head = [3,2,0,-4], pos = 1
Output: true
Пояснение: В связанном списке есть цикл, где хвост соединяется с 1-м узлом (индексированным как 0).
*/

func main() {
	head := createList([]int{3, 2, 0, -4})
	fmt.Println(hasCycle(head)) // true
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle определяет, есть ли цикл в связанном списке.
// time: O(n), space: O(1)
func hasCycle(head *ListNode) bool {
	slow := head // Медленный указатель
	fast := head // Быстрый указатель

	for fast != nil && fast.Next != nil { // Проверяем наличие узлов в списке
		slow = slow.Next      // Перемещаем медленный указатель на один шаг
		fast = fast.Next.Next // Перемещаем быстрый указатель на два шага

		if slow == fast { // Если медленный и быстрый указатели сходятся, есть цикл
			return true // Цикл найден
		}
	}
	return false // Цикл не найден
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
