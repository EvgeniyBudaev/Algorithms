package leetcode.java.trees.ValidateBinarySearchTree;

import java.util.LinkedList;
import java.util.Queue;

/* 98. Validate Binary Search Tree
https://leetcode.com/problems/validate-binary-search-tree/description/

Учитывая корень двоичного дерева, определите, является ли оно допустимым двоичным деревом поиска (BST).
Допустимое BST определяется следующим образом:

Левое поддерево узла содержит только узлы с ключами, меньшими ключа узла.
Правильное поддерево узла содержит только узлы с ключами, большими ключа узла.
И левое, и правое поддеревья также должны быть двоичными деревьями поиска.

Input: root = [2,1,3]
Output: true

Input: root = [5,1,4,null,null,3,6]
Output: false
Пояснение: Значение корневого узла равно 5, но значение его правого дочернего узла равно 4.
*/

public class ValidateBinarySearchTree {
    public static void main(String[] args) {
        TreeNode root1 = createTree(new Integer[]{2, 1, 3});
        System.out.println(isValidBST(root1)); // true

        TreeNode root2 = createTree(new Integer[]{5, 1, 4, null, null, 3, 6});
        System.out.println(isValidBST(root2)); // false
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

    // isValidBST проверяет, является ли дерево корректным BST.
    // time: O(n), space: O(h) (где h - высота дерева для стека рекурсии)
    private static boolean isValidBST(TreeNode root) {
        long[] minVale = {Long.MIN_VALUE};
        boolean[] isValid = {true};

        visit(root, minVale, isValid);

        return isValid[0];
    }

    // visit обходит дерево in-order (левое поддерево, корень, правое поддерево)
    // и проверяет условие BST.
    private static void visit(TreeNode node, long[] minVale, boolean[] isValid) {
        if (node == null) {
            return;
        }

        visit(node.left, minVale, isValid); // Обходим левое поддерево

        // Если значение узла меньше или равно минимальному значению, то дерево не является корректным BST
        if (node.val <= minVale[0]) {
            isValid[0] = false;
        }

        minVale[0] = node.val; // Обновляем минимальное значение
        visit(node.right, minVale, isValid); // Обходим правое поддерево
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
