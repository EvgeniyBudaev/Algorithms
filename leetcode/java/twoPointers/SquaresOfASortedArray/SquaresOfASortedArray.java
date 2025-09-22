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
        int[] result = new int[n];
        int left = 0; // Индекс первого элемента массива nums
        int right = n - 1; // Индекс последнего элемента массива nums
        int index = n - 1; // Индекс последнего элемента массива result

        while (left <= right) {
            int leftSquare = nums[left] * nums[left];
            int rightSquare = nums[right] * nums[right];

            if (leftSquare > rightSquare) {
                result[index] = leftSquare;
                left++;
            } else {
                result[index] = rightSquare;
                right--;
            }
            index--;
        }

        return result;
    }

    // sortedSquares возвращает массив квадратов каждого числа, отсортированного в неубывающем порядке.
    // time: O(n), space: O(n)
//    private int[] sortedSquares(int[] nums) {
//        for (int i = 0; i < nums.length; i++) {
//            nums[i] = nums[i] * nums[i];
//        }
//
//        Arrays.sort(nums);
//        return nums;
//    }

    // sortedSquares возвращает массив квадратов каждого числа, отсортированного в неубывающем порядке.
    // time: O(n), space: O(n)
//    public int[] sortedSquares(int[] nums) {
//        return Arrays.stream(nums)
//                .map(num -> num * num)
//                .sorted()
//                .toArray();
//    }
}
