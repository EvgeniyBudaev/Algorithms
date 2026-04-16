package leetcode.java.backtracking.GeneralizedAbbreviation;

import java.util.ArrayList;
import java.util.List;

/* https://algo.monster/liteproblems/320

Учитывая строковое слово, верните список всех возможных обобщенных сокращений слова. Верните ответ в любом порядке.

Input: word = "word"
Output: ["4","3d","2r1","2rd","1o2","1o1d","1or1","1ord","w3","w2d","w1r1","w1rd","wo2","wo1d","wor1","word"]

Input: word = "a"
Output: ["1","a"]
*/

public class GeneralizedAbbreviation {
    public static void main(String[] args) {
        System.out.println(generateAbbreviations("word"));
        // Вывод: [4, 3d, 2r1, 2rd, 1o2, 1o1d, 1or1, 1ord, w3, w2d, w1r1, w1rd, wo2, wo1d, wor1, word]
    }

    /**
     * Генерирует все возможные сокращения для слова.
     * Сокращение заменяет последовательные символы их количеством.
     * Например: "word" → "4", "3d", "2r1", "w3", "word" и т.д.
     * <p>
     * Time: O(n * 2^n)
     * - 2^n — общее количество возможных сокращений (каждый символ либо входит в число, либо остаётся буквой)
     * - n — время на конкатенацию частей при построении каждой строки-результата
     * Space: O(n) — не считая выходных данных
     * - O(n) — глубина рекурсивного стека
     * - O(n) — хранение текущего списка частей сокращения
     * - O(n * 2^n) — хранение всех результатов (выходные данные, не учитывается в пространственной сложности)
     */
    private static List<String> generateAbbreviations(String word) {
        List<String> abbreviations = new ArrayList<>();
        List<String> current = new ArrayList<>();
        dfs(word, current, abbreviations);
        return abbreviations;
    }

    /**
     * Вспомогательная рекурсивная функция с обратным ходом (backtracking)
     * для генерации всех возможных сокращений слова.
     *
     * @param remaining     оставшаяся часть слова для обработки
     * @param current       текущая формируемая комбинация частей сокращения
     * @param abbreviations список для хранения всех найденных сокращений
     *                      <p>
     *                      Time: O(n * 2^n) — перебор всех комбинаций с конкатенацией
     *                      Space: O(n) — глубина рекурсивного стека
     */
    private static void dfs(String remaining, List<String> current, List<String> abbreviations) {
        // Базовый случай: строка полностью обработана
        if (remaining.isEmpty()) {
            abbreviations.add(String.join("", current));
            return;
        }

        // Вариант 1: Пробуем сократить префикс длиной i (заменить i символов числом)
        for (int i = 1; i <= remaining.length(); i++) {
            // Добавляем числовое сокращение
            current.add(String.valueOf(i));

            if (i < remaining.length()) {
                // После числа обязательно добавляем следующий символ,
                // чтобы разделить числа (иначе "11" было бы неоднозначно)
                current.add(String.valueOf(remaining.charAt(i)));
                // Рекурсия для оставшейся части после сокращения + разделителя
                dfs(remaining.substring(i + 1), current, abbreviations);
                // Backtracking: удаляем разделитель
                current.remove(current.size() - 1);
            } else {
                // Сократили всю оставшуюся строку — завершаем комбинацию
                dfs("", current, abbreviations);
            }

            // Backtracking: удаляем числовое сокращение
            current.remove(current.size() - 1);
        }

        // Вариант 2: Оставляем первый символ как букву (без сокращения)
        current.add(String.valueOf(remaining.charAt(0)));
        dfs(remaining.substring(1), current, abbreviations);
        // Backtracking: удаляем добавленный символ
        current.remove(current.size() - 1);
    }
}
