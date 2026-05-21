package leetcode.java.sorting.sortAnArray

import java.util.Arrays;

public class SortAnArray {
    public static void main(String[] args) {
        int[] nums = {5, 2, 3, 1};
        System.out.println(Arrays.toString(sortArray(nums))); // [1, 2, 3, 5]
    }

    // sortArray - сортировка массива пузырьком.
    // time: O(n^2), space: O(1)
    private static int[] sortArray(int[] nums) {
        return bubbleSort(nums); // Вызываем функцию сортировки пузырьком
    }

    // bubbleSort - сортировка пузырьком.
    // time: O(n^2), space: O(1)
    private static int[] bubbleSort(int[] arr) {
        int n = arr.length; // Длина массива

        for (int i = 0; i < n - 1; i++) {
            boolean swapped = false; // Флаг для отслеживания перестановок
            for (int j = 0; j < n - i - 1; j++) { // Проходим по неотсортированной части массива
                if (arr[j] > arr[j + 1]) { // Если текущий элемент больше следующего,
                    // Меняем местами соседние элементы
                    int temp = arr[j];
                    arr[j] = arr[j + 1];
                    arr[j + 1] = temp; // Перестановка элементов
                    swapped = true; // Устанавливаем флаг перестановки
                }
            }
            // Если за проход не было перестановок, массив уже отсортирован
            if (!swapped) {
                break;
            }
        }

        return arr;
    }
}
