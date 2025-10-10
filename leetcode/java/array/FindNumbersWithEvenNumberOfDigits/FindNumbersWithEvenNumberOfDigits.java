package array.FindNumbersWithEvenNumberOfDigits;

/* 1295. Find Numbers with Even Number of Digits
https://leetcode.com/problems/find-numbers-with-even-number-of-digits/description/

Найдите числа с четным числом цифр.
Дан массив целых чисел, верните, сколько из них содержат четное количество цифр.

Input: nums = [12,345,2,6,7896]
Output: 2

Input: nums = [555,901,482,1771]
Output: 1
*/

public class FindNumbersWithEvenNumberOfDigits {
    public static void main(String[] args) {
        int[] nums = {12, 345, 2, 6, 7896};
        System.out.println(findNumbers(nums)); // 2
    }

    // findNumbers возвращает количество чисел с четным числом цифр.
    // time: O(n), space: O(1)
    private static int findNumbers(int[] nums) {
        int counter = 0; // Счетчик чисел с четным числом цифр

        for (int num : nums) {
            // Проверяем, является ли число четным
            if (String.valueOf(num).length() % 2 == 0) {
                counter++; // Увеличиваем счетчик
            }
        }

        return counter; // Возвращаем счетчик чисел с четным числом цифр
    }
}
