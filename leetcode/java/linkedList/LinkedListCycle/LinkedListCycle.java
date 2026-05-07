package leetcode.java.linkedList.LinkedListCycle;

/* 141. Linked List Cycle
https://leetcode.com/problems/linked-list-cycle/description/

Если задан head, head связанного списка, определите, есть ли в связанном списке цикл.

Цикл в связанном списке есть, если в списке есть некоторый узел, к которому можно снова прийти, непрерывно следуя
указателю next. Внутри pos используется для обозначения индекса узла, к которому подключен указатель next в tail.
Обратите внимание, что pos не передается как параметр.

Верните true, если в связанном списке есть цикл. В противном случае верните false.0.

Input: head = [3,2,0,-4], pos = 1
Output: true
Пояснение: В связанном списке есть цикл, где хвост соединяется с 1-м узлом (индексированным как 0).
*/

// Обнаружение цикла в связном списке (алгоритм Флойда)
public class LinkedListCycle {
    public static void main(String[] args) {
        // Создаем список: 3 → 2 → 0 → -4
        ListNode head = createList(new int[]{3, 2, 0, -4});

        // Создаем цикл: последний узел (-4) → узел с индексом 1 (значение 2)
        // Аналогично поведению задачи LeetCode #141
        if (head != null && head.next != null) {
            ListNode tail = head;
            ListNode cycleEntry = head.next; // узел со значением 2
            while (tail.next != null) {
                tail = tail.next;
            }
            tail.next = cycleEntry; // замыкаем цикл
        }

        System.out.println(hasCycle(head)); // true
    }

    static class ListNode {
        int val;
        ListNode next;

        ListNode(int val) {
            this.val = val;
        }
    }

    // hasCycle определяет, есть ли цикл в связанном списке.
    // time: O(n), space: O(1) — алгоритм «черепаха и заяц» (Флойда)
    private static boolean hasCycle(ListNode head) {
        if (head == null) {
            return false;
        }

        ListNode slow = head; // Медленный указатель (черепаха)
        ListNode fast = head; // Быстрый указатель (заяц)

        while (fast != null && fast.next != null) {
            slow = slow.next;           // +1 шаг
            fast = fast.next.next;      // +2 шага

            if (slow == fast) {         // Встреча указателей = цикл
                return true;
            }
        }

        return false; // Быстрый указатель достиг конца — цикла нет
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
}
