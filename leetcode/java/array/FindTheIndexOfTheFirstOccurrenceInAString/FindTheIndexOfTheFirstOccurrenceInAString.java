package array.FindTheIndexOfTheFirstOccurrenceInAString;

/* 28. Find the Index of the First Occurrence in a String
https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/

Для двух строк «needle» и «haystack» вернуть индекс первого вхождения «needle» в «haystack» или -1,
если «needle» не является частью «haystack».

Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Объяснение: «sad» встречается в индексах 0 и 6. Первое вхождение находится в индексе 0, поэтому мы возвращаем 0.
*/

public class FindTheIndexOfTheFirstOccurrenceInAString {
    public static void main(String[] args) {
        String haystack = "sadbutsad";
        String needle = "sad";
        System.out.println(strStr(haystack, needle)); // 0
    }

    // strStr возвращает индекс первого вхождения needle в haystack.
    // time: O(n*m), где n - длина haystack, m - длина needle, space: O(1)
    private static int strStr(String haystack, String needle) {
        // Если needle пустой, то возвращаем 0
        if (needle.isEmpty()) {
            return 0;
        }

        // Проходим по строке haystack
        for (int i = 0; i <= haystack.length() - needle.length(); i++) {
            // Проверяем, что в строке есть needle
            if (haystack.substring(i, i + needle.length()).equals(needle)) {
                return i; // Возвращаем индекс
            }
        }

        return -1; // Если needle не найдено, возвращаем -1
    }
}
