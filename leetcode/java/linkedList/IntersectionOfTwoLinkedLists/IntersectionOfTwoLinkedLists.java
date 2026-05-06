package leetcode.java.linkedList.IntersectionOfTwoLinkedLists;

import java.lang.classfile.components.ClassPrinter.ListNode;

/* 160. Intersection of Two Linked Lists
https://leetcode.com/problems/intersection-of-two-linked-lists/description/

Даны заголовки двух односвязных списков headA и headB, вернуть узел, в котором пересекаются два списка.
Если два связанных списка вообще не пересекаются, вернуть null.

Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
Output: Intersected at '8'
Объяснение: значение пересекаемого узла равно 8 (обратите внимание, что оно не должно быть равно 0, если два списка
пересекаются).
Из заголовка A это читается как [4,1,8,4,5]. Из заголовка B это читается как [5,6,1,8,4,5]. Перед пересекаемым узлом в A
есть 2 узла; Перед пересекаемым узлом в B есть 3 узла.
- Обратите внимание, что значение пересекаемого узла не равно 1, поскольку узлы со значением 1 в A и B (2-й узел в A и
3-й узел в B) являются разными ссылками на узлы. Другими словами, они указывают на два разных места в памяти, в то время
как узлы со значением 8 в A и B (3-й узел в A и 4-й узел в B) указывают на одно и то же место в памяти.
*/

public class IntersectionOfTwoLinkedLists {
    // Поиск точки пересечения двух связных списков
    public static void main(String[] args) {
        // Создаем общую часть списков
        ListNode common = createList(new int[]{8, 4, 5});

        // Создаем список A: [4,1] -> common [8,4,5]
        ListNode listA = new ListNode(4);
        listA.next = new ListNode(1);
        listA.next.next = common;

        // Создаем список B: [5,6,1] -> common [8,4,5]
        ListNode listB = new ListNode(5);
        listB.next = new ListNode(6);
        listB.next.next = new ListNode(1);
        listB.next.next.next = common;

        // Находим точку пересечения
        ListNode intersection = getIntersectionNode(listA, listB);

        // Выводим результат
        if (intersection != null) {
            System.out.printf("Intersected at '%d'%n", intersection.val);
        } else {
            System.out.println("No intersection");
        }

        // Выводим списки для наглядности
        System.out.print("List A: ");
        printList(listA);
        System.out.print("List B: ");
        printList(listB);
    }

    static class ListNode {
        int val;
        ListNode next;

        ListNode(int val) {
            this.val = val;
        }
    }

    // getIntersectionNode находит точку пересечения двух связных списков.
    // time: O(n+m), где n и m - длины списков A и B, space: O(1)
    private static ListNode getIntersectionNode(ListNode headA, ListNode headB) {
        if (headA == null || headB == null) { // Если хотя бы один список пуст
            return null;
        }

        ListNode pointer1 = headA; // Указатель для списка A
        ListNode pointer2 = headB; // Указатель для списка B

        // Двигаем указатели, пока они не встретятся
        // Когда указатель достигает конца списка, он переходит к голове другого списка
        while (pointer1 != pointer2) {
            pointer1 = (pointer1 != null) ? pointer1.next : headB;
            pointer2 = (pointer2 != null) ? pointer2.next : headA;
        }

        // Возвращаем точку пересечения или null, если пересечения нет
        return pointer1;
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
