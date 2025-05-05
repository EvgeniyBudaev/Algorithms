package main

import "fmt"

/* 21. Merge Two Sorted Lists
https://leetcode.com/problems/merge-two-sorted-lists/

Вам даны заголовки двух отсортированных связанных списков list1 и list2.
Объедините два списка в один отсортированный список. Список должен быть создан путем сращивания узлов первых двух списков.
Верните заголовок объединенного связанного списка.

Example 1:
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
*/

// Объединение двух отсортированных связных списков
func main() {
	list1 := createList([]int{1, 2, 4})
	list2 := createList([]int{1, 3, 4})
	result := mergeTwoLists(list1, list2)
	printList(result) // [1,1,2,3,4,4]
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoLists объединяет два отсортированных связных списка в один отсортированный список.
// time complexity: O(n), где n - количество узлов в списке
// space complexity: O(1)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{} // Фиктивная начальная нода
	current := dummy     // Указатель на текущую ноду

	for list1 != nil || list2 != nil {
		// Получаем значения или "бесконечность", если список закончился
		val1 := getVal(list1)
		val2 := getVal(list2)

		// Выбираем меньшее значение и перемещаем указатель на следующую ноду в соответствующем списке
		if val1 <= val2 {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		// Перемещаем указатель на следующую ноду
		current = current.Next
	}

	return dummy.Next // Пропускаем фиктивную ноду
}

// getVal возвращает значение ноды или максимальное int, если нода nil
func getVal(node *ListNode) int {
	if node == nil {
		return 1<<31 - 1 // Максимальное int (аналог float('inf'))
	}
	return node.Val
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
