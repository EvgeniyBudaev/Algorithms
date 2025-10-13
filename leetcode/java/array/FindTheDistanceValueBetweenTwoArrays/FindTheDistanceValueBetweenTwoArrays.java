package array.FindTheDistanceValueBetweenTwoArrays;

import java.util.Arrays;

/* 1385. Find the Distance Value Between Two Arrays
https://leetcode.com/problems/find-the-distance-value-between-two-arrays/description/

Даны два целочисленных массива arr1 и arr2, и целое число d, вернуть значение расстояния между двумя массивами.
Значение расстояния определяется как количество элементов arr1[i], таких, что нет ни одного элемента arr2[j],
где |arr1[i]-arr2[j]| <= d.

Input: arr1 = [4,5,8], arr2 = [10,9,1,8], d = 2
Output: 2
Explanation:
For arr1[0]=4 we have:
|4-10|=6 > d=2
|4-9|=5 > d=2
|4-1|=3 > d=2
|4-8|=4 > d=2
For arr1[1]=5 we have:
|5-10|=5 > d=2
|5-9|=4 > d=2
|5-1|=4 > d=2
|5-8|=3 > d=2
For arr1[2]=8 we have:
|8-10|=2 <= d=2
|8-9|=1 <= d=2
|8-1|=7 > d=2
|8-8|=0 <= d=2
*/

public class FindTheDistanceValueBetweenTwoArrays {
    public static void main(String[] args) {
        int[] arr1 = {4, 5, 8};
        int[] arr2 = {10, 9, 1, 8};
        int d = 2; // d - дистанция
        System.out.println(findTheDistanceValue(arr1, arr2, d)); // 2
    }

    // findTheDistanceValue - возвращает значение расстояния между двумя массивами.
    // time: O(n log n), space: O(1)
    private static int findTheDistanceValue(int[] arr1, int[] arr2, int d) {
        Arrays.sort(arr2); // Сортируем arr2
        int count = 0; // Инициализируем счетчик подходящих элементов

        // Проходим по каждому элементу x в arr1
        for (int x : arr1) {
            // Используем бинарный поиск для поиска ближайшего элемента в arr2
            int index = Arrays.binarySearch(arr2, x);

            if (index < 0) {
                index = -index - 1;
            }

            boolean valid = true; // Флаг, указывающий на то, что элемент подходит

            // Проверяем ближайшие элементы в отсортированном массиве
            // Проверяем следующий элемент
            if (index < arr2.length && Math.abs(x - arr2[index]) <= d) {
                valid = false;
            }

            // Проверяем предыдущий элемент
            if (index > 0 && Math.abs(x - arr2[index - 1]) <= d) {
                valid = false;
            }

            // Если элемент x подходит, увеличиваем счетчик
            if (valid) {
                count++;
            }
        }

        return count; // Возвращаем значение счетчика
    }
}
