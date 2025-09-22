package twoPointers.MoveZeroes;

import java.util.Arrays;

/* 283. Move Zeroes
https://leetcode.com/problems/move-zeroes/description/

Учитывая целочисленный массив nums, переместите все 0 в его конец, сохраняя относительный порядок ненулевых элементов.
Обратите внимание, что вы должны сделать это на месте, не копируя массив.

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

Input: nums = [0]
Output: [0]
*/

public class MoveZeroes {
    public static void main(String[] args) {
        int[] nums = {0, 1, 0, 3, 12};
        moveZeroes(nums);
        System.out.println(Arrays.toString(nums)); // [1, 3, 12, 0, 0]
    }

    // moveZeroes перемещает все нули в конец массива nums, сохраняя относительный порядок ненулевых элементов.
    // time: O(n), space: O(1)
    private static void moveZeroes(int[] nums) {
        int left = 0; // Индекс для следующего ненулевого элемента

        for (int right = 0; right < nums.length; right++) {
            if (nums[right] != 0) { // Если текущий элемент не равен 0
                // Меняем местами текущий элемент и следующий ненулевой элемент
                int temp = nums[right];
                nums[right] = nums[left];
                nums[left] = temp;
                left++; // Увеличиваем индекс для следующего ненулевого элемента
            }
        }
    }
}
