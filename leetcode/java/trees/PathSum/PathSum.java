package leetcode.java.trees.PathSum;

import java.util.LinkedList;
import java.util.Queue;

/* 112. Path Sum
https://leetcode.com/problems/path-sum/description/

При наличии корня бинарного дерева и целочисленной targetSum вернуть true, если дерево имеет путь от корня к листьям,
такой, что сложение всех значений вдоль пути равно targetSum.
Лист — это узел без дочерних элементов.

Example 1:
Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
Output: true
Пояснение: Показан путь от корня к листьям с целевой суммой.
*/

public class PathSum {
    public static void main(String[] args) {
        TreeNode root = createTree(new Integer[]{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, null, 1});
        int targetSum = 22;
        System.out.println(hasPathSum(root, targetSum)); // true
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
     * hasPathSum проверяет, есть ли в дереве путь от корня до листа, сумма чисел которого равна заданной сумме.
     * time: O(n), где n - количество узлов в дереве
     * space: O(h), где h - высота дерева. В худшем случае h = n.
     */
    private static boolean hasPathSum(TreeNode root, int targetSum) {
        return checkPathSum(root, 0, targetSum);
    }

    /**
     * checkPathSum проверяет наличие пути от заданного узла до листа, сумма чисел которого равна заданной сумме.
     * time: O(n), где n - количество узлов в дереве
     * space: O(h), где h - высота дерева. В худшем случае h = n.
     */
    private static boolean checkPathSum(TreeNode node, int currentSum, int targetSum) {
        // Если узел пустой, возвращаем false
        if (node == null) {
            return false;
        }

        // currentSum - префиксная сумма.
        // Добавляем значение текущего узла к сумме
        currentSum += node.val;

        // Если это лист и сумма совпадает
        if (isLeaf(node) && currentSum == targetSum) {
            return true;
        }

        // Проверяем левое и правое поддеревья
        return checkPathSum(node.left, currentSum, targetSum) ||
                checkPathSum(node.right, currentSum, targetSum);
    }

    /**
     * isLeaf проверяет, является ли узел листом.
     */
    private static boolean isLeaf(TreeNode node) {
        return node.left == null && node.right == null;
    }

    /**
     * createTree создает бинарное дерево из массива значений (обход в ширину / level-order).
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
            // Если правый потомок не равен null, то создаем его и добавляем его в очередь
            if (i < values.length && values[i] != null) {
                node.right = new TreeNode(values[i]); // Создаем правого потомка
                queue.add(node.right);                // Добавляем правого потомка в очередь
            }
            i++;
        }

        // Возвращаем указатель на корневой узел
        return root;
    }
}
