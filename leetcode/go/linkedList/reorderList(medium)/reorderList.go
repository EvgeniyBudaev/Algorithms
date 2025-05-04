package main

import "fmt"

/* 143. Reorder List
https://leetcode.com/problems/reorder-list/

Вам дан заголовок односвязного списка. Список можно представить как:

L0 → L1 → … → Ln - 1 → Ln
Измените порядок списка, чтобы он имел следующий вид:

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
Вы не можете изменять значения в узлах списка. Изменять можно только сами узлы.

Example 1:
Input: head = [1,2,3,4]
Output: [1,4,2,3]

Example 2:
Input: head = [1,2,3,4,5]
Output: [1,5,2,4,3]

*/

func main() {
	head := createList([]int{1, 2, 3, 4})
	reorderList(head) // [1, 4, 2, 3]
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// reorderList меняет порядок узлов списка на L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
// time complexity: O(n), где n - количество узлов в списке
// space complexity: O(1)
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// Находим узел перед серединой
	preMiddle := preMiddleNode(head)
	// Разворачиваем вторую половину
	secondHalf := reverseList(preMiddle.Next)

	// Разделяем список на две части
	preMiddle.Next = nil

	// Объединяем две половины
	p1 := head
	p2 := secondHalf
	newHead := p1

	for p2 != nil {
		p1Next := p1.Next
		p1.Next = p2
		p1 = p2
		p2 = p1Next
	}

	printList(newHead)
}

// preMiddleNode находит предыдущий узел среднего узла в связанном списке.
// time complexity: O(n), где n - количество узлов в списке
// space complexity: O(1)
func preMiddleNode(head *ListNode) *ListNode {
	slow := head // Скорость медленного указателя, начнем с головы списка
	fast := head // Скорость быстрого указателя, начнем с головы списка

	// Проходим список, пока быстрый указатель не достигнет конца списка или его следующего элемента
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next      // Перемещаем медленный указатель на один шаг
		fast = fast.Next.Next // Перемещаем быстрый указатель на два шага
	}

	// Возвращаем средний узел
	return slow
}

// reverseList переворачивает связный список и возвращает перевернутый список.
// time complexity: O(n), где n - количество узлов в списке
// space complexity: O(1)
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
