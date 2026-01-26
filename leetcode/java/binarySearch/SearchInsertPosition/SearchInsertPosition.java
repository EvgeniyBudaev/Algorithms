package binarySearch.SearchInsertPosition;

/* 35. Search Insert Position
https://leetcode.com/problems/search-insert-position/description/

Учитывая отсортированный массив различных целых чисел и целевое значение, верните индекс, если цель найдена.
Если нет, верните индекс там, где он был бы, если бы он был вставлен по порядку.
Вы должны написать алгоритм со сложностью выполнения O(log n)

Input: nums = [1,3,5,6], target = 5
Output: 2

Input: nums = [1,3,5,6], target = 2
Output: 1

Input: nums = [1,3,5,6], target = 7
Output: 4
*/

public class SearchInsertPosition {
    public static void main(String[] args) {
        int[] nums = {1, 3, 5, 6};
        System.out.println(searchInsert(nums, 5)); // 2
    }

    // searchInsert возвращает индекс элемента массива или индекс, где он должен быть вставлен.
    // time: O(log n), space: O(1)
    private static int searchInsert(int[] nums, int target) {
        int low = 0;
        int high = nums.length - 1;

        while (low <= high) {
            int mid = (low + high) / 2;
            int guess = nums[mid];

            if (guess == target) {
                return mid;
            }
            if (guess > target) {
                high = mid - 1;
            } else {
                low = mid + 1;
            }
        }

        // Цель не найдена в массиве, возвращаем индекс, где она должна быть вставлена.
        return low;
    }
}
