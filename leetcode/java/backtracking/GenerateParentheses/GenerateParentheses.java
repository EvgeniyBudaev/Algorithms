package leetcode.java.backtracking.GenerateParentheses;

import java.util.ArrayList;
import java.util.List;

/* 22. Generate Parentheses
https://leetcode.com/problems/generate-parentheses/description/

Учитывая n пар круглых скобок, напишите функцию, которая генерирует все комбинации правильных круглых скобок.

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

Input: n = 1
Output: ["()"]
*/

public class GenerateParentheses {
    public static void main(String[] args) {
        System.out.println(generateParenthesis(1)); // [()]
        System.out.println(generateParenthesis(2)); // [(()), ()()]
        System.out.println(generateParenthesis(3)); // [((())), (()()), (())(), ()(()), ()()()]  
    }

    /**
     * Генерирует все правильные скобочные последовательности длины 2n.
     *
     * @param n количество пар скобок
     * @return список всех валидных комбинаций
     * <p>
     * Time: O(4^n / √n) — количество валидных последовательностей равно n-му числу Каталана,
     * каждая строка строится за O(n). Upper bound: O(4^n).
     * Space: O(n) — вспомогательная память (глубина рекурсии + StringBuilder),
     * не считая памяти под результат. С учётом вывода: O(4^n / √n * n).
     */
    private static List<String> generateParenthesis(int n) {
        List<String> result = new ArrayList<>();
        backtrack(result, new StringBuilder(), 0, 0, n);
        return result;
    }

    /**
     * Рекурсивная функция обратного отслеживания для построения валидных последовательностей.
     * <p>
     * Принцип работы:
     * - Добавляем '(', если количество открытых скобок меньше n
     * - Добавляем ')', если количество закрытых меньше открытых (чтобы не нарушить валидность)
     * - Когда длина строки достигает 2*n — сохраняем результат
     * - После рекурсивного вызова выполняем «откат» (backtrack): удаляем последний символ,
     * чтобы исследовать другие ветви дерева решений
     *
     * @param result  список для накопления итоговых комбинаций
     * @param current буфер для построения текущей последовательности (StringBuilder для эффективности)
     * @param open    текущее количество открытых скобок '('
     * @param close   текущее количество закрытых скобок ')'
     * @param n       целевое количество пар скобок
     */
    private static void backtrack(List<String> result, StringBuilder current, int open, int close, int n) {
        if (current.length() == 2 * n) {
            result.add(current.toString());
            return;
        }

        if (open < n) {
            current.append('(');
            backtrack(result, current, open + 1, close, n);
            current.deleteCharAt(current.length() - 1); // backtrack
        }

        if (close < open) {
            current.append(')');
            backtrack(result, current, open, close + 1, n);
            current.deleteCharAt(current.length() - 1); // backtrack
        }
    }
}
