package leetcode.java.linkedList.MiddleOfTheLinkedList;

import java.lang.classfile.components.ClassPrinter.ListNode;

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

// Поиск среднего узла связного списка (алгоритм двух указателей)
public class MiddleOfTheLinkedList {
    public static void main(String[] args) {
        ListNode head = createList(new int[]{1, 2, 3, 4, 5});
        ListNode result = middleNode(head);
        printList(result); // [3,4,5]
    }

    static class ListNode {
        int val;
        ListNode next;
        
        ListNode(int val) {
            this.val = val;
        }
    }

    // middleNode возвращает средний узел связного списка.
    // Если чётное количество узлов — возвращается второй из двух средних.
    // time: O(n), где n - количество узлов в списке
    // space: O(1)
    private static ListNode middleNode(ListNode head) {
        if (head == null) {
            return null;
        }
        
        ListNode slow = head; // Медленный указатель: +1 шаг
        ListNode fast = head; // Быстрый указатель: +2 шага

        // Двигаем указатели, пока быстрый не достигнет конца
        while (fast != null && fast.next != null) {
            slow = slow.next;           // +1 шаг
            fast = fast.next.next;      // +2 шага
        }

        // slow теперь указывает на средний узел
        return slow;
    }

    // createList создает связный список из массива значений
    private static ListNode createList(int[] values) {
        if (values == null || values.length == 0) {
            return null;
        }

        ListNode head = new ListNode(values[0]);
        ListNode current = head;

        for (int i = 1; i < values.length; i++) {
            current.next = new ListNode(values[i]);
            current = current.next;
        }

        return head;
    }

    // printList печатает список в формате [v1,v2,v3]
    private static void printList(ListNode head) {
        System.out.print("[");
        while (head != null) {
            System.out.print(head.val);
            if (head.next != null) {
                System.out.print(",");
            }
            head = head.next;
        }
        System.out.println("]");
    }
}
