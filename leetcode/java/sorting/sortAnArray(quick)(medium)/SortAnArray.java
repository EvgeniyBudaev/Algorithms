package leetcode.java.sorting.sortAnArray

(quick)(medium);

import java.util.Arrays;
import java.util.Random;

public class SortAnArray {
    private static final Random random = new Random();

    public static void main(String[] args) {
        int[] nums = {5, 2, 3, 1};
        System.out.println(Arrays.toString(sortArray(nums))); // [1, 2, 3, 5]
    }

    // sortArray - быстрая сортировка массива.
    // time: O(n log n), space: O(1)
    private static int[] sortArray(int[] nums) {
        qsort(nums, 0, nums.length);
        return nums;
    }

    // getPivot - выбирает случайный pivot в диапазоне [l, r)
    // time: O(1), space: O(1)
    private static int[] getPivot(int[] arr, int l, int r) {
        int pivotIdx = random.nextInt(r - l) + l;  // Генерируем случайный индекс в диапазоне [l, r)
        return new int[]{arr[pivotIdx], pivotIdx}; // Возвращаем pivot и его индекс
    }

    // partition - разделяет массив на две части относительно pivot
    // time: O(n), space: O(1)
    private static int partition(int[] arr, int lIdx, int rIdx) {
        int[] pivotData = getPivot(arr, lIdx, rIdx); // Получаем pivot и его индекс
        int pivot = pivotData[0];
        int pivotIdx = pivotData[1];

        // Перемещаем pivot в начало
        int temp = arr[lIdx];
        arr[lIdx] = arr[pivotIdx];
        arr[pivotIdx] = temp; // Меняем местами первый и pivot элементы

        int l = lIdx + 1;
        int r = rIdx - 1;     // Устанавливаем указатели на концы массива

        while (l <= r) {
            if (arr[l] < pivot) { // Если элемент слева меньше pivot, двигаем левый указатель
                l++;
            } else if (arr[r] > pivot) { // Если элемент справа больше pivot, двигаем правый указатель
                r--;
            } else { // Иначе меняем местами элементы и двигаем указатели
                // Меняем местами элементы
                temp = arr[l];
                arr[l] = arr[r];
                arr[r] = temp;
                l++;
                r--;
            }
        }

        // Возвращаем pivot на правильную позицию
        temp = arr[lIdx];
        arr[lIdx] = arr[r];
        arr[r] = temp;
        return r; // Возвращаем индекс pivot после сортировки
    }

    // qsort - рекурсивная функция быстрой сортировки
    // time: O(n log n), space: O(1)
    private static void qsort(int[] nums, int l, int r) {
        if (l >= r) {
            return;
        }
        int pivotIdx = partition(nums, l, r); // Получаем индекс pivot после сортировки
        qsort(nums, l, pivotIdx);          // Рекурсивно сортируем левую часть массива
        qsort(nums, pivotIdx + 1, r);      // Рекурсивно сортируем правую часть массива
    }
}
