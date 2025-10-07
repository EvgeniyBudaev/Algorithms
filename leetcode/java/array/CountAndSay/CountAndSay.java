package array.CountAndSay;

/* 38. Count and Say
https://leetcode.com/problems/count-and-say/description/

Последовательность count-and-say — это последовательность строк цифр, определяемая рекурсивной формулой:

countAndSay(1) = "1"
countAndSay(n) — это кодирование длины серии countAndSay(n - 1).
Кодирование длины серии (RLE) — это метод сжатия строк, который работает путем замены последовательных одинаковых
символов (повторяющихся 2 или более раз) на конкатенацию символа и числа, обозначающего количество символов
(длину серии). Например, чтобы сжать строку "3322251", мы заменяем "33" на "23", заменяем "222" на "32",
заменяем "5" на "15" и заменяем "1" на "11". Таким образом, сжатая строка становится "23321511".

Дано положительное целое число n, вернуть n-й элемент последовательности count-and-say.

Input: n = 4
Output: "1211"

Explanation:
countAndSay(1) = "1"
countAndSay(2) = RLE of "1" = "11"
countAndSay(3) = RLE of "11" = "21"
countAndSay(4) = RLE of "21" = "1211"
*/

public class CountAndSay {
    public static void main(String[] args) {
        System.out.println(countAndSay(4)); // "1211"
    }

    private static String countAndSay(int n) {
        String res = "1"; // Начинаем с первого элемента последовательности

        while (n > 1) { // Пока n больше 1, продолжаем генерировать последовательность
            StringBuilder cur = new StringBuilder(); // Текущая строка
            int i = 0; // Индекс текущего символа

            while (i < res.length()) { // Проходим по всем символам в текущей строке
                int count = 1; // Количество текущего символа

                // Если текущий символ равен следующему, увеличиваем счетчик
                while (i + 1 < res.length() && res.charAt(i) == res.charAt(i + 1)) {
                    i++; // Переходим к следующему символу
                    count++; // Увеличиваем счетчик
                }

                cur.append(count).append(res.charAt(i)); // Добавляем количество и текущий символ в текущую строку
                i++; // Переходим к следующему символу
            }

            res = cur.toString(); // Обновляем текущую строку
            n--; // Уменьшаем счетчик n на 1
        }

        return res; // Возвращаем результат
    }
}
