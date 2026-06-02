package leetcode.java.trees.InorderSuccessorInBST;

import java.util.LinkedList;
import java.util.Queue;

/* 285. Inorder Successor in BST
https://leetcode.ca/2016-09-10-285-Inorder-Successor-in-BST/

Если задан корень бинарного дерева поиска и узел p в нем, вернуть упорядоченного преемника этого узла в BST.
Если у заданного узла нет упорядоченного преемника в дереве, вернуть null.

Узел-преемник узла p — это узел с наименьшим ключом, большим, чем p.val.

Input: root = [2,1,3], p = 1
Output: 2
Объяснение: следующий узел 1 по порядку — 2. Обратите внимание, что и p, и возвращаемое значение имеют тип TreeNode.
*/

public class InorderSuccessorInBST {
    public static void main(String[] args) {
        TreeNode root = createTree(new Integer[]{2, 1, 3});
        TreeNode p = createTree(new Integer[]{1});
        TreeNode successor = inorderSuccessor(root, p);
        System.out.println(successor.val); // 2
    }

    static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode(int val) {
            this.val = val;
        }
    }

    // inorderSuccessor возвращает упорядоченного преемника узла p в бинарном дереве поиска.
    // time: O(h), где h - высота дерева.
    // space: O(1)
    private static TreeNode inorderSuccessor(TreeNode root, TreeNode p) {
        TreeNode ans = null; // Переменная для хранения упорядоченного преемника

        while (root != null) { // Обход дерева
            // Если значение текущего узла больше значения узла p, то следующий узел по порядку может быть в левом поддереве.
            if (root.val > p.val) {
                ans = root; // Обновляем упорядоченного преемника до текущего узла
                root = root.left; // Переходим в левое поддерево
            } else {
                root = root.right; // Переходим в правое поддерево
            }
        }

        return ans; // Возвращаем упорядоченного преемника
    }

    // createTree создает бинарное дерево из массива значений.
    private static TreeNode createTree(Integer[] values) {
        if (values == null || values.length == 0 || values[0] == null) {
            return null;
        }

        TreeNode root = new TreeNode(values[0]); // Создаем корневой узел с первым значением в массиве
        Queue<TreeNode> queue = new LinkedList<>(); // Создаем очередь для обхода дерева по уровням
        queue.add(root);
        int i = 1; // Индекс для обхода массива значений

        while (!queue.isEmpty() && i < values.length) {
            TreeNode node = queue.poll(); // Берем первый элемент из очереди

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
