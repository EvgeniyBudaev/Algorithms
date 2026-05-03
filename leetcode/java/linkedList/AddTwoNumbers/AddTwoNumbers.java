package leetcode.java.linkedList.AddTwoNumbers;

import java.lang.classfile.components.ClassPrinter.ListNode;
import java.util.ArrayList;
import java.util.List;

/* 2. Add Two Numbers
https://leetcode.com/problems/add-two-numbers/description/

Вам даны два непустых связанных списка, представляющих два неотрицательных целых числа. Цифры хранятся в обратном
порядке, и каждый из их узлов содержит одну цифру. Сложите два числа и верните сумму в виде связанного списка.

Вы можете предположить, что два числа не содержат начальных нулей, кроме самого числа 0.

Example 1:
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
*/

public class AddTwoNumbers {
    public static void main(String[] args) {
        ListNode l1 = createList(new int[]{2, 4, 3}); // 342
        ListNode l2 = createList(new int[]{5, 6, 4}); // 465
        ListNode result = addTwoNumbers(l1, l2);      // 807
        System.out.println(listToArray(result));      // [7, 0, 8]
    }

    // ListNode - узел связного списка
    static class ListNode {
        int val;
        ListNode next;

        ListNode(int val) {
            this.val = val;
        }
    }

    // addTwoNumbers добавляет два связанных списка и возвращает результат в виде связанного списка.
    // time: O(max(n, m)), space: O(1) - не учитывая память для результата
    private static ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode dummy = new ListNode(0); // фиктивный узел для удобства
        ListNode curr = dummy;            // текущий узел, который собирает ответ
        int carry = 0;                    // перенос разряда

        // Цикл по спискам: пока есть цифры в l1, l2 или есть перенос
        while (l1 != null || l2 != null || carry != 0) {
            // Получаем текущие значения или 0, если список закончился
            int l1Val = (l1 != null) ? l1.val : 0;
            int l2Val = (l2 != null) ? l2.val : 0;

            // Вычисляем сумму и перенос
            int sum = l1Val + l2Val + carry;
            ListNode newNode = new ListNode(sum % 10); // новая нода с цифрой
            carry = sum / 10;                          // новый перенос

            // Добавляем новую ноду в результат
            curr.next = newNode;
            curr = curr.next;

            // Переходим к следующим узлам входных списков
            if (l1 != null) l1 = l1.next;
            if (l2 != null) l2 = l2.next;
        }

        // Возвращаем результат (пропускаем фиктивный узел)
        return dummy.next;
    }

    // createList создает связный список из массива значений
    public static ListNode createList(int[] values) {
        if (values.length == 0) {
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
    public static void printList(ListNode head) {
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

    // listToArray преобразовывает список в массив
    public static List<Integer> listToArray(ListNode head) {
        List<Integer> result = new ArrayList<>();
        while (head != null) {
            result.add(head.val);
            head = head.next;
        }
        return result;
    }
}
