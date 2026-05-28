package leetcode.java.trees.BinaryTreeRightSideView;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

/* 199. Binary Tree Right Side View
https://leetcode.com/problems/binary-tree-right-side-view/description/

Имея корень двоичного дерева, представьте, что вы стоите с правой стороны от него, и верните значения узлов,
которые вы видите, упорядоченные сверху вниз.

Example 1:
Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]
*/

public class BinaryTreeRightSideView {
    public static void main(String[] args) {
        //TreeNode root = createTree(new Integer[]{1, 2, 3, null, 5, null, 4});
        TreeNode root = createTree(new Integer[]{1, 2, 3});
        System.out.println(rightSideView(root)); // [1, 3]
    }

    static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode(int val) {
            this.val = val;
        }
    }

    // rightSideView возвращает список значений узлов, которые вы видите с правой стороны, 
    // упорядоченные сверху вниз.
    // time: O(n), space: O(n), где n - количество узлов в дереве
    private static List<Integer> rightSideView(TreeNode root) {
        List<Integer> result = new ArrayList<>();
        preOrder(root, 0, result);
        return result;
    }

    // preOrder рекурсивно обходит дерево и заполняет уровни.
    // time: O(n), space: O(n)
    private static void preOrder(TreeNode node, int level, List<Integer> list) {
        // Если узел равен null, возвращаемся
        if (node == null) {
            return;
        }

        // Делаем PreOrder и на каждом уровне запоминаем последнюю вершину для уровня.
        // Если текущего уровня еще нет в списке, добавляем новый уровень
        if (level == list.size()) {
            list.add(0);
        }

        // Запоминаем последнюю вершину для уровня (перезаписываем значение)
        list.set(level, node.val);
        // Рекурсивно обходим левое и правое поддеревья
        preOrder(node.left, level + 1, list);
        preOrder(node.right, level + 1, list);
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
