package array.TopKFrequentWords;

import java.util.ArrayList;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/* 692. Top K Frequent Words
https://leetcode.com/problems/top-k-frequent-words/description/

Учитывая массив строковых слов и целое число k, верните k наиболее часто встречающихся строк.
Возвращает ответ, отсортированный по частоте от самого высокого до самого низкого. Отсортируйте слова с одинаковой
частотой по их лексикографическому порядку.

Input: words = ["i","love","leetcode","i","love","coding"], k = 2
Output: ["i","love"]
Пояснение: «i» и «love» — два наиболее часто встречающихся слова.
Обратите внимание, что «i» стоит перед словом «love» из-за более низкого алфавитного порядка.
*/

public class TopKFrequentWords {
    public static void main(String[] args) {
        String[] words = {"i", "love", "leetcode", "i", "love", "coding"};
        System.out.println(topKFrequent(words, 2)); // ["i", "love"]
    }

    // topKFrequent возвращает k наиболее часто встречающихся слов.
    // time: O(n log n), где n - длина массива words, space: O(n), где n - длина массива words
    private static List<String> topKFrequent(String[] words, int k) {
        // Создаем map для подсчета частоты слов
        Map<String, Integer> freqMap = new HashMap<>();
        for (String word : words) {
            freqMap.put(word, freqMap.getOrDefault(word, 0) + 1);
        }

        // Создаем список для хранения пар (слово, частота)
        List<WordFreq> wordFreqs = new ArrayList<>();
        for (Map.Entry<String, Integer> entry : freqMap.entrySet()) {
            wordFreqs.add(new WordFreq(entry.getKey(), entry.getValue()));
        }

        // Сортируем список по частоте и лексикографическому порядку
        Collections.sort(wordFreqs, (a, b) -> {
            if (a.count != b.count) {
                return b.count - a.count; // По убыванию частоты
            }
            return a.word.compareTo(b.word); // По возрастанию лексикографического порядка
        });

        // Выбираем первые k элементов
        List<String> result = new ArrayList<>();
        for (int i = 0; i < k && i < wordFreqs.size(); i++) {
            result.add(wordFreqs.get(i).word);
        }

        return result; // Возвращаем k наиболее часто встречающихся слов
    }

    // Вспомогательный класс для хранения пары (слово, частота)
    static class WordFreq {
        String word;
        int count;

        WordFreq(String word, int count) {
            this.word = word;
            this.count = count;
        }
    }

    // Версия с Stream API: time: O(n log n), space: O(n)
    // public static List<String> topKFrequent(String[] words, int k) {
    //     return Arrays.stream(words)
    //         .collect(Collectors.groupingBy(
    //             word -> word,
    //             Collectors.counting()
    //         ))
    //         .entrySet()
    //         .stream()
    //         .sorted((a, b) -> {
    //             if (a.getValue().equals(b.getValue())) {
    //                 return a.getKey().compareTo(b.getKey()); // Лексикографический порядок
    //             }
    //             return b.getValue().compareTo(a.getValue()); // По убыванию частоты
    //         })
    //         .limit(k)
    //         .map(Map.Entry::getKey)
    //         .collect(Collectors.toList());
    // }
}