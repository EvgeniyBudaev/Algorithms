package leetcode.java.hashing.JewelsAndStones;

import java.util.HashSet;
import java.util.Set;

/* 771. Jewels and Stones
https://leetcode.com/problems/jewels-and-stones/description/

Вам даны строки jewels, представляющие типы камней, которые являются jewels, и stones,
представляющие те камни, которые у вас есть. Каждый символ в stones — это тип камня, который у вас есть.
Вы хотите знать, сколько stones у вас тоже являются jewels.
Буквы чувствительны к регистру, поэтому «а» считается другим типом камня, чем «А».

Input: jewels = "aA", stones = "aAAbbbb"
Output: 3

Input: jewels = "z", stones = "ZZ"
Output: 0
*/

public class JewelsAndStones {
    public static void main(String[] args) {
        
    }

    // numJewelsInStones получает количество камней из набора камней в наборе камней.
    // time: O(n), space: O(n)
    private static int numJewelsInStones(String jewels, String stones) {
        Set<Character> jewelSet = new HashSet<>(); // Множество камней
        for (char j : jewels.toCharArray()) { // Создаем набор камней для поиска O(1)
            jewelSet.add(j);
        }

        int count = 0; // Счетчик камней
        for (char s : stones.toCharArray()) { // Считаем сколько камней в наборе
            if (jewelSet.contains(s)) { // Если камень есть в наборе, то увеличиваем счетчик
                count++;
            }
        }

        return count; // Возвращаем счетчик камней
    }
}
