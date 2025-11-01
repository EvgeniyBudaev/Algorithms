package array.MonotonicArray;

/* 896. Monotonic Array
https://leetcode.com/problems/monotonic-array/description/

Массив монотонен, если он либо монотонно возрастает, либо монотонно убывает.
Массив nums монотонно возрастает, если для всех i <= j, nums[i] <= nums[j]. Массив nums монотонно убывает,
если для всех i <= j, nums[i] >= nums[j].
Если задан целочисленный массив nums, вернуть true, если заданный массив монотонен, или false в противном случае.

Input: nums = [1,2,2,3]
Output: true

Input: nums = [6,5,4,4]
Output: true

Input: nums = [1,3,2]
Output: false
*/

public class MonotonicArray {
    public static void main(String[] args) {
        int[] nums1 = {1, 2, 2, 3};
        System.out.println(isMonotonic(nums1)); // true

        int[] nums2 = {6, 5, 4, 4};
        System.out.println(isMonotonic(nums2)); // true

        int[] nums3 = {1, 3, 2};
        System.out.println(isMonotonic(nums3)); // false
    }

    private static boolean isMonotonic(int[] nums) {
        boolean isIncreasing = true; // монотонно возрастает
        boolean isDecreasing = true; // монотонно убывает

        for (int i = 1; i < nums.length; i++) {
            if (nums[i] > nums[i - 1]) { // если монотонно возрастает
                isDecreasing = false;
            } else if (nums[i] < nums[i - 1]) { // если монотонно убывает
                isIncreasing = false;
            }
            // если не монотонно возрастает и не монотонно убывает
            if (!isIncreasing && !isDecreasing) {
                return false;
            }
        }

        return true; // если монотонно возрастает или монотонно убывает
    }
}
