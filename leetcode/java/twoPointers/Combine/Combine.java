package twoPointers.Combine;

import java.util.Arrays;

/*
Учитывая два отсортированных массива целых чисел arr1 и arr2,
верните новый массив, который объединяет их оба и также отсортирован.
*/

public class Combine {
    public static void main(String[] args) {
        int[] arr1 = {1, 4, 7, 20};
        int[] arr2 = {3, 5, 6};
        System.out.println(Arrays.toString(combine(arr1, arr2))); // [1, 3, 4, 5, 6, 7, 20]
    }

    /**
     * Объединяет два отсортированных массива arr1 и arr2 в один отсортированный массив.
     * time: O(n+m), т.к. мы проходим по двум массивам один раз
     * space: O(n+m), т.к. мы создаем новый массив с размером n+m
     */
    private static int[] combine(int[] arr1, int[] arr2) {
        int m = arr1.length, n = arr2.length; // Размеры массивов
        int p1 = m - 1, p2 = n - 1; // Указатели на последние элементы массивов

        // Создаем новый массив размером m + n
        int[] result = new int[m + n];

        // Копируем элементы из arr1 в result
        System.arraycopy(arr1, 0, result, 0, m);

        // Проходим по result с конца
        for (int i = m + n - 1; i >= 0; i--) {
            // Если все элементы из arr2 были добавлены, завершаем цикл
            if (p2 < 0) {
                break;
            }
            // Сравниваем последний элемент каждого массива и добавляем больший элемент в конец
            if (p1 >= 0 && result[p1] > arr2[p2]) {
                result[i] = result[p1];
                p1--;
            } else {
                result[i] = arr2[p2];
                p2--;
            }
        }

        return result;
    }

    /**
     * Объединяет два отсортированных массива arr1 и arr2 в один отсортированный массив.
     * time: O(n+m), т.к. мы проходим по двум массивам один раз
     * space: O(n+m), т.к. мы создаем новый массив с размером n+m
     */
    // public static int[] combine(int[] arr1, int[] arr2) {
    //     return IntStream.concat(Arrays.stream(arr1), Arrays.stream(arr2))
    //             .sorted()
    //             .toArray();
    // }
}
