package main

import "fmt"

/* 234. Palindrome Linked List
https://leetcode.com/problems/palindrome-linked-list/

Если задан заголовок односвязного списка, вернуть значение true, если он является палиндромом, или false в противном случае.

Example 1:
Input: head = [1,2,2,1]
Output: true

Example 2:
Input: head = [1,2]
Output: false

*/

func main() {
	head := createList([]int{1, 2, 2, 1})
	fmt.Println(isPalindrome(head)) // true
	head2 := createList([]int{1, 2})
	fmt.Println(isPalindrome(head2)) // false
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// isPalindrome возвращает true, если связный список является палиндромом, иначе - false.
// time complexity: O(n), где n - количество узлов в списке
// memory complexity: O(1)
func isPalindrome(head *ListNode) bool {
	// Если список пуст или содержит только один узел, он является палиндромом
	if head == nil || head.Next == nil {
		return true
	}

	// Находим конец первой половины
	firstHalfEnd := findMiddleNode(head)
	// Разворачиваем вторую половину
	secondHalfStart := reverseList(firstHalfEnd)

	// Сравниваем две половины
	p1 := head            // Указатель на первый узел первой половины
	p2 := secondHalfStart // Указатель на первый узел второй половины

	for p1 != nil && p2 != nil {
		// Сравниваем значения текущих узлов первой и второй половин
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// Восстанавливаем оригинальный список (опционально)
	//firstHalfEnd.Next = reverseList(secondHalfStart)

	return true
}

// reverseList переворачивает связный список и возвращает перевернутый список.
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode // Переменная для хранения предыдущего узла
	curr := head       // Текущий узел, начинаем с головы списка

	for curr != nil {
		tmp := curr      // Сохраняем текущий узел в tmp, чтобы не потерять ссылку на него
		curr = curr.Next // Перемещаем текущий узел на следующий узел
		tmp.Next = prev  // Переворачиваем указатель следующего узла, чтобы он указывал на предыдущий узел
		prev = tmp       // Перемещаем предыдущий узел на текущий узел, чтобы продолжить процесс переворота
	}

	// Возвращаем новый заголовок списка, который теперь является последним узлом
	return prev
}

// findMiddleNode возвращает средний узел связанного списка.
func findMiddleNode(head *ListNode) *ListNode {
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
