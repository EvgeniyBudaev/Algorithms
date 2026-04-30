package leetcode.java.heaps.SortCharactersByFrequency;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/* 451. Sort Characters By Frequency
https://leetcode.com/problems/sort-characters-by-frequency/description/

Дана строка s, отсортируйте ее в порядке убывания на основе частоты символов. Частота символа — это количество раз,
которое он появляется в строке.
Верните отсортированную строку. Если есть несколько ответов, верните любой из них.

Input: s = "tree"
Output: "eert"
Объяснение: «e» встречается дважды, а «r» и «t» встречаются по одному разу.
Поэтому «e» должно встречаться перед «r» и «t». Поэтому «eetr» также является допустимым ответом.
*/

public class SortCharactersByFrequency {
    public static void main(String[] args) {
        System.out.println(frequencySort("tree"));  // "eert" или "eetr" (оба корректны)
        System.out.println(frequencySort("cccaaa")); // "cccaaa" или "aaaccc"
        System.out.println(frequencySort("Aabb"));   // "bbAa" или "bbaA"
    }

    /**
     * CharFreq - вспомогательная запись для хранения пары «символ-частота».
     * В Java 14+ можно использовать record, для старых версий — обычный класс.
     */
    private static class CharFreq {
        char ch;      // Символ
        int freq;     // Частота появления

        public CharFreq(char ch, int freq) {
            this.ch = ch;
            this.freq = freq;
        }
    }

    /**
     * frequencySort - сортировка символов строки по частоте в порядке убывания.
     * time: O(n*log(n)), space: O(n)
     *
     * @param s входная строка
     * @return строка с символами, отсортированными по убыванию частоты
     */
    private static String frequencySort(String s) {
        // Подсчитываем частоту символов
        Map<Character, Integer> counter = new HashMap<>();
        for (char ch : s.toCharArray()) {
            counter.put(ch, counter.getOrDefault(ch, 0) + 1);
        }

        // Создаём список пар «символ-частота»
        List<CharFreq> charFreqList = new ArrayList<>(counter.size());
        for (Map.Entry<Character, Integer> entry : counter.entrySet()) {
            charFreqList.add(new CharFreq(entry.getKey(), entry.getValue()));
        }

        // Сортируем по частоте в порядке убывания
        charFreqList.sort((a, b) -> b.freq - a.freq);

        // Строим строку результата
        StringBuilder result = new StringBuilder(s.length());
        for (CharFreq cf : charFreqList) {
            for (int i = 0; i < cf.freq; i++) {
                result.append(cf.ch);
            }
        }

        return result.toString();
    }
}
