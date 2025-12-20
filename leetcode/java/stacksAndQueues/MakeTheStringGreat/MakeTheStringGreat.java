package stacksAndQueues.MakeTheStringGreat;

import java.util.ArrayDeque;
import java.util.Deque;

/* 1544. Make The String Great
https://leetcode.com/problems/make-the-string-great/description/

Дана строка s, состоящая из строчных и прописных английских букв.
Хорошая строка — это строка, в которой нет двух соседних символов s[i] и s[i + 1], где:
0 <= i <= s.length - 2
s[i] — строчная буква, а s[i + 1] — та же буква, но в верхнем регистре или наоборот.
Чтобы сделать строку хорошей, вы можете выбрать два соседних символа, которые делают строку плохой, и удалить их.
Вы можете продолжать делать это, пока строка не станет хорошей.
Верните строку после того, как исправили ее. Ответ гарантированно будет уникальным при заданных ограничениях.
Обратите внимание, что пустая строка тоже подойдет.

Input: s = "leEeetcode"
Output: "leetcode"
Explanation: На первом этапе, либо вы выберете i = 1, либо i = 2, в результате оба варианта «leEeetcode» будут сокращены
до «leetcode».

Input: s = "abBAcC"
Output: ""
Explanation: У нас есть много возможных сценариев, и все они ведут к одному и тому же ответу. Например:
"abBAcC" --> "aAcC" --> "cC" --> ""
"abBAcC" --> "abBA" --> "aA" --> ""

Input: s = "s"
Output: "s"
*/

public class MakeTheStringGreat {
    public static void main(String[] args) {
        System.out.println(makeGood("leEeetcode")); // "leetcode"
    }

    // makeGood возвращает хорошую строку.
    // time: O(n), space: O(n), где n - длина строки
    private static String makeGood(String s) {
        if (s.isEmpty()) {
            return "";
        }

        Deque<Character> stack = new ArrayDeque<>(); // Стек символов

        for (char ch : s.toCharArray()) {
            if (!stack.isEmpty()) {
                char lastChar = stack.peek(); // Получаем последний символ из стека
                // Если символы являются противоположными регистрами
                if (ch != lastChar && Character.toLowerCase(ch) == Character.toLowerCase(lastChar)) {
                    stack.pop(); // Удаляем последний символ
                    continue;
                }
            }
            stack.push(ch); // Добавляем символ в стек
        }

        // Собираем результат в обратном порядке (т.к. стек LIFO)
        StringBuilder result = new StringBuilder();
        while (!stack.isEmpty()) {
            result.append(stack.pop());
        }

        return result.reverse().toString();
    }
}
