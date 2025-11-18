package array.SumOfTwoIntegers;

/* 371. Sum of Two Integers
https://leetcode.com/problems/sum-of-two-integers/description/

Даны два целых числа a и b, вернуть сумму двух целых чисел без использования операторов + и -.

Input: a = 1, b = 2
Output: 3
*/

public class SumOfTwoIntegers {
    public static void main(String[] args) {
        System.out.println(getSum(1, 2)); // 3
    }

    // getSum возвращает сумму двух целых чисел без использования операторов + и -.
    // time: O(1), space: O(1)
    private static int getSum(int a, int b) {
        while (b != 0) { // Продолжаем процесс, пока есть перенос
            int carry = a & b; // Операция a & b определяет биты, где оба числа имеют 1, что указывает на необходимость переноса
            a = a ^ b;      // Операция XOR складывает биты чисел без учета переноса
            b = carry << 1; // Перенос сдвигается на один бит влево, чтобы правильно добавить его к следующему биту суммы
        }

        return a; // Когда перенос становится нулевым, переменная a содержит итоговую сумму
    }
}