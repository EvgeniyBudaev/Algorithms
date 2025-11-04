package array.PascalsTriangle;

import java.util.ArrayList;
import java.util.List;

/* 118. Pascal's Triangle
https://leetcode.com/problems/pascals-triangle/description/

Дано целое число numRows, вернуть первые numRows треугольника Паскаля.
В треугольнике Паскаля каждое число является суммой двух чисел, расположенных непосредственно над ним, как показано:

Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
*/

public class PascalsTriangle {
    public static void main(String[] args) {
        int numRows = 5;
        System.out.println(generate(numRows)); // [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
    }

    // generate возвращает первые numRows треугольника Паскаля.
    // time: O(numRows^2) - количество итераций в цикле в цикле
    // space: O(numRows^2) - количество итераций в цикле в цикле
    private static List<List<Integer>> generate(int numRows) {
        List<List<Integer>> triangle = new ArrayList<>(); // Треугольник Паскаля

        if (numRows == 0) { // Если numRows равно 0, то возвращаем пустой треугольник Паскаля
            return triangle;
        }

        triangle.add(new ArrayList<>()); // Первый ряд треугольника Паскаля
        triangle.get(0).add(1); // Первый ряд всегда состоит из одного числа 1

        for (int i = 1; i < numRows; i++) {
            List<Integer> prevRow = triangle.get(i - 1); // Предыдущий ряд треугольника Паскаля
            List<Integer> newRow = new ArrayList<>(); // Текущий ряд треугольника Паскаля
            newRow.add(1); // Первый элемент текущего ряда всегда равен 1

            for (int j = 1; j < prevRow.size(); j++) { // Для каждого элемента в предыдущем ряду
                newRow.add(prevRow.get(j - 1) + prevRow.get(j)); // Сумма двух элементов, расположенных непосредственно над ним
            }

            newRow.add(1); // Последний элемент текущего ряда всегда равен 1
            triangle.add(newRow); // Добавляем текущий ряд в треугольник Паскаля
        }

        return triangle; // Возвращаем треугольник Паскаля
    }
}
