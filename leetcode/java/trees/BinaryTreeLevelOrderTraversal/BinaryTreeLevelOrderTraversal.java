package leetcode.java.trees.BinaryTreeLevelOrderTraversal;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

/* 102. Binary Tree Level Order Traversal
https://leetcode.com/problems/binary-tree-level-order-traversal/description/

Учитывая корень двоичного дерева, вернуть порядок обхода значений его узлов (т. е. слева направо, уровень за уровнем).

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]

Example 2:
Input: root = [1]
Output: [[1]]
*/

public class BinaryTreeLevelOrderTraversal {
    public static void main(String[] args) {
        TreeNode root = createTree(new Integer[]{3, 9, 20, null, null, 15, 7});
        System.out.println(levelOrder(root)); // [[3], [9, 20], [15, 7]] 
    }

    static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode(int val) {
            this.val = val;
        }
    }

    // levelOrder возвращает список значений узлов дерева в порядке обхода уровнями.
    // time: O(n), space: O(n), где n - количество узлов в дереве
    private static List<List<Integer>> levelOrder(TreeNode root) {
        List<List<Integer>> levels = new ArrayList<>();
        preOrder(root, 0, levels);
        return levels;
    }

    // preOrder рекурсивно обходит дерево и заполняет уровни.
    // time: O(n), space: O(n)
    private static void preOrder(TreeNode node, int level, List<List<Integer>> levels) {
        // Если узел равен null, возвращаемся
        if (node == null) {
            return;
        }

        // Если текущего уровня еще нет в списке, добавляем новый уровень
        if (level == levels.size()) {
            levels.add(new ArrayList<>());
        }

        // Делаем PreOrder обход дерева и добавляем вершину на текущий уровень в конец списка.
        levels.get(level).add(node.val);
        // Рекурсивно обходим левое и правое поддеревья
        preOrder(node.left, level + 1, levels);
        preOrder(node.right, level + 1, levels);
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
