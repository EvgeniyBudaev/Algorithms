package twoPointers.IntervalListIntersections;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/* 986. Interval List Intersections
https://leetcode.com/problems/interval-list-intersections/description/

Вам даны два списка закрытых интервалов, firstList и SecondList,
где firstList[i] = [starti, endi] и SecondList[j] = [startj, endj]. Каждый список интервалов попарно непересекающийся и
отсортирован.
Верните пересечение этих двух списков интервалов.
Замкнутый интервал [a, b] (с a <= b) обозначает набор действительных чисел x с a <= x <= b.
Пересечение двух закрытых интервалов представляет собой набор действительных чисел, которые либо пусты, либо
представлены в виде замкнутого интервала. Например, пересечение [1, 3] и [2, 4] — это [2, 3].

Input: firstList = [[0,2],[5,10],[13,23],[24,25]], secondList = [[1,5],[8,12],[15,24],[25,26]]
Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
*/

public class IntervalListIntersections {
    public static void main(String[] args) {
        int[][] firstList = {{0, 2}, {5, 10}, {13, 23}, {24, 25}};
        int[][] secondList = {{1, 5}, {8, 12}, {15, 24}, {25, 26}};

        int[][] result = intervalIntersection(firstList, secondList);

        // Вывод результата
        for (int[] interval : result) {
            System.out.println(Arrays.toString(interval));
        }
        // Output: [1, 2], [5, 5], [8, 10], [15, 23], [24, 24], [25, 25]
    }

    // intervalIntersection находит пересечение двух списков интервалов.
    // time: O(n), space: O(n)
    private static int[][] intervalIntersection(int[][] firstList, int[][] secondList) {
        List<int[]> result = new ArrayList<>();
        int left = 0, right = 0;

        // Пока не закончатся интервалы в обоих списках
        while (left < firstList.length && right < secondList.length) {
            int[] first = firstList[left]; // Текущий интервал первого списка
            int[] second = secondList[right]; // Текущий интервал второго списка

            int start = Math.max(first[0], second[0]); // Максимальное начало
            int end = Math.min(first[1], second[1]); // Минимальное окончание

            // Если начало меньше окончания, то добавляем интервал в результат
            if (start <= end) {
                result.add(new int[]{start, end});
            }

            // Если окончание первого интервала меньше второго, то двигаем указатель первого интервала, иначе двигаем второй
            if (first[1] < second[1]) {
                left++;
                // Если окончание первого интервала больше второго, то двигаем указатель второго интервала
            } else if (first[1] > second[1]) {
                right++;
            } else {
                left++;
                right++;
            }
        }

        return result.toArray(int[][]::new); // Список интервалов
        // return result.stream().toArray(int[][]::new);
    }
}
