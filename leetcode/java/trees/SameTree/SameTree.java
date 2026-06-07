package leetcode.java.trees.SameTree;

import java.util.LinkedList;
import java.util.Queue;

/* 100. Same Tree
https://leetcode.com/problems/same-tree/description/

Имея корни двух бинарных деревьев p и q, напишите функцию для проверки их идентичности.
Два бинарных дерева считаются одинаковыми, если они структурно идентичны, а узлы имеют одинаковое значение.

Example 1:
Input: p = [1,2,3], q = [1,2,3]
Output: true

Example 2:
Input: p = [1,2], q = [1,null,2]
Output: false
*/

public class SameTree {
    public static void main(String[] args) {
        TreeNode p = createTree(new Integer[]{1, 2, 3});
        TreeNode q = createTree(new Integer[]{1, 2, 3});
        System.out.println(isSameTree(p, q)); // true
    }

    static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode() {
        }

        TreeNode(int val) {
            this.val = val;
        }

        TreeNode(int val, TreeNode left, TreeNode right) {
            this.val = val;
            this.left = left;
            this.right = right;
        }
    }

    // isSameTree проверяет, являются ли два дерева идентичными.
    // time: O(n), где n - количество узлов в дереве
    // space: O(h), где h - высота дерева. В худшем случае h = n.
    private static boolean isSameTree(TreeNode p, TreeNode q) {
        // Если одно из деревьев null, проверяем оба
        if (p == null || q == null) {
            return p == null && q == null;
        }
        // Проверяем значения текущих узлов
        if (p.val != q.val) {
            return false;
        }
        // Рекурсивно проверяем левые и правые поддеревья
        return isSameTree(p.left, q.left) && isSameTree(p.right, q.right);
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
                node.left = new TreeNode(values[i]); // Создаем левого потомка и добавляем его в очередь
                queue.offer(node.left);
            }
            i++;

            // Правый потомок
            if (i < values.length && values[i] != null) {
                node.right = new TreeNode(values[i]); // Создаем правого потомка и добавляем его в очередь
                queue.offer(node.right);
            }
            i++;
        }

        // Возвращаем ссылку на корневой узел
        return root;
    }
}
