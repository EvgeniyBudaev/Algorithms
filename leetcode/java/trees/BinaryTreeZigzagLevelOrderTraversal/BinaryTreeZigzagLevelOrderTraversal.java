package leetcode.java.trees.BinaryTreeZigzagLevelOrderTraversal;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

/* 103. Binary Tree Zigzag Level Order Traversal
https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/description/

Учитывая корень двоичного дерева, вернуть зигзагообразный порядок обхода значений его узлов (т. е. слева направо, затем
справа налево для следующего уровня и попеременно).

Input: root = [3,9,20,null,null,15,7]
Output: [[3],[20,9],[15,7]]
*/

public class BinaryTreeZigzagLevelOrderTraversal {
    public static void main(String[] args) {
        TreeNode root = createTree(new Integer[]{3, 9, 20, null, null, 15, 7});
        System.out.println(zigzagLevelOrder(root)); // [[3], [20, 9], [15, 7]]
    }

    static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode(int val) {
            this.val = val;
        }
    }

    // zigzagLevelOrder выполняет зигзагообразный обход дерева и возвращает результат.
    // time: O(n), где n - количество узлов в дереве
    // space: O(n), где n - количество узлов в дереве
    private static List<List<Integer>> zigzagLevelOrder(TreeNode root) {
        List<List<Integer>> result = new ArrayList<>(); // Инициализируем пустой результат
        if (root == null) {    // Если корень равен null, возвращаем пустой результат
            return result;
        }

        Queue<TreeNode> queue = new LinkedList<>(); // Создаем очередь и добавляем корень
        queue.offer(root);

        while (!queue.isEmpty()) {       // Пока очередь не пуста
            int levelSize = queue.size();          // Получаем количество узлов на текущем уровне
            LinkedList<Integer> levelValues = new LinkedList<>(); // Создаем список для значений на текущем уровне

            for (int i = 0; i < levelSize; i++) { // Обрабатываем все узлы на текущем уровне
                TreeNode node = queue.poll();   // Берем первый элемент из очереди

                if (node.left != null) { // Если узел имеет левого потомка, добавляем его в очередь
                    queue.offer(node.left); // Добавляем левого потомка в очередь
                }
                if (node.right != null) { // Если узел имеет правого потомка, добавляем его в очередь
                    queue.offer(node.right); // Добавляем правого потомка в очередь
                }

                if (result.size() % 2 == 0) { // Если текущий уровень четный, добавляем значения слева направо
                    levelValues.addLast(node.val); // Добавляем значение в конец списка
                } else { // Если текущий уровень нечетный, добавляем значения справа налево
                    levelValues.addFirst(node.val); // Добавляем значение в начало списка
                }
            }
            result.add(new ArrayList<>(levelValues)); // Добавляем список значений текущего уровня в результат
        }

        return result; // Возвращаем результат
    }

    // createTree создает бинарное дерево из массива значений.
    private static TreeNode createTree(Integer[] values) {
        if (values == null || values.length == 0 || values[0] == null) {
            return null;
        }

        TreeNode root = new TreeNode(values[0]); // Создаем корневой узел с первым значением в массиве
        Queue<TreeNode> queue = new LinkedList<>(); // Создаем очередь для обхода дерева по уровням
        queue.offer(root);
        int i = 1; // Индекс для обхода массива значений

        while (!queue.isEmpty() && i < values.length) {
            TreeNode node = queue.poll(); // Берем первый элемент из очереди

            // Левый потомок
            if (i < values.length && values[i] != null) {
                node.left = new TreeNode(values[i]); // Создаем левый потомок и добавляем его в очередь
                queue.offer(node.left); // Добавляем левого потомка в очередь
            }
            i++;

            // Правый потомок
            // Если правый потомок не равен null, то создаем его и добавляем его в очередь
            if (i < values.length && values[i] != null) {
                node.right = new TreeNode(values[i]); // Создаем правого потомка и добавляем его в очередь
                queue.offer(node.right); // Добавляем правого потомка в очередь
            }
            i++;
        }

        // Возвращаем ссылку на корневой узел
        return root;
    }
}
