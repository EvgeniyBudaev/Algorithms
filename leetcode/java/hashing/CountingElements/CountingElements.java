package leetcode.java.hashing.CountingElements;

import java.util.HashSet;
import java.util.Set;

/* 1426. Counting Elements

Учитывая целочисленный массив arr, подсчитайте, сколько элементов x существует так, что x + 1 также находится в arr.
Если в arr есть дубликаты, посчитайте их отдельно.

Input: arr = [1,2,3]
Output: 2
Пояснение: 1 и 2 учитываются, потому что 2 и 3 находятся в arr.

Input: arr = [1,1,3,3,5,5,7,7]
Output: 0
Пояснение: Никакие числа не учитываются, поскольку в arr нет цифр 2, 4, 6 или 8.

Input: arr = [1,3,2,3,5,0]
Output: 3
Пояснение: 0, 1 и 2 учитываются, потому что 1, 2 и 3 находятся в arr.

Input: arr = [1,1,2,2]
Output: 2
Пояснение: Две единицы засчитываются, потому что 2 находится в arr.

Input: arr = [1,1,2]
Output: 2
Объяснение: Обе единицы учитываются, поскольку в массиве есть 2.
*/

public class CountingElements {
    public static void main(String[] args) {
        System.out.println(countElements(new int[]{1, 2, 3})); // 2
        System.out.println(countElements(new int[]{1, 1, 2})); // 2 
    }

    // countElements подсчитывает количество элементов x в массиве, таких что x + 1 также находится в массиве.
    // time: O(n), space: O(n)
    private static int countElements(int[] arr) {
        // Множество для быстрой проверки наличия элементов в массиве.
        Set<Integer> elementSet = new HashSet<>();
        
        // Заполняем множество всеми уникальными элементами массива
        for (int num : arr) {
            elementSet.add(num);
        }

        int count = 0; // Счетчик для количества элементов x.
        // Проходим по всем элементам массива и проверяем, есть ли (x + 1) в множестве.
        for (int num : arr) {
            // Если num + 1 есть в множестве, то увеличиваем счетчик.
            if (elementSet.contains(num + 1)) {
                count++;
            }
        }

        return count; // Возвращаем счетчик.
    }
}
