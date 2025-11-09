package array.ReverseInteger;

/* 7. Reverse Integer
https://leetcode.com/problems/reverse-integer/description/

Дано знаковое 32-битное целое число x, вернуть x с перевернутыми цифрами. Если переворот x приводит к выходу значения
за пределы знакового 32-битного целого числа [-231, 231 - 1], то вернуть 0.

Предположим, что среда не позволяет вам хранить 64-битные целые числа (со знаком или без знака).

Input: x = 123
Output: 321
*/

public class ReverseInteger {
    public static void main(String[] args) {
        System.out.println(reverse(123)); // 321
    }

    // reverse переворачивает число.
    // Time: O(n), space: O(1)
    private static int reverse(int x) {
        long finalNum = 0;

        while (x != 0) {
            int lastDig = x % 10;
            finalNum += lastDig;
            finalNum = finalNum * 10;
            x = x / 10;
        }

        finalNum = finalNum / 10;

        if (finalNum > Integer.MAX_VALUE || finalNum < Integer.MIN_VALUE) {
            return 0;
        }

        if (x < 0) {
            return (int) (-1 * finalNum);
        }
        
        return (int) finalNum;
    }
}
