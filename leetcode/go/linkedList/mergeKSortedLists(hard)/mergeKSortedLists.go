package main

import (
	"fmt"
	"math"
)

/* 23. Merge k Sorted Lists
https://leetcode.com/problems/merge-k-sorted-lists/description/

Вам дан массив из k связанных списков, каждый связанный список отсортирован в порядке возрастания.
Объедините все связанные списки в один отсортированный связанный список и верните его.

Example 1:
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
*/

// Объединение K отсортированных связных списков
func main() {
	list1 := createList([]int{1, 4, 5})
	list2 := createList([]int{1, 3, 4})
	list3 := createList([]int{2, 6})
	lists := []*ListNode{list1, list2, list3}
	result := mergeKLists(lists)
	printList(result) // [1,1,2,3,4,4,5,6]
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeKLists объединяет все связные списки в один отсортированный связный список.
// time complexity: O(n), где n - количество узлов в списке
// space complexity: O(1)
func mergeKLists(lists []*ListNode) *ListNode {
	// Решение будет основано на "нескольких указателях".
	// Т.е. мы каждый раз находим наименьший элемент на которые сейчас указывают списки
	// и двигаем указатель для этого списка.
	// Продолжаем так делать пока все списки не станут указывать на nil.
	dummy := &ListNode{} // Фиктивная начальная нода
	current := dummy     // Указатель на текущую ноду

	for {
		// Находим индекс списка с наименьшей нодой
		minIdx := getMinNodeIdx(lists)
		if minIdx == -1 {
			break // Все списки обработаны
		}

		// Добавляем минимальный узел в результат
		current.Next = lists[minIdx]
		current = current.Next

		// Перемещаем указатель в списке с минимальным узлом
		lists[minIdx] = lists[minIdx].Next
	}

	return dummy.Next // Пропускаем фиктивную ноду
}

// getMinNodeIdx возвращает индекс, в котором текущая нода наименьшая.
func getMinNodeIdx(lists []*ListNode) int {
	minVal := math.MaxInt32 // Максимальное значение int32
	minIdx := -1            // Индекс минимального значения

	for i, node := range lists {
		// Если нода не nil и ее значение меньше текущего минимального, то обновляем минимальное значение и индекс
		if node != nil && node.Val < minVal {
			minVal = node.Val
			minIdx = i
		}
	}

	// Если ни одна нода не была найдена, возвращаем -1
	return minIdx
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
