package stacksAndQueues.MinimumRemoveToMakeValidParentheses;

import java.util.ArrayDeque;
import java.util.Deque;

/* 1249. Minimum Remove to Make Valid Parentheses
https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/description/

Дана строка s из '(' , ')' и строчных английских символов.

Ваша задача — удалить минимальное количество скобок ('(' или ')', в любых позициях ) так, чтобы результирующая строка
скобок была допустимой и возвращала любую допустимую строку.

Формально строка скобок допустима тогда и только тогда, когда:

Это пустая строка, содержит только строчные символы или
Она может быть записана как AB (A, объединенная с B), где A и B — допустимые строки, или
Она может быть записана как (A), где A — допустимая строка.

Input: s = "lee(t(c)o)de)"
Output: "lee(t(c)o)de"
Пояснение: «lee(t(co)de)», «lee(t(c)ode)» также будут приняты.

Input: s = "a)b(c)d"
Output: "ab(c)d"
*/

public class MinimumRemoveToMakeValidParentheses {
    public static void main(String[] args) {
        String s = "lee(t(c)o)de)";
        System.out.println(minRemoveToMakeValid(s)); // "lee(t(c)o)de"
    }

    // minRemoveToMakeValid удаляет минимальное количество скобок ('(' или ')', в любых позициях ) так,
    // чтобы результирующая строка скобок была допустимой.
    // time: O(n), space: O(n)
    private static String minRemoveToMakeValid(String s) {
        char[] chars = s.toCharArray(); // Массив символов для модификации
        Deque<Integer> stack = new ArrayDeque<>(); // Стек для индексов открывающих скобок '('

        // Первый проход: помечаем лишние закрывающие скобки
        for (int i = 0; i < chars.length; i++) {
            if (chars[i] == '(') {
                stack.push(i);
            } else if (chars[i] == ')') {
                if (!stack.isEmpty()) {
                    stack.pop(); // Есть пара — убираем из стека
                } else {
                    chars[i] = 0; // Помечаем лишнюю ')' для удаления
                }
            }
        }

        // Помечаем все оставшиеся непарные '('
        while (!stack.isEmpty()) {
            chars[stack.pop()] = 0;
        }

        // Собираем результат, пропуская помеченные символы (char == 0)
        StringBuilder result = new StringBuilder();
        for (char c : chars) {
            if (c != 0) {
                result.append(c);
            }
        }

        return result.toString();
    }
}
