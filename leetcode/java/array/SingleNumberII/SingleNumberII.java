package array.SingleNumberII;

import java.util.Arrays;

/* 137. Single Number II
https://leetcode.com/problems/single-number-ii/description/

Дан целочисленный массив nums, где каждый элемент встречается три раза, за исключением одного, который встречается ровно
один раз. Найдите единственный элемент и верните его.
Вам необходимо реализовать решение с линейной сложностью времени выполнения и использовать только постоянное
дополнительное пространство.

Input: nums = [2,2,3,2]
Output: 3
*/

public class SingleNumberII {
    public static void main(String[] args) {
        int nums[] = {2, 2, 3, 2};
        System.out.println(singleNumber(nums)); // 3
    }

    // singleNumber находит уникальный элемент в массиве.
    // time: O(n logn), space: O(1)
    private static int singleNumber(int[] nums) {
        if (nums.length == 1) { // Если массив состоит из одного элемента, то он и является единственным
            return nums[0];
        }

        Arrays.sort(nums); // Сортируем массив

        if (nums[0] != nums[1]) { // Если первый элемент не равен второму, значит он и является единственным
            return nums[0];
        }

        // Если последний элемент не равен предпоследнему, значит он и является единственным
        if (nums[nums.length - 1] != nums[nums.length - 2]) {
            return nums[nums.length - 1];
        }

        for (int i = 1; i < nums.length - 1; i++) { // Проходимся по массиву от второго до предпоследнего элемента
            // Если текущий элемент не равен предыдущему и следующему, значит он и является единственным
            if (nums[i] != nums[i - 1] && nums[i] != nums[i + 1]) {
                return nums[i];
            }
        }

        return -1; // Если не найден, возвращаем -1
    }
}
