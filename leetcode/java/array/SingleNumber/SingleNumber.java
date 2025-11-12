package array.SingleNumber;

import java.util.Arrays;

/* 136. Single Number
https://leetcode.com/problems/single-number/description/

Учитывая непустой массив целых чисел nums, каждый элемент появляется дважды, кроме одного. Найдите этого единственного.
Вы должны реализовать решение с линейной сложностью времени выполнения и использовать только постоянное дополнительное
пространство.

Input: nums = [2,2,1]
Output: 1

Input: nums = [4,1,2,1,2]
Output: 4

Input: nums = [1]
Output: 1
*/

public class SingleNumber {
    public static void main(String[] args) {
        int[] nums = {2, 2, 1};
        System.out.println(singleNumber(nums)); // 1
    }

    // singleNumber находит уникальное число в слайсе.
    // time: O(n logn), space: O(1)
    private static int singleNumber(int[] nums) {
        Arrays.sort(nums); // time: O(n logn), space: O(1)

        // Цикл по всем элементам, кроме последнего. time: O(n), space: O(1)
        for (int i = 0; i < nums.length - 1; i += 2) {
            // Если текущий элемент не равен следующему, то текущий элемент является уникальным.
            if (nums[i] != nums[i + 1]) {
                return nums[i];
            }
        }

        // Если все элементы встречаются дважды, то последний элемент будет уникальным.
        return nums[nums.length - 1];
    }
}
