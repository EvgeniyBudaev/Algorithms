package binarySearch.SearchInRotatedSortedArray;

/* 33. Search in Rotated Sorted Array
https://leetcode.com/problems/search-in-rotated-sorted-array/description/

Существует целочисленный массив nums, отсортированный по возрастанию (с разными значениями).

Перед передачей в вашу функцию nums, возможно, поворачивается с неизвестным индексом поворота k (1 <= k < nums.length),
так что результирующий массив имеет вид [nums[k], nums[k+1],... , nums[n-1], nums[0], nums[1], ..., nums[k-1]]
(с индексом 0). Например, [0,1,2,4,5,6,7] можно повернуть с опорным индексом 3 и превратить в [4,5,6,7,0,1,2].

Учитывая массив nums после возможного поворота и целочисленную цель, верните индекс цели, если он находится в nums, или
-1, если он не в nums.

Вы должны написать алгоритм с компиляцией времени выполнения O(log n).

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
*/

public class SearchInRotatedSortedArray {
    public static void main(String[] args) {
        int[] nums = {4, 5, 6, 7, 0, 1, 2};
        System.out.println(search(nums, 0)); // 4
    }

    // search ищет индекс цели в отсортированном массиве после возможного поворота.
    // time: O(log n), space: O(1)
    private static int search(int[] nums, int target) {
        int low = 0;
        int high = nums.length - 1;

        while (low <= high) {
            int mid = (low + high) / 2;
            int guess = nums[mid];
            int leftNum = nums[low];
            int rightNum = nums[high];

            if (guess == target) {
                return mid;
            }

            // Если левая половина отсортирована
            if (leftNum <= guess) {
                // Если цель находится между левой и центральной числами
                if (guess > target && leftNum <= target) {
                    high = mid - 1;
                } else {
                    low = mid + 1;
                }
            } else {
                // Правая половина отсортирована
                // Если цель находится между правой и центральной числами
                if (guess < target && rightNum >= target) {
                    low = mid + 1;
                } else {
                    high = mid - 1;
                }
            }
        }

        return -1; 
    }
}
