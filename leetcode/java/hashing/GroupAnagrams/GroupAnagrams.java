package leetcode.java.hashing.GroupAnagrams;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/* 49. Group Anagrams
https://leetcode.com/problems/group-anagrams/description/

Учитывая массив строк strs, сгруппируйте анаграммы вместе. Вы можете вернуть ответ в любом порядке.
Анаграмма — это слово или фраза, образованная путем перестановки букв другого слова или фразы, обычно с использованием
всех исходных букв ровно один раз.

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Input: strs = [""]
Output: [[""]]

Input: strs = ["a"]
Output: [["a"]]
*/

public class GroupAnagrams {
    public static void main(String[] args) {
        String[] strs = {"eat", "tea", "tan", "ate", "nat", "bat"};
        List<List<String>> result = groupAnagrams(strs);
        System.out.println(result); // [["bat"],["nat","tan"],["ate","eat","tea"]]
    }

    // groupAnagrams группирует анаграммы в исходном срезе строк.
    // time: O(n * m * log(m)), где n - количество строк в массиве, m - максимальная длина строки
    // space: O(n * m), где n - количество строк в массиве, m - максимальная длина строки
    private static List<List<String>> groupAnagrams(String[] strs) {
        if (strs == null || strs.length == 0) { // Если массив пуст, возвращаем пустой
            return new ArrayList<>();
        }

        Map<String, List<String>> anagramMap = new HashMap<>(); // Мапа для хранения групп анаграмм

        for (String str : strs) {
            // Сортируем символы строки для получения ключа-анаграммы
            char[] chars = str.toCharArray();
            Arrays.sort(chars); // Сортируем фрагмент рун
            String sortedStr = new String(chars); // aet, aet, ant, aet, ant, abt

            // Добавляем строку в соответствующую группу
            anagramMap.computeIfAbsent(sortedStr, k -> new ArrayList<>()).add(str);
        }

        // Преобразуем значения мапы в список списков
        return new ArrayList<>(anagramMap.values());  
    }
}
