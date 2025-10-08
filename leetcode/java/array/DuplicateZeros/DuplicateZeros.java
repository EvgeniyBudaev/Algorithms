package array.DuplicateZeros;

/* 1089. Duplicate Zeros
https://leetcode.com/problems/duplicate-zeros/description/

Учитывая целочисленный массив фиксированной длины arr, дублируйте каждое вхождение нуля, сдвигая оставшиеся элементы
вправо. Обратите внимание, что элементы, превышающие длину исходного массива, не записываются. Внесите вышеуказанные
изменения во входной массив и ничего не возвращайте.

Input: arr = [1,0,2,3,0,4,5,0]
Output: [1,0,0,2,3,0,0,4]
Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]

Input: arr = [1,2,3]
Output: [1,2,3]
Explanation: After calling your function, the input array is modified to: [1,2,3]
*/

public class DuplicateZeros {
    public static void main(String[] args) {
        int[] arr = {1, 0, 2, 3, 0, 4, 5, 0};
        duplicateZeros(arr);
        System.out.println(arr); // [1,0,0,2,3,0,0,4]
    }

    // duplicateZeros дублирует каждое вхождение нуля, сдвигая оставшиеся элементы вправо.
    // Элементы, превышающие длину исходного массива, не записываются.
    // time: O(n), space: O(1)
    private static void duplicateZeros(int[] arr) {
        for (int i = 0; i < arr.length; i++) { // Проходим по массиву
            if (arr[i] == 0 && arr.length != 0) { // Если ноль и длина массива не равна 0
                System.arraycopy(arr, i, arr, i + 1, arr.length - i - 1); // Копируем элементы вправо
                i++; // Увеличиваем индекс на 1
            }
        }
    }
}
