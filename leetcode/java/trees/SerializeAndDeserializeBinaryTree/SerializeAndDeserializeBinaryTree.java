package leetcode.java.trees.SerializeAndDeserializeBinaryTree;

import java.util.Arrays;
import java.util.LinkedList;
import java.util.Queue;

/* 297. Serialize and Deserialize Binary Tree
https://leetcode.com/problems/serialize-and-deserialize-binary-tree/description/

Сериализация — это процесс преобразования структуры данных или объекта в последовательность битов, чтобы ее можно было
сохранить в файле или буфере памяти или передать по сетевому соединению для последующего восстановления в той же или
другой компьютерной среде.

Разработайте алгоритм сериализации и десериализации двоичного дерева. Нет никаких ограничений на то, как должен работать
ваш алгоритм сериализации/десериализации. Вам просто нужно убедиться, что двоичное дерево может быть сериализовано в
строку, а эта строка может быть десериализована в исходную древовидную структуру.

Пояснение: формат ввода/вывода такой же, как и формат, в котором LeetCode сериализует двоичное дерево. Вам не
обязательно следовать этому формату, поэтому, пожалуйста, проявите креативность и придумайте другие подходы
самостоятельно.

Input: root = [1,2,3,null,null,4,5]
Output: [1,2,3,null,null,4,5]
*/

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode(int x) {
        val = x;
    }

    // Дополнительный конструктор для удобного создания узла с потомками
    TreeNode(int x, TreeNode left, TreeNode right) {
        val = x;
        this.left = left;
        this.right = right;
    }
}

public class SerializeAndDeserializeBinaryTree {
    public static void main(String[] args) {
        SerializeAndDeserializeBinaryTree codec = new SerializeAndDeserializeBinaryTree();

        // Пример дерева: [1,2,3,null,null,4,5]
        TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.right = new TreeNode(3);
        root.right.left = new TreeNode(4);
        root.right.right = new TreeNode(5);

        // Сериализация
        String serialized = codec.serialize(root);
        System.out.println("Serialized: " + serialized);

        // Десериализация
        TreeNode deserialized = codec.deserialize(serialized);

        // Проверка
        System.out.println("Deserialized root: " + deserialized.val);
    }

    // Сериализует дерево в одну строку
    private String serialize(TreeNode root) {
        StringBuilder sb = new StringBuilder();
        serializeHelper(root, sb);
        return sb.toString();
    }

    private void serializeHelper(TreeNode node, StringBuilder sb) {
        if (node == null) {
            sb.append("N,");
        } else {
            sb.append(node.val).append(",");
            serializeHelper(node.left, sb);
            serializeHelper(node.right, sb);
        }
    }

    // Десериализует закодированные данные в дерево
    private TreeNode deserialize(String data) {
        // Разбиваем строку на токены и помещаем их в очередь для последовательного потребления
        Queue<String> queue = new LinkedList<>(Arrays.asList(data.split(",")));
        return deserializeHelper(queue);
    }

    private TreeNode deserializeHelper(Queue<String> queue) {
        String token = queue.poll();
        if (token == null || token.equals("N")) {
            return null;
        }

        // Аргументы в Java вычисляются строго слева направо, 
        // что гарантирует правильный порядок обхода (сначала левое, потом правое поддерево)
        return new TreeNode(Integer.parseInt(token), deserializeHelper(queue), deserializeHelper(queue));
    }
}
