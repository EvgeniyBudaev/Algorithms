package leetcode.java.intervals.MinimumNumberOfArrowsToBurstBalloons;

import java.util.Arrays;

/* 452. Minimum Number of Arrows to Burst Balloons
https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/description/

На плоской стене, представляющей плоскость XY, приклеены несколько сферических воздушных шаров.
Воздушные шары представлены в виде двумерного целочисленного массива точек, где points[i] = [xstart, xend]
обозначает воздушный шар, горизонтальный диаметр которого простирается между xstart и xend. Вы не знаете точных
координат y воздушных шаров.

Стрелы можно выпускать вертикально вверх (в положительном направлении y) из разных точек вдоль оси x.
Воздушный шар с xstart и xend лопается стрелой, выпущенной в x, если xstart <= x <= xend.
Нет ограничений на количество выпущенных стрел. Выпущенная стрела продолжает двигаться вверх бесконечно, лопая все
воздушные шары на своем пути.

Учитывая массив точек, верните минимальное количество стрел, которое необходимо выпустить, чтобы лопнуть все воздушные шары.

Example 1:
Input: points = [[10,16],[2,8],[1,6],[7,12]]
Output: 2
Объяснение: Воздушные шары можно лопнуть двумя стрелами:
- Выстрелите стрелой в точку x = 6, лопнув шары [2,8] и [1,6].
- Выстрелите стрелой в точку x = 11, лопнув шары [10,16] и [7,12].
*/

public class MinimumNumberOfArrowsToBurstBalloons {
    public static void main(String[] args) {
        int[][] points = {{10, 16}, {2, 8}, {1, 6}, {7, 12}};
        System.out.println(findMinArrowShots(points)); // 2   
    }

    // findMinArrowShots - функция, которая находит минимальное количество стрел, которое необходимо выпустить, чтобы лопнуть все воздушные шары.
    // time: O(n * log n), space: O(1)
    private static int findMinArrowShots(int[][] points) {
        // Если нет воздушных шаров, то нет стрел
        if (points.length == 0) {
            return 0;
        }

        // Сортируем интервалы по начальной точке
        Arrays.sort(points, (a, b) -> Integer.compare(a[0], b[0]));
        // [[1, 6] [2, 8] [7, 12] [10, 16]]

        int result = 1;                      // Первый интервал всегда лопается стрелой
        int[] lastInterval = points[0];   // Последний интервал, который лопается стрелой

        for (int[] interval : points) {
            // Если интервал пересекается с последним лопающим интервалом,
            // обновляем зону пересечения для текущей стрелы
            if (isOverlapping(lastInterval, interval)) {
                lastInterval = overlapTwoIntervals(lastInterval, interval);
                continue;
            }
            // Иначе нужна новая стрела
            lastInterval = interval;  // Обновляем последний лопающий интервал
            result++;                 // Увеличиваем количество стрел
        }

        return result;
    }

    // isOverlapping - функция, которая проверяет, пересекаются ли два интервала.
    private static boolean isOverlapping(int[] a, int[] b) {
        return Math.max(a[0], b[0]) <= Math.min(a[1], b[1]);
    }

    // overlapTwoIntervals - функция, которая находит пересечение двух интервалов.
    private static int[] overlapTwoIntervals(int[] a, int[] b) {
        // Возвращаем зону пересечения: [максимум начал, минимум концов]
        return new int[]{Math.max(a[0], b[0]), Math.min(a[1], b[1])};
    }
}
