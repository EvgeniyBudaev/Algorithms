package array.LongestCommonPrefix;

/* 14. Longest Common Prefix
https://leetcode.com/problems/longest-common-prefix/description/

Напишите функцию для поиска самой длинной общей строки префикса среди массива строк.
Если общего префикса нет, верните пустую строку "".

Input: strs = ["flower","flow","flight"]
Output: "fl"
*/

public class LongestCommonPrefix {
    public static void main(String[] args) {
        String[] strs = {"flower", "flow", "flight"};
        System.out.println(longestCommonPrefix(strs)); // fl
    }

    // longestCommonPrefix - возвращает самый длинный общий префикс среди массива строк.
    // time: O(S), где S - суммарная длина всех строк в массиве strs.
    // space: O(1)
    private static String longestCommonPrefix(String[] strs) {
        // Если массив пуст или null, возвращаем пустую строку
        if (strs == null || strs.length == 0) {
            return "";
        }

        String prefix = strs[0]; // Задаем префикс как первую строку
        for (int i = 1; i < strs.length; i++) { // Проходим по остальным строкам
            // Проверяем, что текущая строка начинается с префикса
            while (strs[i].indexOf(prefix) != 0) { // indexOf возвращает 0, если строка начинается с префикса
                // Удаляем последний символ из prefix, пока не найдем общий префикс
                prefix = prefix.substring(0, prefix.length() - 1);
                if (prefix.isEmpty()) { // Если префикс стал пустым, возвращаем его
                    return "";
                }
            }
        }

        return prefix; // Возвращаем найденный префикс
    }
}
