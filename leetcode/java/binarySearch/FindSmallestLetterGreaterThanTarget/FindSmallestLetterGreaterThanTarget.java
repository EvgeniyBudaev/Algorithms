package binarySearch.FindSmallestLetterGreaterThanTarget;

/* 744. Find Smallest Letter Greater Than Target
https://leetcode.com/problems/find-smallest-letter-greater-than-target/description/

Вам дан массив символов-букв, отсортированный в неубывающем порядке, и целевой символ. В буквах есть как минимум два
разных символа.
Верните наименьший буквенный символ, который лексикографически больше целевого. Если такого символа не существует,
вернуть первый символ буквами.

Input: letters = ["c","f","j"], target = "a"
Output: "c"
Пояснение: Самый маленький символ, который лексикографически больше буквы «a», — это «c».
*/

public class FindSmallestLetterGreaterThanTarget {
    public static void main(String[] args) {
        char[] letters = {'c', 'f', 'j'};
        System.out.println(nextGreatestLetter(letters, 'a')); // 'c'
    }

    // nextGreatestLetter возвращает наименьший буквенный символ, который лексикографически больше целевого символа.
    // time: O(log n), space: O(1)
    private static char nextGreatestLetter(char[] letters, char target) {
        int length = letters.length; // Длина массива

        // Обработка крайних случаев
        if (target < letters[0] || target >= letters[length - 1]) {
            return letters[0];
        }

        // Бинарный поиск
        int low = 0, high = length - 1;
        while (low <= high) {
            int mid = low + (high - low) / 2; // предотвращает переполнение
            char guess = letters[mid];

            // Если текущий символ больше целевого, и предыдущий <= целевого — возвращаем текущий символ
            if (guess > target && letters[mid - 1] <= target) {
                return guess;
            }

            if (guess > target) {
                high = mid - 1;
            } else {
                low = mid + 1;
            }
        }

        return letters[0]; // На случай, если ничего не найдено (по логике не должно выполняться)
    }
}
