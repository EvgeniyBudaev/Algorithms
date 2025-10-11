package array.FindPivotIndex;

/* 724. Find Pivot Index
https://leetcode.com/problems/find-pivot-index/description/

Учитывая массив целых чисел nums, вычислите индекс поворота этого массива.
Индекс поворота — это индекс, где сумма всех чисел строго слева от индекса равна сумме всех чисел строго справа от индекса.
Если индекс находится на левом краю массива, то левая сумма равна 0, поскольку слева нет элементов. Это также применимо к правому краю массива.
Верните самый левый индекс поворота. Если такого индекса не существует, верните -1.

Input: nums = [1,7,3,6,5,6]
Output: 3
Объяснение:
Индекс поворота равен 3.
Левая сумма = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11
Правая сумма = nums[4] + nums[5] = 5 + 6 = 11
*/

public class FindPivotIndex {
    public static void main(String[] args) {
        int[] nums = {1, 7, 3, 6, 5, 6};
        System.out.println(pivotIndex(nums)); // 3
    }

    // pivotIndex находит индекс поворота массива.
    // time: O(n), space: O(1)
    private static int pivotIndex(int[] nums) {
        int totalSum = 0; // Общая сумма всех элементов массива.

        for (int num : nums) { // Вычисляем общую сумму всех элементов.
            totalSum += num;
        }

        int leftSum = 0; // Сумма всех элементов слева от текущего элемента.

        for (int i = 0; i < nums.length; i++) {
            int num = nums[i];
            if (leftSum == totalSum - leftSum - num) {
                return i;
            }
            leftSum += num; // Обновляем левую сумму.
        }

        return -1; // Если не найден индекс поворота, возвращаем -1.
    }
}
