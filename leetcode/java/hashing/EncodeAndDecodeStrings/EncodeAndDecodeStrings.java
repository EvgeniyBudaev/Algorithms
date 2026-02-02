package leetcode.java.hashing.EncodeAndDecodeStrings;

import java.util.ArrayList;
import java.util.List;

/* 659. Encode and Decode Strings

Разработайте алгоритм для кодирования списка строк в одну строку. Закодированная строка затем декодируется обратно в
исходный список строк.
Пожалуйста, реализуйте кодирование и декодирование

Input: ["neet","code","love","you"]
Output:["neet","code","love","you"]

Input: ["we","say",":","yes"]
Output: ["we","say",":","yes"]
*/

public class EncodeAndDecodeStrings {
    public static void main(String[] args) {
        List<String> input = List.of("Hello", "World");
        String encoded = encode(input);
        System.out.println("encode: " + encoded); // 00000101Hello00000101World

        List<String> decoded = decode(encoded);
        System.out.println("decode: " + decoded); // [Hello, World]
    }

    // Кодирует список строк в одну строку.
    // Длина каждой строки сохраняется как префикс фиксированного размера перед фактическим содержимым строки.
    private static String encode(List<String> strs) {
        StringBuilder sb = new StringBuilder();
        for (String str : strs) {
            sb.append(getLengthPrefix(str));
            sb.append(str);
        }
        return sb.toString();
    }

    // getLengthPrefix - формирует 8‑символьный двоичный префикс длины строки (с ведущими нулями).
    private static String getLengthPrefix(String str) {
        int length = str.length();
        String binary = Integer.toBinaryString(length);
        // Добавляем ведущие нули до длины 8 символов
        return String.format("%8s", binary).replace(' ', '0');
    }

    // Декодирует одну строку в список строк, считывая префикс длины фиксированного размера
    // и затем считываем соответствующее количество символов.
    private static List<String> decode(String encoded) {
        List<String> result = new ArrayList<>();
        int left = 0;
        int totalLength = encoded.length();

        while (left < totalLength) {
            int PREFIX_BITS = 8; // фиксированная ширина префикса в символах
            int right = left + PREFIX_BITS;
            if (right > totalLength) {
                throw new IllegalArgumentException(
                    "Некорректный формат: неполный префикс длины на позиции " + left);
            }

            String binaryPrefix = encoded.substring(left, right);
            int length;
            try {
                length = Integer.parseInt(binaryPrefix, 2);
            } catch (NumberFormatException e) {
                throw new IllegalArgumentException(
                    "Некорректный двоичный префикс: '" + binaryPrefix + "'", e);
            }

            int contentEnd = right + length;
            if (contentEnd > totalLength) {
                throw new IllegalArgumentException(
                    "Некорректный формат: содержимое выходит за границы строки (требуется " +
                    length + " символов, доступно " + (totalLength - right) + ")");
            }

            String decoded = encoded.substring(right, contentEnd);
            result.add(decoded);
            left = contentEnd;
        }

        return result;
    }
}
