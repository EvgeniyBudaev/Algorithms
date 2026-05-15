package leetcode.java.linkedList.OddEvenLinkedList;

import java.lang.classfile.components.ClassPrinter.ListNode;

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

// Переупорядочивание связного списка: нечетные и четные узлы
public class OddEvenLinkedList {
    public static void main(String[] args) {
        ListNode head = createList(new int[]{1, 2, 3, 4, 5});
        printList(oddEvenList(head)); // [1,3,5,2,4]
    }

    static class ListNode {
        int val;
        ListNode next;

        ListNode(int val) {
            this.val = val;
        }
    }

    // oddEvenList переупорядочивает связный список, разделяя его на нечетные и четные узлы.
    // time: O(n), space: O(1)
    private static ListNode oddEvenList(ListNode head) {
        if (head == null) { // Если список пустой
            return null;
        }

        ListNode dummyEven = new ListNode(0); // Заглушка для четных узлов
        dummyEven.next = head.next;
        ListNode odd = head;           // Указатель на нечетные узлы
        ListNode even = head.next;     // Указатель на четные узлы

        while (even != null && even.next != null) {
            odd.next = even.next;      // Пропускаем четный узел
            odd = odd.next;            // Перемещаем указатель на следующий нечетный
            even.next = odd.next;      // Пропускаем нечетный узел
            even = even.next;          // Перемещаем указатель на следующий четный
        }

        odd.next = dummyEven.next; // Соединяем нечетную и четную части
        return head;
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
