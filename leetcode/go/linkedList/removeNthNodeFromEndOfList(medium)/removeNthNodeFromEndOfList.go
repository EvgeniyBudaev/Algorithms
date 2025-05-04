package main

import "fmt"

/* 19. Remove Nth Node From End of List
https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/

Дан заголовок связанного списка, удалить n-й узел из конца списка и вернуть его заголовок.

Example 1:
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Example 2:
Input: head = [1], n = 1
Output: []
*/

func main() {
	head1 := createList([]int{1, 2, 3, 4, 5})
	n := 2
	result1 := removeNthFromEnd(head1, n)
	printList(result1) // [1, 2, 3, 5]

	head2 := createList([]int{1, 2, 3, 4, 5})
	result2 := removeNthFromEndVariantSlowFast(head2, n)
	printList(result2) // [1, 2, 3, 5]
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// removeNthFromEnd удаляет n-й узел из конца списка и возвращает его заголовок
// Паттерн Dummy
// time complexity: O(n), где n - количество узлов в списке
// space complexity: O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Создаем фиктивный узел dummyNode, чтобы избежать специальных случаев с пустым списком или удалением первого узла.
	dummyNode := &ListNode{Val: 0, Next: head}

	// Вычисляем длину списка с учетом dummyNode
	length := 0
	curr := dummyNode
	for curr != nil {
		curr = curr.Next
		length++
	}

	// Доходим до (n-1)-ой вершины с конца
	curr = dummyNode
	// Тут условие i < length-n-1, потому что нам нужно дойти до (n-1)-ой вершины с конца, а не до n-ой вершины.
	for i := 0; i < length-n-1; i++ {
		curr = curr.Next // Двигаемся к следующему узлу
	}

	// Удаляем вершину
	// Проверяем, что узел существует, прежде чем удалить его
	if curr.Next != nil {
		// Удаляем вершину, изменяя ссылку на следующую вершину
		curr.Next = curr.Next.Next
	}

	// Возвращаем заголовок списка, начиная с dummyNode
	return dummyNode.Next
}

// removeNthFromEndVariantSlowFast удаляет n-й узел из конца списка и возвращает его заголовок
// time complexity: O(n), где n - количество узлов в списке
// space complexity: O(1)
func removeNthFromEndVariantSlowFast(head *ListNode, n int) *ListNode {
	// Создаем фиктивный узел dummyNode, чтобы избежать специальных случаев с пустым списком или удалением первого узла.
	dummyNode := &ListNode{Val: 0, Next: head}

	fast := dummyNode
	// Перемещаем быстрый указатель на n+1 позиций вперед
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	slow := dummyNode
	// Перемещаем оба указателя, пока быстрый не достигнет конца
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// Удаляем n-й узел с конца
	slow.Next = slow.Next.Next

	// Возвращаем заголовок списка, начиная с dummyNode
	return dummyNode.Next
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
