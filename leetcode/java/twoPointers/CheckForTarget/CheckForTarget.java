package twoPointers.CheckForTarget;

/*
Учитывая отсортированный массив уникальных целых чисел и целевое целое число, верните true, если существует пара чисел,
сумма которых равна целевому значению, и false в противном случае. Эта задача аналогична задаче «Две суммы».
(В режиме Two Sum входные данные не сортируются).

Например, если заданы числа = [1, 2, 4, 6, 8, 9, 14, 15] и цель = 13, верните true, потому что 4 + 9 = 13.
*/

public class CheckForTarget {
    public static void main(String[] args) {
        int[] nums = {1, 2, 4, 6, 8, 9, 14, 15};
        System.out.println(checkForTarget(nums, 13)); // true
    }

    // checkForTarget проверяет, существует ли в отсортированном массиве nums пара чисел, сумма которых равна заданному
    // целевому значению target.
    // time: O(n), space: O(1)
    private static boolean checkForTarget(int[] nums, int target) {
        int left = 0, right = nums.length - 1; // левая и правая границы массива

        while (left < right) { // пока левая граница меньше правой
            int sum = nums[left] + nums[right]; // сумма пары
            if (sum == target) {
                return true;
            } else if (sum > target) { // если сумма больше целевого значения, то уменьшаем правую границу
                right--;
            } else { // если сумма меньше целевого значения, то увеличиваем левую границу
                left++;
            }
        }

        return false; // если не найдено
    }
}
