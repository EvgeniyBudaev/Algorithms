package leetcode.java.linkedList.ReorderList;

import java.lang.classfile.components.ClassPrinter.ListNode;

/* 143. Reorder List
https://leetcode.com/problems/reorder-list/

Вам дан заголовок односвязного списка. Список можно представить как:

L0 → L1 → … → Ln - 1 → Ln
Измените порядок списка, чтобы он имел следующий вид:

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
Вы не можете изменять значения в узлах списка. Изменять можно только сами узлы.

Example 1:
Input: head = [1,2,3,4]
Output: [1,4,2,3]

Example 2:
Input: head = [1,2,3,4,5]
Output: [1,5,2,4,3]
*/

public class ReorderList {
    public static void main(String[] args) {
        ListNode head = createList(new int[]{1, 2, 3, 4});
        reorderList(head); // [1, 4, 2, 3]
    }

    static class ListNode {
        int val;
        ListNode next;

        ListNode(int val) {
            this.val = val;
        }
    }

    /**
     * reorderList меняет порядок узлов списка на:
     * L0 → Ln → L1 → Ln-1 → L2 → Ln-2 → …
     * time: O(n), где n - количество узлов в списке
     * space: O(1)
     */
    private static void reorderList(ListNode head) {
        if (head == null || head.next == null) {
            return;
        }

        // Находим узел перед серединой
        ListNode preMiddle = preMiddleNode(head);
        // Разворачиваем вторую половину
        ListNode secondHalf = reverseList(preMiddle.next);

        // Разделяем список на две части
        preMiddle.next = null;

        // Объединяем две половины
        ListNode p1 = head;
        ListNode p2 = secondHalf;
        ListNode newHead = p1;

        while (p2 != null) {
            ListNode p1Next = p1.next;
            p1.next = p2;
            p1 = p2;
            p2 = p1Next;
        }

        printList(newHead);
    }

    /**
     * preMiddleNode находит предыдущий узел среднего узла в связанном списке.
     * Использует алгоритм "черепаха и заяц" (slow/fast pointers).
     * time: O(n), space: O(1)
     */
    private static ListNode preMiddleNode(ListNode head) {
        ListNode slow = head;  // Медленный указатель
        ListNode fast = head;  // Быстрый указатель

        // Проходим список, пока быстрый указатель может сделать два шага
        // Условие fast.next.next != null гарантирует, что slow остановится
        // на узле ПЕРЕД серединой (для чётного и нечётного количества узлов)
        while (fast != null && fast.next != null && fast.next.next != null) {
            slow = slow.next;        // Перемещаем медленный указатель на 1 шаг
            fast = fast.next.next;   // Перемещаем быстрый указатель на 2 шага
        }

        // Возвращаем узел перед серединой
        return slow;
    }

    /**
     * reverseList переворачивает связный список и возвращает перевернутый список.
     * time: O(n), space: O(1)
     */
    private static ListNode reverseList(ListNode head) {
        ListNode prev = null;  // Предыдущий узел
        ListNode curr = head;  // Текущий узел

        while (curr != null) {
            ListNode tmp = curr;        // Сохраняем текущий узел
            curr = curr.next;           // Переходим к следующему
            tmp.next = prev;            // Переворачиваем ссылку
            prev = tmp;                 // Двигаем prev вперёд
        }

        // Возвращаем новый заголовок (бывший последний узел)
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
