package leetcode.java.trees.BinaryTreeInorderTraversal;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

/* 94. Binary Tree Inorder Traversal
https://leetcode.com/problems/binary-tree-inorder-traversal/description/

Для заданного корня бинарного дерева вернуть симметричный обход значений его узлов.

Example 1:
Input: root = [1,null,2,3]
Output: [1,3,2]

Example 2:
Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]
Output: [4,2,6,5,7,1,3,9,8]
*/

public class BinaryTreeInorderTraversal {
    public static void main(String[] args) {
        TreeNode root1 = createTree(new Integer[]{1, null, 2, 3});
        System.out.println(inorderTraversal(root1)); // [1, 3, 2]

        TreeNode root2 = createTree(new Integer[]{1, 2, 3, 4, 5, null, 8, null, null, 6, 7, 9});
        System.out.println(inorderTraversal(root2)); // [4, 2, 6, 5, 7, 1, 3, 9, 8]
    }

    static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode(int val) {
            this.val = val;
        }
    }

    // inorderTraversal возвращает прямой обход значений узлов заданного дерева.
    // time: O(n), где n - количество узлов в дереве
    // space: O(h), где h - высота дерева
    private static List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> result = new ArrayList<>();
        traverse(root, result);
        return result;
    }

    // traverse рекурсивно проходит по дереву и добавляет значения узлов в ответ.
    // time: O(n), space: O(h)
    private static void traverse(TreeNode node, List<Integer> result) {
        if (node == null) {
            return;
        }
        // In order - левый потомок, текущий узел, правый потомок.
        traverse(node.left, result);           // Добавляем значение левого потомка
        result.add(node.val);                   // Добавляем значение текущего узла в ответ
        traverse(node.right, result);          // Добавляем значение правого потомка
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
