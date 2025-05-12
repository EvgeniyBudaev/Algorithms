package main

import "fmt"

/* 160. Intersection of Two Linked Lists
https://leetcode.com/problems/intersection-of-two-linked-lists/description/

Даны заголовки двух односвязных списков headA и headB, вернуть узел, в котором пересекаются два списка.
Если два связанных списка вообще не пересекаются, вернуть null.

Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
Output: Intersected at '8'
Объяснение: значение пересекаемого узла равно 8 (обратите внимание, что оно не должно быть равно 0, если два списка
пересекаются).
Из заголовка A это читается как [4,1,8,4,5]. Из заголовка B это читается как [5,6,1,8,4,5]. Перед пересекаемым узлом в A
есть 2 узла; Перед пересекаемым узлом в B есть 3 узла.
- Обратите внимание, что значение пересекаемого узла не равно 1, поскольку узлы со значением 1 в A и B (2-й узел в A и
3-й узел в B) являются разными ссылками на узлы. Другими словами, они указывают на два разных места в памяти, в то время
как узлы со значением 8 в A и B (3-й узел в A и 4-й узел в B) указывают на одно и то же место в памяти.
*/

func main() {
	// Создаем общую часть списков
	common := createList([]int{8, 4, 5})

	// Создаем список A: [4,1] -> common [8,4,5]
	listA := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: common}}

	// Создаем список B: [5,6,1] -> common [8,4,5]
	listB := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: common}}}

	// Находим точку пересечения
	intersection := getIntersectionNode(listA, listB)

	// Выводим результат
	if intersection != nil {
		fmt.Printf("Intersected at '%d'\n", intersection.Val)
	} else {
		fmt.Println("No intersection")
	}

	// Выводим списки для наглядности
	fmt.Print("List A: ")
	printList(listA)
	fmt.Print("List B: ")
	printList(listB)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// getIntersectionNode находит точку пересечения двух связных списков.
// time: O(n+m), где n и m - длины списков A и B, space: O(1)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil { // Если хотя бы один список пуст, то нет пересечения
		return nil
	}

	dummy1, dummy2 := headA, headB // Создаем временные указатели на начало списков

	for dummy1 != dummy2 { // Пока временные указатели не сходятся
		if dummy1 != nil { // Если текущий узел в списке A не равен nil, то переходим к следующему узлу в списке A
			dummy1 = dummy1.Next // Переходим к следующему узлу в списке A
		} else { // Если текущий узел в списке A равен nil, то переходим в начало списка B
			dummy1 = headB // Переходим в начало списка B
		}

		if dummy2 != nil { // Если текущий узел в списке B не равен nil, то переходим к следующему узлу в списке B
			dummy2 = dummy2.Next // Переходим к следующему узлу в списке B
		} else { // Если текущий узел в списке B равен nil, то переходим в начало списка A
			dummy2 = headA // Переходим в начало списка A
		}
	}

	return dummy1 // Возвращаем пересекаемый узел или nil, если нет пересечения
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
