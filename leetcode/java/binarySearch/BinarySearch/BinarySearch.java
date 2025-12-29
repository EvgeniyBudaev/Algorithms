package binarySearch.BinarySearch;

/* 704. Binary Search
https://leetcode.com/problems/binary-search/description/

Учитывая массив целых чисел nums, отсортированный в порядке возрастания, и целочисленную цель, напишите функцию для
поиска цели в nums. Если цель существует, верните ее индекс. В противном случае верните -1.
Вы должны написать алгоритм со сложностью выполнения O(log n).

Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4

Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1
*/

public class BinarySearch {
    public static void main(String[] args) {
        int[] nums = {-1, 0, 3, 5, 9, 12};
        System.out.println(search(nums, 9)); // 4 
    }

    // search возвращает индекс элемента массива или -1, если элемент отсутствует в массиве.
    // time: O(log n), space: O(1)
    private static int search(int[] nums, int target) {
        int low = 0;
        int high = nums.length - 1;

        while (low <= high) {
            int mid = (low + high) / 2; // mid - индекс среднего элемента
            int guess = nums[mid];      // guess - предполагаемое значение

            if (guess == target) {
                return mid;
            }
            if (guess > target) {
                high = mid - 1;
            } else {
                low = mid + 1;
            }
        }

        return -1;
    }
}
