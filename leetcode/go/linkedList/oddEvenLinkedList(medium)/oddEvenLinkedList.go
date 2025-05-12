package main

import "fmt"

/* 328. Odd Even Linked List
https://leetcode.com/problems/odd-even-linked-list/description/

Учитывая заголовок односвязного списка, сгруппируйте все узлы с нечетными индексами вместе, за которыми следуют узлы с
четными индексами, и верните переупорядоченный список.
Первый узел считается нечетным, а второй — четным и т. д.
Обратите внимание, что относительный порядок внутри как четных, так и нечетных групп должен оставаться таким же, как и
во входных данных.
Вам необходимо решить задачу за O(1) дополнительной пространственной сложности и O(n) временной сложности.

Input: head = [1,2,3,4,5]
Output: [1,3,5,2,4]
*/

func main() {
	head := createList([]int{1, 2, 3, 4, 5})
	printList(oddEvenList(head)) // [1,3,5,2,4]
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// oddEvenList переупорядочивает связный список, разделяя его на нечетные и четные узлы.
// time: O(n), space: O(1)
func oddEvenList(head *ListNode) *ListNode {
	if head == nil { // Если список пустой, то возвращаем nil
		return nil
	}

	dummyEven := &ListNode{Next: head.Next} // Создаем заглушку для четных узлов
	odd, even := head, head.Next            // Создаем указатели на первый нечетный и первый четный узлы

	for even != nil && even.Next != nil { // Проходим по списку, пока есть хотя бы два следующих узла
		odd.Next = even.Next // Связываем нечетный узел с четным узлом
		odd = odd.Next       // Перемещаем указатель на следующий нечетный узел
		even.Next = odd.Next // Связываем четный узел с четным узлом
		even = even.Next     // Перемещаем указатель на следующий четный узел
	}

	odd.Next = dummyEven.Next // Связываем последний нечетный узел с первым четным узлом
	return head               // Возвращаем начальный узел списка
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
