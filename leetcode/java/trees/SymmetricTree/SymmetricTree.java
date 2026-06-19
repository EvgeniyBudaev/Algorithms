package leetcode.java.trees.SymmetricTree;

import java.util.LinkedList;
import java.util.Queue;

/* 101. Symmetric Tree
https://leetcode.com/problems/symmetric-tree/description/

Дан корень двоичного дерева, проверьте, является ли он зеркальным отражением самого себя
(т.е. симметричен ли относительно своего центра).

Example 1:
Input: root = [1,2,2,3,4,4,3]
Output: true

Example 2:
Input: root = [1,2,2,null,3,null,3]
Output: false
*/

public class SymmetricTree {
    // Проверка симметричности бинарного дерева
    public static void main(String[] args) {
        TreeNode root1 = createTree(new Integer[]{1, 2, 2, 3, 4, 4, 3});
        System.out.println(isSymmetric(root1)); // true

        TreeNode root2 = createTree(new Integer[]{1, 2, 2, null, 3, null, 3});
        System.out.println(isSymmetric(root2)); // false  
    }

    // Определение узла бинарного дерева
    static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode(int x) {
            val = x;
        }
    }

    // isSymmetric проверяет, является ли двоичное дерево симметричным относительно своего центра.
    // time: O(n), где n - количество узлов в дереве, space: O(h), где h - высота дерева.
    private static boolean isSymmetric(TreeNode root) {
        // Проверяем, является ли дерево пустым
        if (root == null) {
            return true;
        }
        // Проверяем симметричность двух поддеревьев
        return check(root.left, root.right);
    }

    // check рекурсивно проверяет симметричность двух поддеревьев.
    private static boolean check(TreeNode left, TreeNode right) {
        // Проверяем, являются ли оба поддерева пустыми
        if (left == null || right == null) {
            return left == null && right == null;
        }
        // Проверяем, равны ли значения корневых узлов
        if (left.val != right.val) {
            return false;
        }
        // Рекурсивно проверяем левый потомок левого поддерева с правым потомком правого,
        // и правый потомок левого с левым потомком правого
        return check(left.left, right.right) && check(left.right, right.left);
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
                node.left = new TreeNode(values[i]); // Создаем левого потомка
                queue.offer(node.left);              // Добавляем левого потомка в очередь
            }
            i++;

            // Правый потомок
            if (i < values.length && values[i] != null) {
                node.right = new TreeNode(values[i]); // Создаем правого потомка
                queue.offer(node.right);              // Добавляем правого потомка в очередь
            }
            i++;
        }

        // Возвращаем ссылку на корневой узел
        return root;
    }
}
