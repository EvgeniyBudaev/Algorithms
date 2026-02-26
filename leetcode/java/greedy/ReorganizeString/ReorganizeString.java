package greedy.ReorganizeString;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/* 767. Reorganize String
https://leetcode.com/problems/reorganize-string/description/
solution https://leetcode.com/problems/reorganize-string/submissions/1333724578/

Дана строка s, переставьте символы s так, чтобы любые два соседних символа не были одинаковыми.
Верните любую возможную перестановку s или верните "", если это невозможно.

Input: s = "aab"
Output: "aba"

Input: s = "aaab"
Output: ""
*/

public class ReorganizeString {
    public static void main(String[] args) {
        System.out.println(reorganizeString("aab")); // "aba"
    }

    // reorganizeString переставляет символы в строке так, чтобы любые два соседних символа не были одинаковыми.
    // time: O(n), space: O(n)
    private static String reorganizeString(String s) {
        Map<Character, Integer> freqMap = new HashMap<>(); // Создаем карту для подсчета частот символов
        for (char c : s.toCharArray()) {
            freqMap.put(c, freqMap.getOrDefault(c, 0) + 1);
        }

        List<CharCount> chars = new ArrayList<>(); // Создаем список для хранения символов и их частот
        // Добавляем символы и их частоты в список
        for (Map.Entry<Character, Integer> entry : freqMap.entrySet()) {
            chars.add(new CharCount(entry.getKey(), entry.getValue())); // Добавляем символ и его частоту
        }

        // Сортируем по убыванию частоты
        chars.sort((a, b) -> b.count - a.count);

        // Проверяем возможность реорганизации
        if (chars.get(0).count > (s.length() + 1) / 2) {
            return "";
        }

        char[] result = new char[s.length()]; // Создаем результат и заполняем его
        int position = 0; // Индекс для заполнения результата

        for (CharCount cc : chars) { // Делаем перебор для каждого символа в отсортированном списке
            for (int i = 0; i < cc.count; i++) { // Для каждого символа заполняем его в результирующей строке
                result[position] = cc.c; // Заполняем символ в результирующей строке
                position += 2; // Переходим к следующему индексу для заполнения
                // Если мы достигли конца результирующей строки, начинаем с начала
                if (position >= result.length) {
                    position = 1;
                }
            }
        }

        return new String(result); // Возвращаем реорганизованную строку
    }

    // CharCount для сортировки символов по частоте
    private static class CharCount {
        char c; // Символ
        int count; // Частота

        CharCount(char c, int count) {
            this.c = c;
            this.count = count;
        }
    }
}
