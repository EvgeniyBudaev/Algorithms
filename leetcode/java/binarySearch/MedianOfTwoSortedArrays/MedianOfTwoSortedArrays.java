package binarySearch.MedianOfTwoSortedArrays;

/* 4. Median of Two Sorted Arrays
https://leetcode.com/problems/median-of-two-sorted-arrays/description/
solution https://leetcode.com/problems/median-of-two-sorted-arrays/solutions/4070371/94-96-binary-search-two-pointers/

Учитывая два отсортированных массива nums1 и nums2 размера m и n соответственно, верните медиану двух отсортированных
массивов.
Общая сложность времени выполнения должна составлять O(log (m+n)).

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
*/

public class MedianOfTwoSortedArrays {
    public static void main(String[] args) {
        int[] nums1 = {1, 3};
        int[] nums2 = {2};
        System.out.println(findMedianSortedArrays(nums1, nums2)); // 2
    }

    private static double findMedianSortedArrays(int[] nums1, int[] nums2) {
        // Обеспечиваем, чтобы nums1 был более коротким массивом
        if (nums1.length > nums2.length) {
            int[] temp = nums1;
            nums1 = nums2;
            nums2 = temp;
        }

        int m = nums1.length;
        int n = nums2.length;
        int low = 0;
        int high = m;

        while (low <= high) {
            int partitionX = (low + high) / 2;
            int partitionY = (m + n + 1) / 2 - partitionX;

            // Определяем max и min для каждого раздела
            int maxX = (partitionX == 0) ? Integer.MIN_VALUE : nums1[partitionX - 1];
            int maxY = (partitionY == 0) ? Integer.MIN_VALUE : nums2[partitionY - 1];

            int minX = (partitionX == m) ? Integer.MAX_VALUE : nums1[partitionX];
            int minY = (partitionY == n) ? Integer.MAX_VALUE : nums2[partitionY];

            // Проверяем условие нахождения медианы
            if (maxX <= minY && maxY <= minX) {
                if ((m + n) % 2 == 0) {
                    return (Math.max(maxX, maxY) + Math.min(minX, minY)) / 2.0;
                } else {
                    return Math.max(maxX, maxY);
                }
            } else if (maxX > minY) {
                high = partitionX - 1;
            } else {
                low = partitionX + 1;
            }
        }

        throw new IllegalArgumentException("Input arrays are not sorted");
    }
}
