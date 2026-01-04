package binarySearch.FindMinimumInRotatedSortedArray;

/* 153. Find Minimum in Rotated Sorted Array
https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/

Предположим, что массив длины n, отсортированный в порядке возрастания, поворачивается от 1 до n раз. Например, массив
nums = [0,1,2,4,5,6,7] может выглядеть следующим образом:

[4,5,6,7,0,1,2], если его повернули 4 раза.
[0,1,2,4,5,6,7], если его повернули 7 раз.
Обратите внимание, что вращение массива [a[0], a[1], a[2], ..., a[n-1]] 1 раз приводит к созданию массива
[a[n-1], a[0] , а[1], а[2], ..., а[n-2]].

Учитывая отсортированные числа уникальных элементов повернутого массива, верните минимальный элемент этого массива.
Вы должны написать алгоритм, который работает за время O(log n).

Input: nums = [3,4,5,1,2]
Output: 1
Пояснение: Исходный массив был [1,2,3,4,5] повёрнут 3 раза.
*/

public class FindMinimumInRotatedSortedArray {
    public static void main(String[] args) {
        int[] nums = {3, 4, 5, 1, 2};
        System.out.println(findMin(nums)); // 1
    }

    // findMin возвращает минимальный элемент повернутого отсортированного массива.
    // time: O(log n), space: O(1)
    private static int findMin(int[] nums) {
        int left = 0, right = nums.length - 1;

        
        while (left < right) {
            int mid = left + (right - left) / 2;
            int guess = nums[mid];
            int leftNum = nums[left];
            int rightNum = nums[right];
            
            // Если левый элемент меньше правого элемента, то массив отсортирован и минимальный элемент левый.
            if (leftNum < rightNum) {
                return leftNum;
            }
            
            if (leftNum <= guess) {
                left = mid + 1;
            } else {
                right = mid;
            }
        }
        
        return nums[left]; // Возвращаем элемент в левой части массива
    }
}
