package binarySearch.PeakIndexInAMountainArray;

/* 852. Peak Index in a Mountain Array
https://leetcode.com/problems/peak-index-in-a-mountain-array/description/

Вам дан целочисленный горный массив длиной n, значения которого увеличиваются до пикового элемента, а затем уменьшаются.
Верните индекс пикового элемента.
Ваша задача — решить ее за время сложности O(log(n)).

Input: arr = [0,1,0]
Output: 1

Input: arr = [0,2,1,0]
Output: 1

Input: arr = [0,10,5,2]
Output: 1
*/

public class PeakIndexInAMountainArray {
    public static void main(String[] args) {
        int[] arr = {0, 1, 0};
        System.out.println(peakIndexInMountainArray(arr)); // 1
    }

    // peakIndexInMountainArray возвращает индекс пикового элемента в массиве arr.
    // time: O(log n), space: O(1)
    private static int peakIndexInMountainArray(int[] arr) {
        int left = 0;
        int right = arr.length - 1;

        while (left <= right) {
            int mid = (left + right) / 2; // Находим индекс среднего элемента
            // Проверяем, является ли текущий элемент пиковым элементом
            if (arr[mid] < arr[mid + 1]) { // Если текущий элемент меньше следующего, двигаемся вправо
                left = mid + 1;
            } else { // Если текущий элемент больше следующего, двигаемся влево
                right = mid - 1;
            }
        }

        // Возвращаем индекс пикового элемента
        return left;
    }
}
