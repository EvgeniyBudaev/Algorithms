package array.FlattenDeeplyNestedArray;

import java.util.ArrayList;
import java.util.List;
import java.util.function.BiConsumer;

/* 2625. Flatten Deeply Nested Array
https://leetcode.com/problems/flatten-deeply-nested-array/description/

Учитывая многомерный массив arr и глубину n, верните сглаженную версию этого массива.
Многомерный массив — это рекурсивная структура данных, содержащая целые числа или другие многомерные массивы.
Сглаженный массив — это версия этого массива, в которой некоторые или все подмассивы удалены и заменены фактическими
элементами этого подмассива. Эту операцию выравнивания следует выполнять только в том случае, если текущая глубина
вложенности меньше n. Глубина элементов в первом массиве считается равной 0.
Пожалуйста, решите эту проблему без встроенного метода Array.flat.

Input: arr = [1, 2, 3, [4, 5, 6], [7, 8, [9, 10, 11], 12], [13, 14, 15]], n = 0
Output: [1, 2, 3, [4, 5, 6], [7, 8, [9, 10, 11], 12], [13, 14, 15]]
Объяснение:
Передача глубины n=0 всегда приводит к исходному массиву. Это связано с тем, что наименьшая возможная глубина подмассива
(0) не меньше n=0. Таким образом, ни один подмассив не должен быть сглажен.

Input: arr = [1, 2, 3, [4, 5, 6], [7, 8, [9, 10, 11], 12], [13, 14, 15]]
n = 1
Output: [1, 2, 3, 4, 5, 6, 7, 8, [9, 10, 11], 12, 13, 14, 15]
Объяснение:
Все подмассивы, начинающиеся с 4, 7 и 13, выравниваются. Это связано с тем, что их глубина 0 меньше 1.
Однако [9, 10, 11] остается несглаженной, поскольку ее глубина равна 1.

Input: arr = [[1, 2, 3], [4, 5, 6], [7, 8, [9, 10, 11], 12], [13, 14, 15]]
n = 2
Output: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15]
Объяснение:
Максимальная глубина любого подмассива равна 1. Таким образом, все они сглажены.
*/

public class FlattenDeeplyNestedArray {
    public static void main(String[] args) {
        List<Object> arr = List.of(
                1, 2, 3,
                new ArrayList<Object>(List.of(4, 5, 6)),
                new ArrayList<Object>(List.of(7, 8, new ArrayList<Object>(List.of(9, 10, 11)), 12)),
                new ArrayList<Object>(List.of(13, 14, 15))
        );

        int n = 0;
        System.out.println(flat(arr, n)); // [1, 2, 3, [4, 5, 6], [7, 8, [9, 10, 11], 12], [13, 14, 15]]
    }

    // flat - функция для сглаживания многомерного массива с заданной глубиной n.
    // Она использует рекурсивную функцию flatten для обхода массива и сглаживания его подмассивов.
    // time: O(n), space: O(n).
    private static List<Object> flat(List<Object> arr, int n) {
        List<Object> res = new ArrayList<>(); // Сглаженный результат

        // Используем BiConsumer для имитации замыкания как в Go
        BiConsumer<List<Object>, Integer> flatten = new BiConsumer<List<Object>, Integer>() {
            @Override
            public void accept(List<Object> currentArr, Integer depth) {
                for (Object item : currentArr) {
                    // Если элемент является массивом и глубина меньше n, рекурсивно сглаживаем
                    if (item instanceof List && depth < n) {
                        this.accept((List<Object>) item, depth + 1); // Увеличиваем глубину на 1
                    } else { // Иначе добавляем элемент в результат
                        res.add(item);
                    }
                }
            }
        };

        flatten.accept(arr, 0); // Начинаем обход массива с глубиной 0
        return res; // Возвращаем сглаженный результат
    }
}
