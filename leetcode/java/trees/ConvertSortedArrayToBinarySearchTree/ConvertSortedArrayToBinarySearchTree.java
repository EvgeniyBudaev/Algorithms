package leetcode.java.trees.ConvertSortedArrayToBinarySearchTree;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

/* 108. Convert Sorted Array to Binary Search Tree
https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/

Дан целочисленный массив nums, элементы которого отсортированы в порядке возрастания, преобразовать его в
сбалансированное по высоте двоичное дерево поиска.

Input: nums = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Пояснение: [0,-10,5,null,-3,null,9] также принимается
*/

public class ConvertSortedArrayToBinarySearchTree {
    public static void main(String[] args) {
        int[] nums = {-10, -3, 0, 5, 9};
        System.out.println(printTree(sortedArrayToBST(nums))); // [0,-3,9,-10,null,5]  
    }

    static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode(int val) {
            this.val = val;
        }
    }

    // sortedArrayToBST - рекурсивный алгоритм, который берет средний элемент массива и делит его на два подмассива, которые будут
    // являться левым и правым поддеревом. После чего рекурсивно вызывает функцию для каждого подмассива.
    // time: O(n), space: O(n)
    private static TreeNode sortedArrayToBST(int[] nums) {
        return buildBST(nums, 0, nums.length - 1);
    }

    private static TreeNode buildBST(int[] nums, int left, int right) {
        if (left > right) return null;

        int mid = left + (right - left) / 2; // Защита от переполнения
        TreeNode node = new TreeNode(nums[mid]);
        node.left = buildBST(nums, left, mid - 1);
        node.right = buildBST(nums, mid + 1, right);
        return node;
    }

    // printTree - функция для вывода дерева в формате [val,left,right]
    // time: O(n), space: O(n)
    private static String printTree(TreeNode root) {
        if (root == null) { // Если дерево пустое, возвращаем пустой срез
            return "[]";
        }

        List<String> result = new ArrayList<>(); // Список для хранения значений узлов
        Queue<TreeNode> queue = new LinkedList<>(); // Используем очередь для BFS
        queue.offer(root);

        while (!queue.isEmpty()) { // BFS
            TreeNode node = queue.poll(); // Берем первый узел из очереди

            if (node == null) { // Если узел пустой, добавляем null в результат
                result.add("null"); // добавляем null в результат
                continue; // Пропускаем остальные шаги
            }

            result.add(String.valueOf(node.val)); // Добавляем значение узла в результат
            queue.offer(node.left); // Добавляем левого потомка в очередь
            queue.offer(node.right); // Добавляем правого потомка в очередь
        }

        // Удаляем trailing nulls
        int i = result.size() - 1; // Инициализируем индекс последнего ненулевого элемента
        while (i >= 0 && result.get(i).equals("null")) { // Если элемент равен null, уменьшаем индекс
            i--; // уменьшаем индекс
        }
        result = result.subList(0, i + 1); // Обрезаем ненужные элементы

        return "[" + String.join(",", result) + "]"; // Возвращаем результат в виде строки
    }
}
