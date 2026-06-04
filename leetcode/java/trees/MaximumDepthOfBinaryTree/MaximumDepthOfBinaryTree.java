package leetcode.java.trees.MaximumDepthOfBinaryTree;

import java.util.LinkedList;
import java.util.Queue;

/* 104. Maximum Depth of Binary Tree
https://leetcode.com/problems/maximum-depth-of-binary-tree/description/

Учитывая корень бинарного дерева, вернуть его максимальную глубину.
Максимальная глубина бинарного дерева — это количество узлов на самом длинном пути от корневого узла до самого дальнего
листового узла.

Input: root = [3,9,20,null,null,15,7]
Output: 3
*/

public class MaximumDepthOfBinaryTree {
    public static void main(String[] args) {
        TreeNode root = createTree(new Integer[]{3, 9, 20, null, null, 15, 7});
        System.out.println(maxDepth(root)); // 3
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
     * maxDepth возвращает максимальную глубину дерева.
     * time: O(n), где n - количество узлов в дереве.
     * space: O(h), где h - высота дерева (из-за стека вызовов рекурсии).
     */
    private static int maxDepth(TreeNode root) {
        if (root == null) { // Если дерево пустое, то возвращаем 0
            return 0;
        }

        int leftDepth = maxDepth(root.left);   // Рекурсивно вычисляем максимальную глубину левого поддерева
        int rightDepth = maxDepth(root.right); // Рекурсивно вычисляем максимальную глубину правого поддерева

        return 1 + Math.max(leftDepth, rightDepth); // Возвращаем максимальную глубину дерева, увеличенную на 1 (текущий уровень)
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
            TreeNode node = queue.poll(); // Берем первый элемент из очереди и удаляем его

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
