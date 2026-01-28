package binarySearch.Sqrt;

/* 69. Sqrt(x)
https://leetcode.com/problems/sqrtx/description/

Если задано неотрицательное целое число x, вернуть квадратный корень x, округленный до ближайшего целого числа.
Возвращаемое целое число также должно быть неотрицательным.
Вы не должны использовать никакую встроенную функцию или оператор экспоненты.
Например, не используйте pow(x, 0.5) в c++ или x ** 0.5 в python.

Input: x = 4
Output: 2
Пояснение: Квадратный корень из 4 равен 2, поэтому мы возвращаем 2.
*/

public class Sqrt {
    public static void main(String[] args) {
        System.out.println(mySqrt(4)); // 2
    }

    // mySqrt возвращает квадратный корень x, округленный до ближайшего целого числа.
    //	time: O(log n), space: O(1)
    private static int mySqrt(int x) {
        if (x == 0) return 0;
        
        int left = 1;
        int right = x;

        while (left <= right) {
            int mid = left + (right - left) / 2;
            if (good(mid, x)) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }

        return right;
    }

    private static boolean good(int i, int x) {
        return (long) i * i <= x;
    }
}
