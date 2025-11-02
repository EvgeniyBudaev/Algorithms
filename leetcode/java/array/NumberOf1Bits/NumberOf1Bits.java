package array.NumberOf1Bits;

/* 191. Number of 1 Bits
https://leetcode.com/problems/number-of-1-bits/description/

Для заданного положительного целого числа n напишите функцию, которая возвращает количество установленных битов в его
двоичном представлении (также известном как вес Хэмминга).

Input: n = 11
Output: 3
Пояснение: Входная двоичная строка 1011 имеет всего три установленных бита.
*/

public class NumberOf1Bits {
    public static void main(String[] args) {
        System.out.println(hammingWeight(11)); // 3
    }

    // hammingWeight принимает положительное целое число и возвращает количество установленных битов в его двоичном представлении.
    // time: O(log(n)), space: O(log(n))
    private static int hammingWeight(int n) {
        // Преобразовать целое число в двоичное представление
        String binaryRepresentation = Integer.toBinaryString(n); // 1011

        // Извлечь двоичное представление в виде строки и подсчитать биты «1»
        int count = 0;
        for (int i = 0; i < binaryRepresentation.length(); i++) {
            if (binaryRepresentation.charAt(i) == '1') {
                count++;
            }
        }

        return count;
    }
}
