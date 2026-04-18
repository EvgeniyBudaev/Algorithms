package leetcode.java.backtracking.LetterCombinationsOfAPhoneNumber;

import java.util.ArrayList;
import java.util.List;

/* 17. Letter Combinations of a Phone Number
https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/

Учитывая строку, содержащую цифры от 2 до 9 включительно, верните все возможные комбинации букв, которые может
представлять число. Верните ответ в любом порядке.
Соответствие цифр буквам (как на кнопках телефона) приведено ниже. Обратите внимание, что 1 не соответствует
никаким буквам.
*/

/*
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
*/

public class LetterCombinationsOfAPhoneNumber {
    public static void main(String[] args) {
        System.out.println(letterCombinations("23")); // [ad, ae, af, bd, be, bf, cd, ce, cf]
        System.out.println(letterCombinations(""));   // []
        System.out.println(letterCombinations("2"));  // [a, b, c]  
    }

    // Маппинг цифр к соответствующим буквам (как на телефонной клавиатуре)
    private static final String[] DIGIT_TO_LETTERS = {
            "",     // 0
            "",     // 1
            "abc",  // 2
            "def",  // 3
            "ghi",  // 4
            "jkl",  // 5
            "mno",  // 6
            "pqrs", // 7
            "tuv",  // 8
            "wxyz"  // 9
    };

    /**
     * Генерирует все возможные комбинации букв, соответствующих введённым цифрам.
     *
     * @param digits строка, содержащая цифры от '2' до '9'
     * @return список всех возможных буквенных комбинаций
     * <p>
     * Time: O(4^n * n), где n — длина входной строки.
     * - 4 — максимальное количество букв на одной клавише (7, 9)
     * - 4^n — верхняя оценка количества комбинаций
     * - *n — время на создание каждой строки результата
     * <p>
     * Space: O(n) — вспомогательная память (глубина рекурсии + StringBuilder),
     * не считая памяти под результат. С учётом вывода: O(4^n * n).
     */
    private static List<String> letterCombinations(String digits) {
        List<String> result = new ArrayList<>();

        // Edge case: пустая строка → пустой результат
        if (digits == null || digits.isEmpty()) {
            return result;
        }

        backtrack(result, new StringBuilder(), digits, 0);
        return result;
    }

    /**
     * Рекурсивная функция обратного отслеживания для генерации комбинаций.
     * <p>
     * Принцип работы:
     * - На каждом шаге обрабатываем одну цифру из входной строки
     * - Для текущей цифры перебираем все соответствующие буквы
     * - Добавляем букву в буфер, рекурсивно переходим к следующей цифре
     * - После возврата из рекурсии выполняем «откат» (backtrack): удаляем
     * последнюю добавленную букву, чтобы попробовать следующую опцию
     * - Когда индекс достигает конца строки — сохраняем готовую комбинацию
     *
     * @param result  список для накопления итоговых комбинаций
     * @param current буфер для построения текущей комбинации (StringBuilder для эффективности)
     * @param digits  исходная строка с цифрами
     * @param index   текущая позиция в строке digits (какую цифру обрабатываем)
     */
    private static void backtrack(List<String> result, StringBuilder current, String digits, int index) {
        // Базовый случай: все цифры обработаны — сохраняем комбинацию
        if (index == digits.length()) {
            result.add(current.toString());
            return;
        }

        // Получаем буквы, соответствующие текущей цифре
        char digit = digits.charAt(index);
        String letters = DIGIT_TO_LETTERS[digit - '0'];

        // Перебираем все возможные буквы для текущей цифры
        for (int i = 0; i < letters.length(); i++) {
            current.append(letters.charAt(i));           // Выбираем букву
            backtrack(result, current, digits, index + 1); // Рекурсия к следующей цифре
            current.deleteCharAt(current.length() - 1);  // Откат: удаляем букву для следующей итерации
        }
    }
}
