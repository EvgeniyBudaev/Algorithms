package leetcode.java.dynamicProgramming.CountingBits;

import java.util.Arrays;

/* 338. Counting Bits
https://leetcode.com/problems/counting-bits/description/

Учитывая целое число n, верните массив ans длины n + 1 такой, что для каждого i (0 <= i <= n) ans[i] — это количество
единиц в двоичном представлении i.

Input: n = 2
Output: [0,1,1]
Explanation:
0 --> 0
1 --> 1
2 --> 10
*/

public class CountingBits {
    public static void main(String[] args) {
        System.out.println(Arrays.toString(countBits(2))); // [0, 1, 1]
    }

    // countBits возвращает массив, где каждый элемент содержит количество 1 в двоичном представлении числа
    // time: O(n), space: O(n)
    private static int[] countBits(int n) {
        // Инициализируем массив результатов с нулями
        int[] dp = new int[n + 1];

        // offset - текущая степень двойки
        int offset = 1;

        for (int i = 1; i <= n; i++) {
            // Если достигли новой степени двойки, обновляем offset
            if (offset * 2 == i) {
                offset = i;
            }
            // Количество единиц = 1 (для нового старшего бита) + количество в оставшейся части
            dp[i] = 1 + dp[i - offset];
        }

        return dp;
    }
}
