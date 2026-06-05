package leetcode.java.trees.PopulatingNextRightPointersInEachNode;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

/* 116. Populating Next Right Pointers in Each Node
https://leetcode.com/problems/populating-next-right-pointers-in-each-node/description/

Вам дано идеальное двоичное дерево, в котором все листья находятся на одном уровне, а у каждого родителя есть два
потомка. Двоичное дерево имеет следующее определение:

struct Node {
int val;
Node *left;
Node *right;
Node *next;
}

Заполните каждый указатель next, чтобы он указывал на его следующий правый узел. Если следующего правого узла нет,
указатель next должен быть установлен в NULL.
Изначально все указатели next установлены в NULL.

Input: root = [1,2,3,4,5,6,7]
Output: [1,#,2,3,#,4,5,6,7,#]
Пояснение: Учитывая приведенное выше идеальное двоичное дерево (рисунок A), ваша функция должна заполнить каждый
следующий указатель так, чтобы он указывал на его следующий правый узел, как на рисунке B. Сериализованный вывод
располагается в порядке уровней, соединенных указателями next, при этом «#» обозначает конец каждого уровня.
*/

public class PopulatingNextRightPointersInEachNode {
    public static void main(String[] args) {
        Node root = createTree(new Integer[]{1, 2, 3, 4, 5, 6, 7});
        Node connected = connect(root);
        System.out.println(printTreeWithNext(connected)); // [1,#,2,3,#,4,5,6,7,#]
    }

    static class Node {
        public int val;
        public Node left;
        public Node right;
        public Node next;

        public Node() {
        }

        public Node(int _val) {
            val = _val;
        }

        public Node(int _val, Node _left, Node _right, Node _next) {
            val = _val;
            left = _left;
            right = _right;
            next = _next;
        }
    }

    /**
     * connect связывает указатели next узлов в идеальном двоичном дереве.
     * time: O(n), space: O(1)
     */
    private static Node connect(Node root) {
        if (root == null) { // Если дерево пустое, возвращаем его
            return null;
        }

        // Начинаем с корня
        Node leftmost = root;

        // Пока есть следующий уровень
        while (leftmost.left != null) {
            Node head = leftmost;

            // Проходим по текущему уровню и связываем узлы следующего уровня
            while (head != null) {
                // Связываем левого потомка с правым
                head.left.next = head.right;

                // Связываем правого потомка с левым потомком следующего узла (если есть)
                if (head.next != null) {
                    head.right.next = head.next.left;
                }

                // Переходим к следующему узлу в текущем уровне
                head = head.next;
            }

            // Переходим на следующий уровень
            leftmost = leftmost.left;
        }

        return root; // Возвращаем корневой узел
    }

    /**
     * createTree создает идеальное бинарное дерево из массива значений.
     */
    private static Node createTree(Integer[] values) {
        if (values == null || values.length == 0 || values[0] == null) {
            return null;
        }

        Node root = new Node(values[0]);
        Queue<Node> queue = new LinkedList<>();
        queue.add(root);
        int i = 1;

        while (!queue.isEmpty() && i < values.length) {
            Node node = queue.poll();

            // Левый потомок
            if (i < values.length && values[i] != null) {
                node.left = new Node(values[i]);
                queue.add(node.left);
            }
            i++;

            // Правый потомок
            if (i < values.length && values[i] != null) {
                node.right = new Node(values[i]);
                queue.add(node.right);
            }
            i++;
        }

        return root;
    }

    /**
     * printTreeWithNext выводит дерево в формате [val,#,...]
     */
    private static String printTreeWithNext(Node root) {
        if (root == null) {
            return "[]";
        }

        List<String> result = new ArrayList<>();
        Queue<Node> queue = new LinkedList<>();
        queue.add(root);

        while (!queue.isEmpty()) {
            int levelSize = queue.size();
            for (int i = 0; i < levelSize; i++) {
                Node node = queue.poll();

                // Добавляем значение текущего узла
                result.add(String.valueOf(node.val));

                // Если это последний узел на уровне (next == null), добавляем "#"
                // (Исправлено дублирование значений next из оригинального Go-кода)
                if (node.next == null) {
                    result.add("#");
                }

                // Добавляем потомков в очередь
                if (node.left != null) {
                    queue.add(node.left);
                }
                if (node.right != null) {
                    queue.add(node.right);
                }
            }
        }

        return "[" + String.join(",", result) + "]";
    }
}
