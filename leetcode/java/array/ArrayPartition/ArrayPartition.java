package array.ArrayPartition;

import java.util.Arrays;

/* 561. Array Partition
https://leetcode.com/problems/array-partition/description/

Дан массив целых чисел nums из 2n целых чисел, сгруппируйте эти целые числа в n пар (a1, b1), (a2, b2), ..., (an, bn)
так, чтобы сумма min(ai, bi) для всех i была максимальной. Верните максимальную сумму.

Input: nums = [1,4,3,2]
Output: 4
Пояснение: Все возможные пары (без учета порядка элементов):
1. (1, 4), (2, 3) -> min(1, 4) + min(2, 3) = 1 + 2 = 3
2. (1, 3), (2, 4) -> min(1, 3) + min(2, 4) = 1 + 2 = 3
3. (1, 2), (3, 4) -> min(1, 2) + min(3, 4) = 1 + 3 = 4
Таким образом, максимально возможная сумма — 4.
*/

public class ArrayPartition {
    public static void main(String[] args) {
        int[] nums = {1, 4, 3, 2};
        System.out.println(arrayPairSum(nums)); // 4
    }

    // arrayPairSum возвращает максимальную сумму пар в массиве чисел.
    // time: O(n), space: O(1)
    private static int arrayPairSum(int[] nums) {
        Arrays.sort(nums); // Сортируем массив O(n*log(n)) [1, 2, 3, 4]
        int sum = 0; // Сумма пар

        // Берем каждый второй элемент (начиная с 0). time: O(n), space: O(1)
        for (int i = 0; i < nums.length; i += 2) {
            sum += nums[i]; // Добавляем его к сумме пар
        }

        return sum; // Возвращаем сумму пар
    }
}
