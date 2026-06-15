package leetcode.java.math.PowerOfThree;

/* 326. Power of Three
https://leetcode.com/problems/power-of-three/description/

Дано целое число n, вернуть true, если оно является степенью трех. В противном случае вернуть false.
Целое число n является степенью трех, если существует целое число x, такое что n == 3x.

Input: n = 27
Output: true
Explanation: 27 = 33
*/

public class PowerOfThree {
    public static void main(String[] args) {
        System.out.println(isPowerOfThree(27)); // true
    }

    // isPowerOfThree возвращает true, если число является степенью трех, иначе false
    // time: O(log3 n), space: O(1)
    private static boolean isPowerOfThree(int n) {
        if (n < 1) { // Если число меньше единицы, то оно не может быть степенью трех
            return false;
        }

        while (n % 3 == 0) { // Если число делится на 3 без остатка, то оно является степенью трех
            n /= 3;          // Делим число на 3 пока не останется 1
        }

        return n == 1;       // Если число равно 1, то оно является степенью трех
    }
}
