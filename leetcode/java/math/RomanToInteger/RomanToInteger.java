package leetcode.java.math.RomanToInteger;

import java.util.HashMap;
import java.util.Map;

/* 13. Roman to Integer
https://leetcode.com/problems/roman-to-integer/description/

Римские цифры представлены семью различными символами: I, V, X, L, C, D и M.

Символ Значение
I 1
V 5
X 10
L 50
C 100
D 500
M 1000
Например, 2 записывается как II в римской цифре, просто две единицы, сложенные вместе. 12 записывается как XII,
что просто X + II. Число 27 записывается как XXVII, что просто XX + V + II.

Римские цифры обычно записываются от большего к меньшему слева направо. Однако цифра для четырех не IIII.
Вместо этого число четыре записывается как IV. Поскольку единица стоит перед пятью, мы вычитаем ее, получая четыре.
Тот же принцип применим к числу девять, которое записывается как IX. Существует шесть случаев, когда используется
вычитание:

I можно разместить перед V (5) и X (10), чтобы получить 4 и 9.
X можно разместить перед L (50) и C (100), чтобы получить 40 и 90.
C можно разместить перед D (500) и M (1000), чтобы получить 400 и 900.
Дав римскую цифру, преобразуйте ее в целое число.

Input: s = "III"
Output: 3
Explanation: III = 3.
*/

public class RomanToInteger {
    public static void main(String[] args) {
        System.out.println(romanToInt("III")); // 3
    }

    // romanToInt преобразует римскую цифру в целое число.
    // time: O(n), space: O(1)
    private static int romanToInt(String s) {
        // Мапа значений римских цифр
        Map<Character, Integer> values = new HashMap<>();
        values.put('I', 1);
        values.put('V', 5);
        values.put('X', 10);
        values.put('L', 50);
        values.put('C', 100);
        values.put('D', 500);
        values.put('M', 1000);

        int sum = 0; // Переменная для хранения результата

        for (int i = 0; i < s.length(); i++) {
            // Проверяем, есть ли следующая цифра и если она больше текущей, вычитаем текущую из суммы
            if (i + 1 < s.length() && values.get(s.charAt(i)) < values.get(s.charAt(i + 1))) {
                sum -= values.get(s.charAt(i));
                // Если нет, добавляем текущую цифру к сумме
            } else {
                sum += values.get(s.charAt(i));
            }
        }

        return sum;
    }
}
