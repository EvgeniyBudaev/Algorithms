package leetcode.java.intervals.MergeInterval;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/* 56. Merge Intervals
https://leetcode.com/problems/merge-intervals/description/

Дан массив интервалов, где intervals[i] = [starti, endi], объединить все перекрывающиеся интервалы и вернуть массив
неперекрывающихся интервалов, которые покрывают все интервалы во входных данных.

Example 1:
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Пояснение: Поскольку интервалы [1,3] и [2,6] перекрываются, объединим их в [1,6].
*/

public class MergeInterval {
    public static void main(String[] args) {
        int[][] intervals = {{1, 3}, {2, 6}, {8, 10}, {15, 18}};
        int[][] merged = merge(intervals);
        System.out.println(Arrays.deepToString(merged)); // [[1, 6], [8, 10], [15, 18]]
    }

    // merge принимает массив интервалов и объединяет все перекрывающиеся интервалы.
    // time: O(n*log(n)), space: O(1)
    private static int[][] merge(int[][] intervals) {
        if (intervals.length == 0) {
            return new int[0][0];
        }

        // Сортируем интервалы по начальному значению. O(n*log(n))
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[0], b[0]));

        List<int[]> result = new ArrayList<>();
        result.add(intervals[0]); // Результат, начинаем с первого интервала

        // O(n)
        for (int i = 1; i < intervals.length; i++) {
            int[] current = intervals[i];                    // Текущий интервал
            int[] last = result.get(result.size() - 1);      // Последний интервал в результате

            // если текущий интервал и последний в ответе пересекаются,
            // значит объединяем их, иначе добавляем интервал к ответу
            if (isOverlapping(last, current)) {
                result.set(result.size() - 1, mergeTwoIntervals(last, current));
            } else {
                result.add(current);
            }
        }

        return result.toArray(int[][]::new);
    }

    // isOverlapping проверяет, пересекаются ли два интервала.
    private static boolean isOverlapping(int[] a, int[] b) {
        return Math.max(a[0], b[0]) <= Math.min(a[1], b[1]);
    }

    // mergeTwoIntervals объединяет два интервала.
    private static int[] mergeTwoIntervals(int[] a, int[] b) {
        // Интервалы обязательно должны пересекаться
        return new int[]{a[0], Math.max(a[1], b[1])};
    }
}
