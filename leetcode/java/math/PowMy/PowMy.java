package leetcode.java.math.PowMy;

/* 50. Pow(x, n)
https://leetcode.com/problems/powx-n/description/

Реализуйте pow(x, n), которая вычисляет x, возведенный в степень n (т. е. xn).

Input: x = 2.00000, n = 10
Output: 1024.00000
*/

public class PowMy {
    public static void main(String[] args) {
        System.out.println(myPow(2.0, 10)); // 1024.0
    }

    // myPow - метод вычисления степени числа.
    private static double myPow(double x, int n) {
        long N = n; // преобразование степени числа в long
        if (N < 0) {    // проверка на отрицательную степень числа
            x = 1 / x;  // деление на единицу
            N = -N;     // преобразование степени числа в положительное
        }
        return powHelper(x, N); // вызов рекурсивного метода для вычисления степени числа 
    }

    // powHelper - рекурсивный метод вычисления степени числа.
    // time: O(log n), space: O(log n)
    private static double powHelper(double x, long n) {
        if (n == 0) { // проверка на равенство степени числа нулю
            return 1; // возвращаем единицу
        }
        double half = powHelper(x, n / 2); // рекурсивный вызов метода для вычисления степени числа
        double result = half * half;       // возведение в квадрат
        if (n % 2 == 1) {                  // проверка на нечетность степени числа
            result *= x;                   // умножение на число
        }
        return result; // возвращаем результат
    }

    // Альтернативный вариант с использованием встроенной функции:
    // public static double myPow(double x, int n) {
    //     return Math.pow(x, n);
    // }
}
