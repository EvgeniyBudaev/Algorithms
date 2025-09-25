package twoPointers.IntersectionOfTwoArraysII;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/* 350. Intersection of Two Arrays II
https://leetcode.com/problems/intersection-of-two-arrays-ii/description/

Учитывая два целочисленных массива nums1 и nums2, верните массив их пересечения. Каждый элемент результата должен
появляться столько раз, сколько он отображается в обоих массивах, и вы можете возвращать результат в любом порядке.

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
Пояснение: [9,4] также принимается.
*/

public class IntersectionOfTwoArraysII {
    public static void main(String[] args) {
        int[] nums1 = {1, 2, 2, 1};
        int[] nums2 = {2, 2};
        System.out.println(Arrays.toString(intersect(nums1, nums2))); // [2, 2]
    }
    
    // intersect находит пересечение двух массивов nums1 и nums2.
    // time: O(n), space: O(n)
    private static int[] intersect(int[] nums1, int[] nums2) {
        Arrays.sort(nums1); // [1, 1, 2, 2]
        Arrays.sort(nums2); // [2, 2]
        
        int left = 0, right = 0;
        List<Integer> result = new ArrayList<>();
        
        // time: O(n), space: O(n)
        while (left < nums1.length && right < nums2.length) {
            if (nums1[left] < nums2[right]) {
                left++;
            } else if (nums1[left] > nums2[right]) {
                right++;
            } else {
                result.add(nums1[left]);
                left++;
                right++;
            }
        }
        
        // Конвертируем List в массив
        int[] arr = new int[result.size()];
        for (int i = 0; i < result.size(); i++) {
            arr[i] = result.get(i);
        }

        return arr;
    }
}
