package array.FindTheDuplicateNumber;

/* 287. Find the Duplicate Number
https://leetcode.com/problems/find-the-duplicate-number/description/

Дан массив целых чисел nums, содержащий n + 1 целых чисел, где каждое целое число находится в диапазоне [1, n]
включительно.
В nums есть только одно повторяющееся число, верните это повторяющееся число.
Вы должны решить проблему, не изменяя числа массива и используя только постоянное дополнительное пространство.

Input: nums = [1,3,4,2,2]
Output: 2

Input: nums = [3,1,3,4,2]
Output: 3
*/

public class FindTheDuplicateNumber {
    public static void main(String[] args) {
        int[] nums = {1, 3, 4, 2, 2};
        System.out.println(findDuplicate(nums)); // 2
    }

    // findDuplicate возвращает целое повторяющееся число.
    // time: O(n), space: O(1)
    private static int findDuplicate(int[] nums) {
        for (int num : nums) {
            int i = Math.abs(num); // индекс в массиве
            if (nums[i] < 0) { // Если элемент уже посещался, значит это повторяющееся число.
                return i;
            } else { // Если элемент еще не посещался, сделаем его отрицательным.
                nums[i] = -nums[i];
            }
        }

        return 0; // Не найдено повторяющееся число.
    }
}
