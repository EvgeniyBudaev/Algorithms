package leetcode.java.linkedList.DeleteNodeInALinkedList;

import java.lang.classfile.components.ClassPrinter.ListNode;

/* 237. Delete Node in a Linked List
https://leetcode.com/problems/delete-node-in-a-linked-list/description/

Есть односвязный заголовок списка, и мы хотим удалить узел node в нем.
Вам дан узел, который нужно удалить node. Вам не будет предоставлен доступ к первому узлу заголовка.
Все значения связанного списка уникальны, и гарантируется, что данный узел node не является последним узлом в связанном списке.
Удалить данный узел. Обратите внимание, что под удалением узла мы не подразумеваем его удаление из памяти. Мы имеем в виду:

Значение данного узла не должно существовать в связанном списке.
Количество узлов в связанном списке должно уменьшиться на единицу.
Все значения до узла должны быть в том же порядке.
Все значения после узла должны быть в том же порядке.

Input: head = [4,5,1,9], node = 5
Output: [4,1,9]
Пояснение: Вам дан второй узел со значением 5, связанный список должен стать 4 -> 1 -> 9 после вызова вашей функции.
*/

public class DeleteNodeInALinkedList {
    // Объединение K отсортированных связных списков
    public static void main(String[] args) {
        ListNode list = createList(new int[]{4, 5, 1, 9});
        ListNode nodeToDelete = findNode(list, 5); // Находим узел со значением 5 для удаления
        deleteNode(nodeToDelete);
        printList(list); // [4,1,9]
    }

    static class ListNode {
        int val;
        ListNode next;

        ListNode(int val) {
            this.val = val;
        }
    }

    // deleteNode удаляет узел с указанным значением из связного списка.
    // time: O(1), space: O(1)
    private static void deleteNode(ListNode node) {
        if (node == null || node.next == null) {
            return; // Нельзя удалить последний узел или null
        }

        // Копируем значение следующего узла в текущий
        node.val = node.next.val;

        // Пропускаем следующий узел в списке
        node.next = node.next.next;
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

    // Вспомогательный метод для поиска узла по значению (для демонстрации)
    private static ListNode findNode(ListNode head, int val) {
        while (head != null) {
            if (head.val == val) {
                return head;
            }
            head = head.next;
        }
        return null;
    }
}
