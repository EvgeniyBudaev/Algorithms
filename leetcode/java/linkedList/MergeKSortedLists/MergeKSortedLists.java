package leetcode.java.linkedList.MergeKSortedLists;

import java.lang.classfile.components.ClassPrinter.ListNode;

/* 23. Merge k Sorted Lists
https://leetcode.com/problems/merge-k-sorted-lists/description/

Вам дан массив из k связанных списков, каждый связанный список отсортирован в порядке возрастания.
Объедините все связанные списки в один отсортированный связанный список и верните его.

Example 1:
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
*/

// Объединение K отсортированных связных списков
public class MergeKSortedLists {
    public static void main(String[] args) {
        ListNode list1 = createList(new int[]{1, 4, 5});
        ListNode list2 = createList(new int[]{1, 3, 4});
        ListNode list3 = createList(new int[]{2, 6});

        ListNode[] lists = {list1, list2, list3};
        ListNode result = mergeKLists(lists);

        printList(result); // [1,1,2,3,4,4,5,6] 
    }

    static class ListNode {
        int val;
        ListNode next;

        ListNode(int val) {
            this.val = val;
        }
    }

    // mergeKLists объединяет все связные списки в один отсортированный связный список.
    // time: O(k * n), где k - количество списков, n - общее число узлов
    // space: O(1)
    private static ListNode mergeKLists(ListNode[] lists) {
        if (lists == null || lists.length == 0) {
            return null;
        }

        // Решение основано на "нескольких указателях".
        // Каждый раз находим наименьший элемент, на который указывают списки,
        // и двигаем указатель для этого списка.
        ListNode dummy = new ListNode(0); // Фиктивная начальная нода
        ListNode current = dummy;         // Указатель на текущую ноду

        while (true) {
            // Находим индекс списка с наименьшей нодой
            int minIdx = getMinNodeIdx(lists);
            if (minIdx == -1) {
                break; // Все списки обработаны
            }

            // Добавляем минимальный узел в результат
            current.next = lists[minIdx];
            current = current.next;

            // Перемещаем указатель в списке с минимальным узлом
            lists[minIdx] = lists[minIdx].next;
        }

        return dummy.next; // Пропускаем фиктивную ноду
    }

    // getMinNodeIdx возвращает индекс, в котором текущая нода наименьшая.
    private static int getMinNodeIdx(ListNode[] lists) {
        int minVal = Integer.MAX_VALUE; // Максимальное значение int
        int minIdx = -1;                // Индекс минимального значения

        for (int i = 0; i < lists.length; i++) {
            // Если нода не null и ее значение меньше текущего минимального
            if (lists[i] != null && lists[i].val < minVal) {
                minVal = lists[i].val;
                minIdx = i;
            }
        }

        // Если ни одна нода не была найдена, возвращаем -1
        return minIdx;
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
