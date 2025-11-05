package array.PlusOne;

import java.util.Arrays;

/* 66. Plus One
https://leetcode.com/problems/plus-one/description/

Вам дано большое целое число, представленное в виде массива целых чисел digits,
где каждая цифра[i] — это i-я цифра целого числа. Цифры упорядочены от наиболее значимых к наименее значимым слева
направо. Большое целое число не содержит начальных нулей.
Увеличьте большое целое число на единицу и верните полученный массив цифр.

Input: digits = [1,2,3]
Output: [1,2,4]
Объяснение: Массив представляет целое число 123.
Увеличение на единицу дает 123 + 1 = 124. Таким образом, результат должен быть [1,2,4].
*/

public class PlusOne {
    public static void main(String[] args) {
        int[] digits1 = {1, 2, 3};
        System.out.println(Arrays.toString(plusOne(digits1))); // [1,2,4]

        int[] digits2 = {9, 9, 9};
        System.out.println(Arrays.toString(plusOne(digits2))); // [1,0,0,0]
    }

    // plusOne увеличивает число на 1.
    // time: O(n), space: O(1)
    private static int[] plusOne(int[] digits) {
        int n = digits.length; // Получаем длину массива цифр

        // Проходим по цифрам справа налево
        for (int i = n - 1; i >= 0; i--) {
            // Если цифра меньше 9, просто увеличиваем её на 1 и возвращаем результат
            if (digits[i] < 9) {
                digits[i]++;
                return digits;
            }
            // Если цифра равна 9, заменяем её на 0 и продолжаем цикл (перенос 1 в следующий разряд)
            digits[i] = 0;
        }

        // Если все цифры были 9 (например, [9,9,9]), создаем новый массив с 1 в начале
        int[] newDigits = new int[n + 1];
        newDigits[0] = 1;
        return newDigits;
    }
}
