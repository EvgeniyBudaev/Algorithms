package leetcode.java.linkedList.MergeTwoSortedLists;

/* 21. Merge Two Sorted Lists
https://leetcode.com/problems/merge-two-sorted-lists/

Вам даны заголовки двух отсортированных связанных списков list1 и list2.
Объедините два списка в один отсортированный список. Список должен быть создан путем сращивания узлов первых двух списков.
Верните заголовок объединенного связанного списка.

Example 1:
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
*/

// Объединение двух отсортированных связных списков
public class MergeTwoSortedLists {
    public static void main(String[] args) {
        ListNode list1 = createList(new int[]{1, 2, 4});
        ListNode list2 = createList(new int[]{1, 3, 4});
        ListNode result = mergeTwoLists(list1, list2);
        printList(result); // [1,1,2,3,4,4]
    }

    static class ListNode {
        int val;
        ListNode next;

        ListNode(int val) {
            this.val = val;
        }
    }

    // mergeTwoLists объединяет два отсортированных связных списка в один отсортированный список.
    // time: O(n + m), где n и m - длины списков
    // space: O(1)
    private static ListNode mergeTwoLists(ListNode list1, ListNode list2) {
        ListNode dummy = new ListNode(0); // Фиктивная начальная нода
        ListNode current = dummy;         // Указатель на текущую ноду

        while (list1 != null || list2 != null) {
            // Получаем значения или "бесконечность", если список закончился
            int val1 = getVal(list1);
            int val2 = getVal(list2);

            // Выбираем меньшее значение и перемещаем указатель
            if (val1 <= val2) {
                current.next = list1;
                list1 = list1.next;
            } else {
                current.next = list2;
                list2 = list2.next;
            }
            // Перемещаем указатель на следующую ноду
            current = current.next;
        }

        return dummy.next; // Пропускаем фиктивную ноду
    }

    // getVal возвращает значение ноды или максимальное int, если нода null
    private static int getVal(ListNode node) {
        if (node == null) {
            return Integer.MAX_VALUE;
        }
        return node.val;
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
