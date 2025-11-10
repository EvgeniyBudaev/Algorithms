package array.ReverseWordsInAString;

import java.util.Arrays;
import java.util.Collections;

/* 151. Reverse Words in a String
https://leetcode.com/problems/reverse-words-in-a-string/description/

Дана входная строка s, обратный порядок слов.
Слово определяется как последовательность непробельных символов. Слова в s будут разделены как минимум одним пробелом.
Верните строку слов в обратном порядке, объединенных одним пробелом.

Обратите внимание, что s может содержать начальные или конечные пробелы или несколько пробелов между двумя словами.
Возвращаемая строка должна содержать только один пробел, разделяющий слова. Не включайте никаких дополнительных пробелов.

Input: s = "the sky is blue"
Output: "blue is sky the"

Input: s = "  hello world  "
Output: "world hello"
*/

public class ReverseWordsInAString {
    public static void main(String[] args) {
        System.out.println(reverseWords("the sky is blue")); // "blue is sky the"
    }

    // reverseWords переворачивает слова в строке.
    // time: O(n), space: O(n)
    private static String reverseWords(String s) {
        // Удаляем начальные и конечные пробелы, разбиваем на слова
        String[] words = s.trim().split("\\s+");

        // Переворачиваем порядок слов
        Collections.reverse(Arrays.asList(words));

        // Объединяем слова в строку с разделителем пробел
        return String.join(" ", words);
    }
}
