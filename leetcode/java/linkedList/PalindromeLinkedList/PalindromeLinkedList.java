package leetcode.java.linkedList.PalindromeLinkedList;

/* 234. Palindrome Linked List
https://leetcode.com/problems/palindrome-linked-list/

Если задан заголовок односвязного списка, вернуть значение true, если он является палиндромом, или false в противном случае.

Example 1:
Input: head = [1,2,2,1]
Output: true

Example 2:
Input: head = [1,2]
Output: false
*/

public class PalindromeLinkedList {
    public static void main(String[] args) {
        ListNode head = createList(new int[]{1, 2, 2, 1});
        System.out.println(isPalindrome(head)); // true

        ListNode head2 = createList(new int[]{1, 2});
        System.out.println(isPalindrome(head2)); // false 
    }

    static class ListNode {
        int val;
        ListNode next;

        ListNode(int val) {
            this.val = val;
        }
    }

    /**
     * isPalindrome возвращает true, если связный список является палиндромом, иначе - false.
     * time: O(n), где n - количество узлов в списке
     * space: O(1)
     */
    private static boolean isPalindrome(ListNode head) {
        // Если список пуст или содержит только один узел, он является палиндромом
        if (head == null || head.next == null) {
            return true;
        }

        // Находим конец первой половины
        ListNode firstHalfEnd = findMiddleNode(head);
        // Разворачиваем вторую половину
        ListNode secondHalfStart = reverseList(firstHalfEnd);

        // Сравниваем две половины
        ListNode p1 = head;               // Указатель на первый узел первой половины
        ListNode p2 = secondHalfStart;    // Указатель на первый узел второй половины

        while (p1 != null && p2 != null) {
            // Сравниваем значения текущих узлов первой и второй половин
            if (p1.val != p2.val) {
                return false;
            }
            p1 = p1.next;
            p2 = p2.next;
        }

        // Восстанавливаем оригинальный список (опционально)
        // firstHalfEnd.next = reverseList(secondHalfStart);

        return true;
    }

    /**
     * reverseList переворачивает связный список и возвращает перевернутый список.
     */
    private static ListNode reverseList(ListNode head) {
        ListNode prev = null;  // Переменная для хранения предыдущего узла
        ListNode curr = head;  // Текущий узел, начинаем с головы списка

        while (curr != null) {
            ListNode tmp = curr;        // Сохраняем текущий узел в tmp
            curr = curr.next;           // Перемещаем текущий узел на следующий узел
            tmp.next = prev;            // Переворачиваем указатель следующего узла
            prev = tmp;                 // Перемещаем предыдущий узел на текущий узел
        }

        // Возвращаем новый заголовок списка
        return prev;
    }

    /**
     * findMiddleNode возвращает средний узел связанного списка.
     * Использует алгоритм "черепаха и заяц" (slow/fast pointers).
     */
    private static ListNode findMiddleNode(ListNode head) {
        ListNode slow = head;  // Медленный указатель
        ListNode fast = head;  // Быстрый указатель

        // Проходим список, пока быстрый указатель не достигнет конца
        while (fast != null && fast.next != null) {
            slow = slow.next;        // Перемещаем медленный указатель на один шаг
            fast = fast.next.next;   // Перемещаем быстрый указатель на два шага
        }

        // Возвращаем средний узел
        return slow;
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
}
