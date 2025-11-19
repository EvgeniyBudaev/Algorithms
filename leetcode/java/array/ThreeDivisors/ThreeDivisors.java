package array.ThreeDivisors;

import java.util.ArrayList;
import java.util.List;

/* 1952. Three Divisors
https://leetcode.com/problems/three-divisors/description/

Дано целое число n, вернуть true, если n имеет ровно три положительных делителя. В противном случае вернуть false.
Целое число m является делителем n, если существует целое число k, такое что n = k * m.

Input: n = 2
Output: false
Объяснение: у числа 2 всего два делителя: 1 и 2.

Input: n = 4
Output: true
Объяснение: 4 имеет три делителя: 1, 2 и 4.
*/

public class ThreeDivisors {
    public static void main(String[] args) {
        System.out.println(isThree(2)); // false
        System.out.println(isThree(4)); // true
    }

    // isThree принимает число и возвращает true если число имеет ровно три положительных делителя, иначе false.
    // time: O(n), space: O(1)
    private static boolean isThree(int n) {
        // Если количество делителей равно 3, значит число имеет ровно три положительных делителя
        return div(n).size() == 3;
    }

    // div принимает число и возвращает его делители.
    // time: O(n), space: O(n)
    private static List<Integer> div(int n) {
        List<Integer> result = new ArrayList<>(); // Создаем пустой список для хранения делителей числа

        for (int i = 1; i <= n; i++) { // Проходимся по всем числам от 1 до n
            if (n % i == 0) { // Если число делится нацело на i, то оно является делителем числа
                result.add(i);
            }
        }

        return result; // Возвращаем список делителей числа
    }
}