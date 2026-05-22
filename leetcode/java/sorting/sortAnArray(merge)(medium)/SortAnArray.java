package leetcode.java.sorting.sortAnArray

(merge)(medium);

import java.util.Arrays;

public class SortAnArray {
    public static void main(String[] args) {
        int[] nums = {5, 2, 3, 1};
        System.out.println(Arrays.toString(sortArray(nums))); // [1, 2, 3, 5] 
    }

    // sortArray - сортировка массива слиянием.
    // time: O(n log(n)), space: O(n)
    private static int[] sortArray(int[] nums) {
        if (nums.length <= 1) {
            return nums;
        }

        int mid = nums.length / 2; // разделение на две половины
        int[] leftPart = sortArray(Arrays.copyOfRange(nums, 0, mid));    // сортировка левой половины
        int[] rightPart = sortArray(Arrays.copyOfRange(nums, mid, nums.length)); // сортировка правой половины
        return merge(leftPart, rightPart);            // слияние отсортированных половин
    }

    // merge - функция слияния двух отсортированных массивов.
    // time: O(n), space: O(n)
    private static int[] merge(int[] arr1, int[] arr2) {
        int[] result = new int[arr1.length + arr2.length]; // объединение двух массивов в один
        int p1 = 0, p2 = 0;                                // указатели на текущие элементы в обоих массивах
        int k = 0;                                         // указатель для результирующего массива

        // слияние массивов, пока один из указателей не достигнет конца своего массива
        while (p1 < arr1.length || p2 < arr2.length) {
            // если текущий элемент первого массива меньше или равен текущему элементу второго массива,
            // добавляем его в объединенный массив и сдвигаем указатель на первый массив
            if (p2 >= arr2.length || (p1 < arr1.length && arr1[p1] <= arr2[p2])) {
                result[k++] = arr1[p1++]; // добавление текущего элемента первого массива в объединенный массив
            } else {
                result[k++] = arr2[p2++]; // добавление текущего элемента второго массива в объединенный массив
            }
        }

        // возвращаем объединенный отсортированный массив
        return result;
    }
}
