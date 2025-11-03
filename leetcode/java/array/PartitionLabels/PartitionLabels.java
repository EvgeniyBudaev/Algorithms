package array.PartitionLabels;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/* 763. Partition Labels
https://leetcode.com/problems/partition-labels/description/

Вам дана строка s. Мы хотим разделить строку на как можно большее количество частей, чтобы каждая буква появлялась не
более чем в одной части. Обратите внимание, что разбиение выполнено так, что после объединения всех частей по порядку
результирующая строка должна быть s. Верните список целых чисел, представляющих размер этих частей.

Input: s = "ababcbacadefegdehijhklij"
Output: [9,7,8]
Объяснение:
Раздел — «ababcbaca», «defegde», «hijhklij».
Это раздел, в котором каждая буква встречается не более чем в одной части.
Раздел типа «ababcbacadefegde», «hijhklij» неверен, поскольку он разбивает s на меньшее количество частей.

Input: s = "eccbbbbdec"
Output: [10]
*/

public class PartitionLabels {
    public static void main(String[] args) {
        System.out.println(partitionLabels("ababcbacadefegdehijhklij")); // [9, 7, 8]
    }

    // partitionLabels разделяет строки на максимальное количество частей так, чтобы каждый символ встречался только в одной части.
    // time: O(n), space: O(1)
    private static List<Integer> partitionLabels(String s) {
        // 1. Создаём мапу, где ключ — символ, значение — его последний индекс в строке
        Map<Character, Integer> lastIdx = new HashMap<>();
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            lastIdx.put(c, i);
        }

        int curLastIdx = 0;  // Текущая граница раздела
        List<Integer> result = new ArrayList<>(); // Результат (длины частей)
        int acc = 0;         // Суммарная длина предыдущих частей (для вычисления длины текущей)

        // 2. Проходим по строке, определяя границы частей
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            // Если последнее вхождение символа дальше текущей границы, расширяем границу
            if (lastIdx.get(c) > curLastIdx) {
                curLastIdx = lastIdx.get(c);
            }
            // Если дошли до границы — текущая часть завершена
            if (i == curLastIdx) {
                result.add(i + 1 - acc); // Добавляем длину части (i+1 - сумма предыдущих)
                acc = i + 1;             // Обновляем суммарную длину
            }
        }

        return result; // Возвращаем список целых чисел
    }
}
