package binarySearch.FindKClosestElements;

import java.util.ArrayList;
import java.util.List;

/* 658. Find K Closest Elements
https://leetcode.com/problems/find-k-closest-elements/description/

Учитывая отсортированный массив целых чисел arr, два целых числа k и x, верните k ближайших к x целых чисел в массиве.
Результат также следует отсортировать по возрастанию.

Целое число a ближе к x, чем целое число b, если:
|a - x| < |b - x|, or
|a - x| == |b - x| and a < b

Input: arr = [1,2,3,4,5], k = 4, x = 3
Output: [1,2,3,4]

Input: arr = [1,2,3,4,5], k = 4, x = -1
Output: [1,2,3,4]
*/

public class FindKClosestElements {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4, 5};
        System.out.println(findClosestElements(arr, 4, 3)); // [1, 2, 3, 4]
    }

    // findClosestElements возвращает k ближайших к x целых чисел в массиве.
    // time: O(log(n-k) + k), space: O(1), где n - длина массива arr, k - количество ближайших чисел.
    private static List<Integer> findClosestElements(int[] arr, int k, int x) {
        int n = arr.length;
        int left = 0;
        int right = n - k;

        while (left < right) {
            int mid = left + (right - left) / 2; // Поиск середины между левым и правым указателями
            // Сравнение расстояния между элементом в середине и x с расстоянием между элементом, находящимся на k позициях
            // Если расстояние до левого элемента больше, чем до правого элемента, двигаем левую границу
            if (x - arr[mid] > arr[mid + k] - x) {
                left = mid + 1;
            } else {
                right = mid;
            }
        }

        // Формируем результат из k элементов, начиная с `left`
        List<Integer> result = new ArrayList<>(k);
        for (int i = left; i < left + k; i++) {
            result.add(arr[i]);
        }

        // Возвращаем срез, начиная с позиции, найденной выше, и длиной k.
        return result;
    }
}
