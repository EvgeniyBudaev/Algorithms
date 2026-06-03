package leetcode.java.trees.KthSmallestElementInABST;

import java.util.LinkedList;
import java.util.Queue;

/* 230. Kth Smallest Element in a BST
https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/

Учитывая корень двоичного дерева поиска и целое число k, вернуть k-е наименьшее значение (индексированное на 1) среди
всех значений узлов в дереве.

Input: root = [3,1,4,null,2], k = 1
Output: 1
*/

public class KthSmallestElementInABST {
    public static void main(String[] args) {
        TreeNode root = createTree(new Integer[]{3, 1, 4, null, 2});
        System.out.println(kthSmallest(root, 1)); // 1
    }

    static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode(int x) {
            val = x;
        }
    }

    /**
     * kthSmallest возвращает k-е наименьшее значение в двоичном дереве поиска.
     * time: O(n), где n - количество узлов в дереве.
     * space: O(h), где h - высота дерева.
     */
    private static int kthSmallest(TreeNode root, int k) {
        int[] count = new int[1];  // Переменная для подсчета количества обработанных узлов
        int[] result = new int[1]; // Переменная для хранения результата

        inOrder(root, k, count, result); // Вызываем функцию для обхода дерева

        return result[0]; // Возвращаем результат
    }

    /**
     * inOrder выполняет обход дерева в порядке возрастания значений и находит k-е наименьшее значение.
     * time: O(n), где n - количество узлов в дереве.
     * space: O(h), где h - высота дерева.
     */
    private static void inOrder(TreeNode node, int k, int[] count, int[] result) {
        if (node == null) { // Если узел равен null, то выходим из функции
            return;
        }

        // Обходим левое поддерево
        inOrder(node.left, k, count, result);

        // Оптимизация: если k-й элемент уже был найден в левом поддереве, 
        // нет смысла продолжать обход и обрабатывать текущий узел.
        if (count[0] == k) {
            return;
        }

        // Обрабатываем текущий узел
        count[0]++; // Увеличиваем счетчик
        if (count[0] == k) { // Если счетчик равен k, то записываем значение текущего узла в результат и выходим
            result[0] = node.val; // Записываем значение текущего узла в результат
            return;
        }

        // Обходим правое поддерево
        inOrder(node.right, k, count, result);
    }

    /**
     * createTree создает бинарное дерево из массива значений (аналог level-order traversal).
     */
    private static TreeNode createTree(Integer[] values) {
        if (values == null || values.length == 0 || values[0] == null) {
            return null;
        }

        TreeNode root = new TreeNode(values[0]); // Создаем корневой узел с первым значением в массиве
        Queue<TreeNode> queue = new LinkedList<>(); // Создаем очередь для обхода дерева по уровням
        queue.add(root);
        int i = 1; // Индекс для обхода массива значений

        while (!queue.isEmpty() && i < values.length) {
            TreeNode node = queue.poll(); // Берем первый элемент из очереди

            // Левый потомок
            if (i < values.length && values[i] != null) {
                node.left = new TreeNode(values[i]); // Создаем левый потомок
                queue.add(node.left);                // Добавляем левого потомка в очередь
            }
            i++;

            // Правый потомок
            if (i < values.length && values[i] != null) {
                node.right = new TreeNode(values[i]); // Создаем правого потомка
                queue.add(node.right);                // Добавляем правого потомка в очередь
            }
            i++;
        }

        return root; // Возвращаем указатель на корневой узел
    }
}
