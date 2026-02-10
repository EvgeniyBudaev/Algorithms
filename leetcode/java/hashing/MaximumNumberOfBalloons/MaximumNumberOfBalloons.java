package leetcode.java.hashing.MaximumNumberOfBalloons;

import java.util.HashMap;
import java.util.Map;

/* 1189. Maximum Number of Balloons
https://leetcode.com/problems/maximum-number-of-balloons/description/

Учитывая строковый текст, вы хотите использовать символы текста для формирования как можно большего количества
экземпляров слова «balloon».
Каждый символ в тексте можно использовать не более одного раза. Верните максимальное количество экземпляров,
которые можно сформировать.

Input: text = "nlaebolko"
Output: 1

Input: text = "loonbalxballpoon"
Output: 2

Input: text = "leetcode"
Output: 0
*/

public class MaximumNumberOfBalloons {
    public static void main(String[] args) {
        System.out.println(maxNumberOfBalloons("nlaebolko")); // 1
    }

    // maxNumberOfBalloons подсчитывает, сколько раз можно составить слово "balloon" из букв строки.
    // time: O(n), space: O(1)
    private static int maxNumberOfBalloons(String text) {
        // Создаем мапу для подсчета нужных букв
        Map<Character, Double> counts = new HashMap<>();
        counts.put('b', 0.0);
        counts.put('a', 0.0);
        counts.put('l', 0.0);
        counts.put('o', 0.0);
        counts.put('n', 0.0);

        // Подсчитываем только те буквы, которые входят в слово "balloon"
        for (char c : text.toCharArray()) {
            if (counts.containsKey(c)) { // Если буква есть в мапе, то увеличиваем ее счетчик
                counts.put(c, counts.get(c) + 1.0);
            }
        }

        // Для букв 'l' и 'o' нужно 2 экземпляра на каждое слово "balloon",
        // поэтому делим их количество на 2
        counts.put('l', counts.get('l') / 2.0);
        counts.put('o', counts.get('o') / 2.0);

        // Находим минимальное значение среди всех требуемых букв
        double minValue = Double.POSITIVE_INFINITY; // Начинаем с положительной бесконечности
        for (double count : counts.values()) {
            minValue = Math.min(minValue, count); // Находим минимальное значение
        }

        // Возвращаем целую часть минимального значения (округляем вниз)
        return (int) Math.floor(minValue);
    }
}
