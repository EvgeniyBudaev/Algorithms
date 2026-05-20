package leetcode.java.linkedList.ReverseLinkedList;

import java.lang.classfile.components.ClassPrinter.ListNode;

/* 206. Reverse Linked List
https://leetcode.com/problems/reverse-linked-list/description/

Дан заголовок односвязного списка, перевернуть список и вернуть перевернутый список.

Example 1:
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

Example 2:
Input: head = [1,2]
Output: [2,1]

Example 3:
Input: head = []
Output: []
*/

public class ReverseLinkedList {
    public static void main(String[] args) {
        ListNode head = createList(new int[]{1, 2, 3, 4, 5});
        ListNode result = reverseList(head);
        printList(result); // [5, 4, 3, 2, 1]
    }

    static class ListNode {
        int val;
        ListNode next;

        ListNode(int val) {
            this.val = val;
        }
    }

    /**
     * reverseList переворачивает связный список и возвращает перевернутый список.
     * Итеративный подход с тремя указателями.
     * <p>
     * time: O(n), где n - количество узлов в списке
     * space: O(1)
     */
    private static ListNode reverseList(ListNode head) {
        ListNode prev = null;  // Предыдущий узел (изначально null)
        ListNode curr = head;  // Текущий узел, начинаем с головы списка

        while (curr != null) {
            ListNode tmp = curr;        // Сохраняем текущий узел во временной переменной
            curr = curr.next;           // Перемещаем curr на следующий узел
            tmp.next = prev;            // Переворачиваем ссылку: текущий указывает на предыдущий
            prev = tmp;                 // Перемещаем prev на текущий узел
        }

        // Возвращаем новый заголовок списка (бывший последний узел)
        return prev;
    }

    /**
     * createList создает связный список из массива значений.
     */
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

    /**
     * printList печатает список в формате [v1,v2,v3]
     */
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
