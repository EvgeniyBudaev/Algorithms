package stacksAndQueues.ValidParentheses;

import java.util.ArrayDeque;
import java.util.Deque;
import java.util.HashMap;
import java.util.Map;

/* 20. Valid Parentheses
https://leetcode.com/problems/valid-parentheses/description/

Учитывая строку s, содержащую только символы '(', ')', '{', '}', '[' и ']', определите, является ли входная строка
допустимой.

Входная строка действительна, если:

Открытые скобки должны закрываться скобками того же типа.
Открытые скобки должны закрываться в правильном порядке.
Каждой закрывающей скобке соответствует открытая скобка того же типа.

Input: s = "()"
Output: true

Input: s = "()[]{}"
Output: true

Input: s = "(]"
Output: false
*/

public class ValidParentheses {
    public static void main(String[] args) {
        System.out.println(isValid("()"));       // true
        System.out.println(isValid("()[]{}"));   // true
        System.out.println(isValid("(]"));       // false
    }

    // isValid проверяет, что каждой закрывающей скобке соответствует открытая скобка того же типа.
    // time: O(n), space: O(n), где n - длина строки.
    private static boolean isValid(String s) {
        Deque<Character> stack = new ArrayDeque<>(); // Стек для хранения открывающих скобок
        Map<Character, Character> brackets = new HashMap<>();
        brackets.put('}', '{');
        brackets.put(']', '[');
        brackets.put(')', '(');

        for (char ch : s.toCharArray()) {
            if (brackets.containsKey(ch)) {
                // Это закрывающая скобка
                char matching = brackets.get(ch);
                // Проверяем, соответствует ли последняя открывающая скобка
                if (stack.isEmpty() || stack.peek() != matching) {
                    return false;
                }
                stack.pop(); // Удаляем соответствующую открывающую скобку из стека
            } else {
                stack.push(ch); // Добавляем открывающую скобку в стек
            }
        }

        return stack.isEmpty(); // Все скобки сбалансированы, если стек пуст
    }
}
