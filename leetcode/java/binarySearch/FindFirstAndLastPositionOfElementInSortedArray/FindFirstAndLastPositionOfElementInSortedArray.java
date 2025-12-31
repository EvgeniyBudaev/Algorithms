package binarySearch.FindFirstAndLastPositionOfElementInSortedArray;

import java.util.Arrays;

/* 34. Find First and Last Position of Element in Sorted Array
https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/

Дан массив целых чисел nums, отсортированный в неубывающем порядке, найти начальную и конечную позицию заданного
целевого значения.
Если target не найден в массиве, вернуть [-1, -1].
Необходимо написать алгоритм со сложностью выполнения O(log n).

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
*/

public class FindFirstAndLastPositionOfElementInSortedArray {
    public static void main(String[] args) {
        int[] arr = {5, 7, 7, 8, 8, 10};
        int target = 8;
        System.out.println(Arrays.toString(searchRange(arr, target))); // [3, 4]

        int[] arr2 = {0, 0, 1, 2, 3, 3, 4};
        int target2 = 2;
        System.out.println(Arrays.toString(searchRange(arr2, target2))); // [3, 3]
    }

    // searchRange находит начальную и конечную позицию целевого значения в массиве.
    // time: O(log n), space: O(1)
    private static int[] searchRange(int[] nums, int target) {
        if (nums == null || nums.length == 0) { // Проверка на пустой массив
            return new int[]{-1, -1};
        }

        int left = findFirst(nums, target);
        if (left == -1) {
            return new int[]{-1, -1};
        }
        
        int right = findLast(nums, target);
        return new int[]{left, right}; // Возврат начальной и конечной позиций
    }

    private static int findFirst(int[] nums, int target) {
        int low = 0, high = nums.length - 1;
        int result = -1;
        while (low <= high) {
            int mid = low + (high - low) / 2;
            if (nums[mid] == target) {
                result = mid;
                high = mid - 1; // продолжаем искать влево
            } else if (nums[mid] < target) {
                low = mid + 1;
            } else {
                high = mid - 1;
            }
        }
        return result;
    }

    private static int findLast(int[] nums, int target) {
        int low = 0, high = nums.length - 1;
        int result = -1;
        while (low <= high) {
            int mid = low + (high - low) / 2;
            if (nums[mid] == target) {
                result = mid;
                low = mid + 1; // продолжаем искать вправо
            } else if (nums[mid] < target) {
                low = mid + 1;
            } else {
                high = mid - 1;
            }
        }
        return result;
    }
}
