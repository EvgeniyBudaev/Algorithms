package leetcode.java.dynamicProgramming.UniquePathsII;

/* 63. Unique Paths II
https://leetcode.com/problems/unique-paths-ii/description/

Вам дан массив целых чисел размером m x n. Изначально в верхнем левом углу находится робот (т. е. сетка[0][0]).
 Робот пытается переместиться в правый нижний угол (т. е. сетку[m - 1][n - 1]).
 В любой момент времени робот может двигаться только вниз или вправо.
Препятствие и пространство помечаются в сетке цифрами 1 или 0 соответственно. Путь, по которому движется робот, не может
 включать квадраты, являющиеся препятствиями.
Возвращает количество возможных уникальных путей, по которым робот может добраться до правого нижнего угла.
Тестовые примеры генерируются так, что ответ будет меньше или равен 2 * 109.

Input: obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
Output: 2
Пояснение: В центре сетки 3х3 выше находится одно препятствие.
Есть два способа добраться до правого нижнего угла:
1. Вправо -> Вправо -> Вниз -> Вниз.
2. Вниз -> Вниз -> Вправо -> Вправо.
*/

public class UniquePathsII {
    public static void main(String[] args) {
        int[][] grid = {
                {0, 0, 0},
                {0, 1, 0},
                {0, 0, 0}
        };
        System.out.println(uniquePathsWithObstacles(grid)); // 2
    }

    // uniquePathsWithObstacles вычисляет количество уникальных путей в сетке с препятствиями
    // time: O(m * n), space: O(n)
    public int uniquePathsWithObstacles(int[][] obstacleGrid) {
        if (obstacleGrid == null || obstacleGrid.length == 0 || obstacleGrid[0].length == 0 || obstacleGrid[0][0] == 1) {
            return 0;
        }

        int m = obstacleGrid.length;      // Количество строк
        int n = obstacleGrid[0].length;   // Количество столбцов

        // Используем два массива для оптимизации памяти
        int[] previous = new int[n];
        int[] current = new int[n];

        // Инициализация начальной точки
        previous[0] = 1;

        for (int i = 0; i < m; i++) {
            // Обработка первого столбца
            if (obstacleGrid[i][0] == 1) {
                current[0] = 0;
            } else {
                current[0] = previous[0];
            }

            // Обработка остальных столбцов
            for (int j = 1; j < n; j++) {
                if (obstacleGrid[i][j] == 1) {
                    current[j] = 0;
                } else {
                    current[j] = current[j - 1] + previous[j];
                }
            }

            // Обновляем previous для следующей строки
            System.arraycopy(current, 0, previous, 0, n);
        }

        return previous[n - 1];
    }
}
