package array.IncreasingTripletSubsequence;

/* 334. Increasing Triplet Subsequence
https://leetcode.com/problems/increasing-triplet-subsequence/description/

Для заданного целочисленного массива nums вернуть true, если существует тройка индексов (i, j, k) такая,
что i < j < k и nums[i] < nums[j] < nums[k]. Если таких индексов не существует, вернуть false.

Input: nums = [1,2,3,4,5]
Output: true
Пояснение: Любая тройка, где i < j < k, действительна.
*/

public class IncreasingTripletSubsequence {
    public static void main(String[] args) {
        int[] nums = {1, 2, 3, 4, 5};
        System.out.println(increasingTriplet(nums)); // true
    }

    // increasingTriplet проверяет, существует ли возрастающая триплетная подпоследовательность
    // в заданном целочисленном массиве.
    // time: O(n), space: O(1)
    private static boolean increasingTriplet(int[] nums) {
        // Инициализируйте две переменные для отслеживания наименьшего и второго наименьшего элементов.
        int first = Integer.MAX_VALUE;
        int second = Integer.MAX_VALUE;

        // Перебрать элементы входного массива
        for (int num : nums) {
            // Проверьте, меньше ли текущее число или равно ли оно первому наименьшему числу.
            if (num <= first) {
                first = num;
            } else if (num <= second) { // Проверьте, меньше ли текущее число или равно ли оно второму наименьшему числу.
                second = num;
            } else { // Если ни одно из условий не выполняется, то находится возрастающая триплетная подпоследовательность
                return true;
            }
        }

        // Не обнаружено возрастающей триплетной подпоследовательности
        return false;
    }
}
