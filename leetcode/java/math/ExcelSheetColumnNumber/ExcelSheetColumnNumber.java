package leetcode.java.math.ExcelSheetColumnNumber;

/* 171. Excel Sheet Column Number
https://leetcode.com/problems/excel-sheet-column-number/description/

Если задана строка columnTitle, представляющая заголовок столбца, как он отображается на листе Excel,
вернуть соответствующий ему номер столбца.

Input: columnTitle = "A"
Output: 1
*/

public class ExcelSheetColumnNumber {
    public static void main(String[] args) {
        System.out.println(titleToNumber("A")); // 1
    }

    // titleToNumber - принимает строку и возвращает ее номер в excel.
    // time: O(n), space: O(1)
    private static int titleToNumber(String columnTitle) {
        int result = 0; // Инициализируем переменную для хранения итогового значения

        // Перебираем все символы строки
        for (int i = 0; i < columnTitle.length(); i++) {
            char v = columnTitle.charAt(i);

            // Берем остаток текущего символа и добавляем 1, потому что номер столбца в Excel начинается с 1
            // 65 это значение 'A' в таблице ASCII
            int num = (v % 65) + 1;

            // Умножаем на 26 и прибавляем текущий символ
            // 26 это количество букв в английском алфавите
            result = (result * 26) + num;
        }

        return result; // Возвращаем итоговое значение
    }
}
