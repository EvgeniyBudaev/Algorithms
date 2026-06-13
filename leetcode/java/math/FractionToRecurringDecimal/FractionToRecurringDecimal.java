package leetcode.java.math.FractionToRecurringDecimal;

import java.util.HashMap;
import java.util.Map;

/* 166. Fraction to Recurring Decimal
https://leetcode.com/problems/fraction-to-recurring-decimal/description/

Даны два целых числа, представляющих числитель и знаменатель дроби, вернуть дробь в строковом формате.
Если дробная часть повторяется, заключите повторяющуюся часть в скобки.
Если возможны несколько ответов, верните любой из них.
Гарантируется, что длина строки ответа меньше 104 для всех заданных входных данных.

Input: numerator = 1, denominator = 2
Output: "0.5"
*/

public class FractionToRecurringDecimal {
    public static void main(String[] args) {
        System.out.println(fractionToDecimal(1, 2)); // "0.5"
    }

    // fractionToDecimal Возвращает дробь в строковом формате.
    // time: O(n), space: O(n)
    private static String fractionToDecimal(int numerator, int denominator) {
        if (numerator == 0) { // если числитель равен 0, то возвращаем 0
            return "0";
        }

        StringBuilder res = new StringBuilder(); // создаем буфер для записи результата

        // если числитель и знаменатель имеют разные знаки, то добавляем знак "-"
        if ((numerator < 0) ^ (denominator < 0)) {
            res.append("-");
        }

        // Используем long, чтобы избежать переполнения при взятии модуля от Integer.MIN_VALUE
        long num = Math.abs((long) numerator);
        long den = Math.abs((long) denominator);

        res.append(num / den); // делим числитель на знаменатель и добавляем результат
        long rem = num % den;  // берем остаток от деления

        if (rem == 0) {        // если остаток равен 0, то возвращаем результат
            return res.toString();
        }

        res.append(".");       // добавляем точку в результат
        Map<Long, Integer> remainderMap = new HashMap<>(); // мапа для записи остатка и его позиции в строке

        while (rem != 0) {     // пока остаток от деления не равен 0
            // если остаток уже встречался, значит началось повторение
            if (remainderMap.containsKey(rem)) {
                int index = remainderMap.get(rem);
                res.insert(index, "("); // вставляем открывающую скобку на нужную позицию
                res.append(")");        // закрывающую скобку в конец
                return res.toString();
            }

            remainderMap.put(rem, res.length()); // записываем остаток и текущую длину строки
            rem *= 10;                           // умножаем остаток на 10
            res.append(rem / den);               // добавляем следующую цифру
            rem %= den;                          // берем новый остаток
        }

        return res.toString(); // возвращаем результат
    }
}
