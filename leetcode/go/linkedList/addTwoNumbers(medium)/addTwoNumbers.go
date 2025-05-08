package main

import "fmt"

/* 2. Add Two Numbers
https://leetcode.com/problems/add-two-numbers/description/

Вам даны два непустых связанных списка, представляющих два неотрицательных целых числа. Цифры хранятся в обратном
порядке, и каждый из их узлов содержит одну цифру. Сложите два числа и верните сумму в виде связанного списка.

Вы можете предположить, что два числа не содержат начальных нулей, кроме самого числа 0.

Example 1:
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
*/

func main() {
	l1 := createList([]int{2, 4, 3})
	l2 := createList([]int{5, 6, 4})
	fmt.Println(addTwoNumbers(l1, l2)) // [7,0,8]
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers добавляет два связанных списка и возвращает результат в виде связанного списка.
// time: O(n), space: O(1)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // Для хранения результата
	curr := dummy        // Текущая нода, которая собирает ответ
	carry := 0           // Остаток который будем переносить

	// Цикл по спискам
	for l1 != nil || l2 != nil || carry != 0 { // Пока есть что добавлять
		// Получаем текущие значения или 0, если список закончился
		l1Val := 0 // Значение первого списка
		// Если список не пуст, получаем его значение
		if l1 != nil {
			l1Val = l1.Val // Значение первого списка
			l1 = l1.Next   // Переходим к следующей ноде
		}

		l2Val := 0 // Значение второго списка
		// Если список не пуст, получаем его значение
		if l2 != nil {
			l2Val = l2.Val // Значение второго списка
			l2 = l2.Next   // Переходим к следующей ноде
		}

		// Вычисляем сумму и перенос
		sum := l1Val + l2Val + carry        // Сумма
		newNode := &ListNode{Val: sum % 10} // Новая нода
		carry = sum / 10                    // Перенос

		// Добавляем новую ноду в результат
		curr.Next = newNode // Добавляем новую ноду в результат
		curr = curr.Next    // Переходим к следующей ноде
	}

	// Возвращаем результат
	return dummy.Next
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
