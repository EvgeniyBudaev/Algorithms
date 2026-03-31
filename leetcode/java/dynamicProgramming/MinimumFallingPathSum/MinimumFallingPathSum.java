package leetcode.java.dynamicProgramming.MinimumFallingPathSum;

/* 931. Minimum Falling Path Sum
https://leetcode.com/problems/minimum-falling-path-sum/description/

Дана матрица n x n целых чисел, вернуть минимальную сумму любого нисходящего пути через матрицу.

Нисходящий путь начинается с любого элемента в первой строке и выбирает элемент в следующей строке, который находится
либо непосредственно под ним, либо по диагонали слева/справа. В частности, следующим элементом из позиции (row, col)
будет (row + 1, col - 1), (row + 1, col) или (row + 1, col + 1).

Input: matrix = [[2,1,3],[6,5,4],[7,8,9]]
Output: 13
Пояснение: Существуют два падающих пути с минимальной суммой, как показано на рисунке.
*/

public class MinimumFallingPathSum {
    public static void main(String[] args) {
        int[][] matrix = {
            {2, 1, 3},
            {6, 5, 4},
            {7, 8, 9}
        };
        System.out.println(minFallingPathSum(matrix)); // 13 (путь 1 → 5 → 7)
    }

    // minFallingPathSum вычисляет минимальную сумму падающего пути в матрице
    // time: O(n^2), где n - размер матрицы (n x n)
    // space: O(n^2) для DP матрицы
    private static int minFallingPathSum(int[][] matrix) {
        int n = matrix.length;
        if (n == 0) {
            return 0;
        }

        // Создаем копию матрицы для хранения промежуточных результатов
        int[][] dp = new int[n][n];
        for (int i = 0; i < n; i++) {
            System.arraycopy(matrix[i], 0, dp[i], 0, n);
        }

        for (int i = 1; i < n; i++) {
            for (int j = 0; j < n; j++) {
                // Инициализируем минимальную сумму как текущую ячейку + значение сверху
                int minSum = dp[i-1][j] + matrix[i][j];

                // Проверяем левую верхнюю диагональ (если существует)
                if (j > 0) {
                    minSum = Math.min(minSum, dp[i-1][j-1] + matrix[i][j]);
                }

                // Проверяем правую верхнюю диагональ (если существует)
                if (j < n - 1) {
                    minSum = Math.min(minSum, dp[i-1][j+1] + matrix[i][j]);
                }

                dp[i][j] = minSum;
            }
        }

        // Находим минимальное значение в последней строке
        int minTotal = dp[n-1][0];
        for (int j = 1; j < n; j++) {
            if (dp[n-1][j] < minTotal) {
                minTotal = dp[n-1][j];
            }
        }

        return minTotal;
    }
}
