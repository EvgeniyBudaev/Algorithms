package leetcode.java.linkedList.RemoveNthNodeFromEndOfList;

/* 19. Remove Nth Node From End of List
https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/

Дан заголовок связанного списка, удалить n-й узел из конца списка и вернуть его заголовок.

Example 1:
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Example 2:
Input: head = [1], n = 1
Output: []
*/

public class RemoveNthNodeFromEndOfList {
    public static void main(String[] args) {
        ListNode head1 = createList(new int[]{1, 2, 3, 4, 5});
        int n = 2;
        ListNode result1 = removeNthFromEnd(head1, n);
        printList(result1); // [1, 2, 3, 5]

        ListNode head2 = createList(new int[]{1, 2, 3, 4, 5});
        ListNode result2 = removeNthFromEndVariantSlowFast(head2, n);
        printList(result2); // [1, 2, 3, 5]
    }

    static class ListNode {
        int val;
        ListNode next;

        ListNode(int val) {
            this.val = val;
        }
    }

    /**
     * removeNthFromEnd удаляет n-й узел из конца списка и возвращает его заголовок.
     * Паттерн Dummy + вычисление длины списка.
     * time: O(n), где n - количество узлов в списке
     * space: O(1)
     */
    private static ListNode removeNthFromEnd(ListNode head, int n) {
        // Создаем фиктивный узел dummyNode, чтобы избежать специальных случаев
        // с пустым списком или удалением первого узла
        ListNode dummyNode = new ListNode(0);
        dummyNode.next = head;

        // Вычисляем длину списка с учетом dummyNode
        int length = 0;
        ListNode curr = dummyNode;
        while (curr != null) {
            curr = curr.next;
            length++;
        }

        // Доходим до (n-1)-ой вершины с конца
        curr = dummyNode;
        // Условие i < length-n-1, потому что нам нужно дойти до узла ПЕРЕД удаляемым
        for (int i = 0; i < length - n - 1; i++) {
            curr = curr.next;
        }

        // Удаляем вершину
        // Проверяем, что узел существует, прежде чем удалить его
        if (curr.next != null) {
            curr.next = curr.next.next;
        }

        // Возвращаем заголовок списка, начиная с dummyNode
        return dummyNode.next;
    }

    /**
     * removeNthFromEndVariantSlowFast удаляет n-й узел из конца списка.
     * Паттерн Two Pointers (slow/fast) + Dummy.
     * time: O(n), где n - количество узлов в списке
     * space: O(1)
     */
    private static ListNode removeNthFromEndVariantSlowFast(ListNode head, int n) {
        // Создаем фиктивный узел dummyNode для обработки краевых случаев
        ListNode dummyNode = new ListNode(0);
        dummyNode.next = head;

        ListNode fast = dummyNode;
        // Перемещаем быстрый указатель на n+1 позиций вперед
        for (int i = 0; i <= n; i++) {
            fast = fast.next;
        }

        ListNode slow = dummyNode;
        // Перемещаем оба указателя, пока быстрый не достигнет конца
        while (fast != null) {
            slow = slow.next;
            fast = fast.next;
        }

        // Удаляем n-й узел с конца
        if (slow.next != null) {
            slow.next = slow.next.next;
        }

        // Возвращаем заголовок списка
        return dummyNode.next;
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
