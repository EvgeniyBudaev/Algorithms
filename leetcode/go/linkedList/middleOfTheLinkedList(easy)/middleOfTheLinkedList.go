package main

import "fmt"

/* 876. Middle of the Linked List
https://leetcode.com/problems/middle-of-the-linked-list/description/

Если задан заголовок односвязного списка, вернуть средний узел связанного списка.
Если есть два средних узла, вернуть второй средний узел.

Example 1:
Input: head = [1,2,3,4,5]
Output: [3,4,5]
Пояснение: Средний узел списка — это узел 3.

Example 2:
Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Пояснение: Поскольку в списке есть два средних узла со значениями 3 и 4, мы возвращаем второй.
*/

func main() {
	head := createList([]int{1, 2, 3, 4, 5})
	result := middleNode(head)
	printList(result) // [3, 4, 5]
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// middleNode возвращает средний узел связанного списка.
// time complexity: O(n), где n - количество узлов в списке
// memory complexity: O(1)
func middleNode(head *ListNode) *ListNode {
	slow := head // Скорость медленного указателя, начнем с головы списка
	fast := head // Скорость быстрого указателя, начнем с головы списка

	// Проходим список, пока быстрый указатель не достигнет конца списка или его следующего элемента
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // Перемещаем медленный указатель на один шаг
		fast = fast.Next.Next // Перемещаем быстрый указатель на два шага
	}

	// Возвращаем средний узел
	return slow
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
