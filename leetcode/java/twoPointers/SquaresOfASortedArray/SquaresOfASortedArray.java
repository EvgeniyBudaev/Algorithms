package twoPointers.SquaresOfASortedArray;

import java.util.Arrays;

/* 977. Squares of a Sorted Array
https://leetcode.com/problems/squares-of-a-sorted-array/description/

Учитывая целочисленный массив nums, отсортированный в неубывающем порядке, верните массив квадратов каждого числа,
отсортированного в неубывающем порядке.

Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].

Input: nums = [-7,-3,2,3,11]
Output: [4,9,9,49,121]

Two pointers
Time complexity: O(n)
Space complexity: O(n)
*/

public class SquaresOfASortedArray {
    public static void main(String[] args) {
        int[] nums = new int[]{-4, -1, 0, 3, 10};
        // Создаем экземпляр класса для вызова нестатического метода
        SquaresOfASortedArray solution = new SquaresOfASortedArray();
        System.out.println(Arrays.toString(solution.sortedSquares(nums))); // [0,1,9,16,100]
    }

    // sortedSquares возвращает массив квадратов каждого числа, отсортированного в неубывающем порядке.
    // time: O(n), space: O(n)
    private int[] sortedSquares(int[] nums) {
        int n = nums.length; // Длина массива
        int[] result = new int[n];     // Массив квадратов
        int p1 = 0, p2 = n - 1;       // Указатели на начало и конец массива

        for (int i = n - 1; p1 <= p2; i--) { // Проход по массиву
            if (abs(nums[p1]) > abs(nums[p2])) { // Если модуль первого элемента больше модуля второго
                result[i] = nums[p1] * nums[p1]; // Записываем квадрат первого элемента
                p1++;
            } else {
                result[i] = nums[p2] * nums[p2];
                p2--;
            }
        }

        return result;
    }

    // abs возвращает абсолютное значение числа.
    // time: O(1), space: O(1)
    private static int abs(int x) {
        if (x < 0) {
		    return -x;
        }
        return x;
    }
}
