package array.FindThePrefixCommonArrayOfTwoArrays;

import java.util.Arrays;
import java.util.HashSet;
import java.util.Set;

/* 2657. Find the Prefix Common Array of Two Arrays
https://leetcode.com/problems/find-the-prefix-common-array-of-two-arrays/description/

Вам даны две целочисленные перестановки A и B с нулевым индексом длины n.
Общий префиксный массив A и B — это массив C, такой, что C[i] равен количеству чисел, которые присутствуют в индексе i
или перед ним как в A, так и в B.
Верните общий массив префиксов A и B.
Последовательность из n целых чисел называется перестановкой, если она содержит все целые числа от 1 до n ровно один раз.

Input: A = [1,3,2,4], B = [3,1,2,4]
Output: [0,2,3,4]
Объяснение: При i = 0: нет общих чисел, поэтому C[0] = 0.
При i = 1: 1 и 3 являются общими в A и B, поэтому C[1] = 2.
При i = 2: 1, 2 и 3 являются общими в A и B, поэтому C[2] = 3.
При i = 3: 1, 2, 3 и 4 являются общими в A и B, поэтому C[3] = 4.
*/

public class FindThePrefixCommonArrayOfTwoArrays {
    public static void main(String[] args) {
        int[] A = {1, 3, 2, 4};
        int[] B = {3, 1, 2, 4};
        System.out.println(Arrays.toString(findThePrefixCommonArray(A, B))); // [0, 2, 3, 4]
    }

    // findThePrefixCommonArray - функция, которая возвращает общий массив префиксов A и B.
    // time: O(n), space: O(n)
    private static int[] findThePrefixCommonArray(int[] A, int[] B) {
        int[] res = new int[A.length]; // общий массив префиксов A и B
        Set<Integer> set = new HashSet<>(); // хранит числа, которые присутствуют в индексе или перед ним как в A, так и в B
        int count = 0; // счетчик чисел, которые присутствуют в индексе или перед ним как в A, так и в B

        for (int i = 0; i < A.length; i++) {
            if (set.contains(A[i])) { // если число присутствует в индексе или перед ним в A, то увеличиваем счетчик на 1
                count++;
            }
            if (set.contains(B[i])) { // если число присутствует в индексе или перед ним в B, то увеличиваем счетчик на 1
                count++;
            }
            if (A[i] == B[i]) { // если числа в индексе i одинаковые, то увеличиваем счетчик на 1
                count++;
            }
            set.add(A[i]); // добавляем число в множество
            set.add(B[i]); // добавляем число в множество
            res[i] = count; // добавляем счетчик в результат
        }

        return res; // возвращаем результат
    }
}
