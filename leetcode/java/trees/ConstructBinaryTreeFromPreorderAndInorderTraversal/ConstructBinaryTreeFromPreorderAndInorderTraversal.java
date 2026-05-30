package leetcode.java.trees.ConstructBinaryTreeFromPreorderAndInorderTraversal;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.Queue;

/* 105. Construct Binary Tree from Preorder and Inorder Traversal
https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/

Даны два целочисленных массива preorder и inorder, где preorder — это предварительный обход двоичного дерева,
а inorder — симметричный обход того же дерева. Построить и вернуть двоичное дерево.

Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
*/

public class ConstructBinaryTreeFromPreorderAndInorderTraversal {
    public static void main(String[] args) {
        int[] preorder = {3, 9, 20, 15, 7};
        int[] inorder = {9, 3, 15, 20, 7};
        System.out.println(printTree(buildTree(preorder, inorder))); // [3,9,20,null,null,15,7] 
    }

    static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode(int val) {
            this.val = val;
        }
    }

    // buildTree строит бинарное дерево на основе предварительного и симметричного обхода.
    // time: O(n), где n - количество узлов в дереве
    // space: O(n), где n - количество узлов в дереве
    private static TreeNode buildTree(int[] preorder, int[] inorder) {
        // Создаем карту для быстрого поиска индекса элемента в inorder
        Map<Integer, Integer> inorderMap = new HashMap<>();
        for (int i = 0; i < inorder.length; i++) {
            inorderMap.put(inorder[i], i); // Ключом является значение элемента, а значением - его индекс
        }

        // Используем массив из одного элемента для имитации изменяемого индекса (аналог замыкания в Go)
        int[] preIndex = {0};

        // Вызываем вспомогательную функцию с начальными индексами
        return helper(preorder, inorderMap, preIndex, 0, inorder.length - 1);
    }

    // Вспомогательная функция для рекурсивного построения дерева
    private static TreeNode helper(int[] preorder, Map<Integer, Integer> inorderMap,
                                   int[] preIndex, int left, int right) {
        if (left > right) { // Если индексы пересекаются, то дерево закончилось
            return null;
        }

        int nodeVal = preorder[preIndex[0]]; // Значение текущего элемента в preorder
        preIndex[0]++; // Переходим к следующему элементу в preorder

        TreeNode node = new TreeNode(nodeVal); // Создаем узел с текущим значением
        int idx = inorderMap.get(nodeVal); // Находим индекс текущего элемента в inorder

        node.left = helper(preorder, inorderMap, preIndex, left, idx - 1); // Рекурсивно строим левое поддерево
        node.right = helper(preorder, inorderMap, preIndex, idx + 1, right); // Рекурсивно строим правое поддерево

        return node; // Возвращаем созданный узел
    }

    // createTree создает бинарное дерево из массива значений.
    private static TreeNode createTree(Integer[] values) {
        if (values == null || values.length == 0 || values[0] == null) {
            return null;
        }

        TreeNode root = new TreeNode(values[0]);
        Queue<TreeNode> queue = new LinkedList<>();
        queue.offer(root);
        int i = 1;

        while (!queue.isEmpty() && i < values.length) {
            TreeNode node = queue.poll();

            if (i < values.length && values[i] != null) {
                node.left = new TreeNode(values[i]);
                queue.offer(node.left);
            }
            i++;

            if (i < values.length && values[i] != null) {
                node.right = new TreeNode(values[i]);
                queue.offer(node.right);
            }
            i++;
        }

        return root;
    }

    // printTree - функция для вывода дерева в формате [val,left,right]
    // time: O(n), space: O(n)
    private static String printTree(TreeNode root) {
        if (root == null) {
            return "[]";
        }

        List<String> result = new ArrayList<>();
        Queue<TreeNode> queue = new LinkedList<>();
        queue.offer(root);

        while (!queue.isEmpty()) {
            TreeNode node = queue.poll();

            if (node == null) {
                result.add("null");
                continue;
            }

            result.add(String.valueOf(node.val));
            queue.offer(node.left);
            queue.offer(node.right);
        }

        // Удаляем trailing nulls
        int i = result.size() - 1;
        while (i >= 0 && result.get(i).equals("null")) {
            i--;
        }
        result = result.subList(0, i + 1);

        return "[" + String.join(",", result) + "]";
    }
}
