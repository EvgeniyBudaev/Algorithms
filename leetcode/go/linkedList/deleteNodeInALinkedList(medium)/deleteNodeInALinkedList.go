package main

import (
	"fmt"
)

/* 237. Delete Node in a Linked List
https://leetcode.com/problems/delete-node-in-a-linked-list/description/

Есть односвязный заголовок списка, и мы хотим удалить узел node в нем.
Вам дан узел, который нужно удалить node. Вам не будет предоставлен доступ к первому узлу заголовка.
Все значения связанного списка уникальны, и гарантируется, что данный узел node не является последним узлом в связанном списке.
Удалить данный узел. Обратите внимание, что под удалением узла мы не подразумеваем его удаление из памяти. Мы имеем в виду:

Значение данного узла не должно существовать в связанном списке.
Количество узлов в связанном списке должно уменьшиться на единицу.
Все значения до узла должны быть в том же порядке.
Все значения после узла должны быть в том же порядке.

Input: head = [4,5,1,9], node = 5
Output: [4,1,9]
Пояснение: Вам дан второй узел со значением 5, связанный список должен стать 4 -> 1 -> 9 после вызова вашей функции.
*/

// Объединение K отсортированных связных списков
func main() {
	list := createList([]int{4, 5, 1, 9})
	deleteNode(list)
	printList(list) // [4,1,9]
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// deleteNode удаляет узел с указанным значением из связного списка.
// time: O(1), space: O(1)
func deleteNode(node *ListNode) {
	if node == nil || node.Next == nil {
		return // Нельзя удалить последний узел или nil
	}

	// Копируем значение следующего узла в текущий
	node.Val = node.Next.Val

	// Пропускаем следующий узел в списке
	node.Next = node.Next.Next
}

// deleteNodeByValue - удаляет узел с указанным значением из связного списка.
// time: O(n), space: O(1)
func deleteNodeByValue(head **ListNode, val int) {
	if head == nil || *head == nil { // Если список пустой или не существует, просто вернемся
		return
	}

	// Специальный случай: удаление головы списка
	if (*head).Val == val {
		*head = (*head).Next // Обновляем голову списка
		return
	}

	current := *head          // Начинаем с головы списка
	for current.Next != nil { // Идем по списку, пока не найдем узел с указанным значением или не дойдем до конца
		if current.Next.Val == val { // Если нашли узел с указанным значением
			current.Next = current.Next.Next // Пропускаем узел с указанным значением
			return
		}
		current = current.Next // Переходим к следующему узлу
	}
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
