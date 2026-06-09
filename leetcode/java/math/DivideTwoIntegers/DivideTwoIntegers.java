package leetcode.java.math.DivideTwoIntegers;

/* 29. Divide Two Integers
https://leetcode.com/problems/divide-two-integers/description/

Даны два целых числа: делимое и делитель. Разделите два целых числа без использования умножения, деления и оператора mod.

Целочисленное деление должно усекаться в сторону нуля, что означает потерю дробной части. Например, 8,345 будет усечено
до 8, а -2,7335 будет усечено до -2.

Верните частное после деления делимого на делитель.

Примечание: предположим, что мы имеем дело со средой, которая может хранить только целые числа в диапазоне 32-битных
знаковых целых чисел: [−231, 231 − 1]. Для этой задачи, если частное строго больше 231 - 1, то верните 231 - 1,
а если частное строго меньше -231, то верните -231.

Input: dividend = 10, divisor = 3
Output: 3
Пояснение: 10/3 = 3,33333, что усекается до 3.
*/

public class DivideTwoIntegers {
    public static void main(String[] args) {
        System.out.println(divide(10, 3)); // 3
    }

    // divide - деление целых чисел.
    // time: O(log(dividend)), space: O(1)
    private static int divide(int dividend, int divisor) {
        // Если делимое равно Integer.MIN_VALUE, а делитель -1, то результат будет больше Integer.MAX_VALUE.
        if (dividend == Integer.MIN_VALUE && divisor == -1) {
            return Integer.MAX_VALUE;
        }

        boolean negative = false; // Результат будет отрицательным?
        // Если делимое и делитель имеют разные знаки, то результат будет отрицательным.
        if ((dividend < 0) != (divisor < 0)) {
            negative = true;
        }

        // В Java int всегда 32-битный. Чтобы избежать переполнения при взятии модуля от Integer.MIN_VALUE,
        // мы предварительно приводим значения к long.
        long dvd = abs(dividend);
        long dvs = abs(divisor);

        long result = 0;                // Результат деления.
        while (dvd >= dvs) {            // Пока делимое больше или равно делителю.
            long power = 1;             // Степень двойки.
            long val = dvs;             // Значение делителя.
            while (val + val <= dvd) {  // Пока значение делителя меньше или равно делимому.
                val += val;             // Увеличиваем значение делителя.
                power += power;         // Увеличиваем степень двойки.
            }

            dvd -= val;                 // Вычитаем значение делителя из делимого.
            result += power;            // Добавляем степень двойки к результату.
        }

        if (negative) {                 // Если результат должен быть отрицательным, то умножаем его на -1.
            result = -result;           // Умножаем на -1.
        }

        // Примечание: эти проверки фактически избыточны, так как единственный случай переполнения 
        // (Integer.MIN_VALUE / -1) мы уже обработали в самом начале функции. 
        // Но мы оставляем их для полного соответствия оригинальному коду.
        if (result > Integer.MAX_VALUE) { // Если результат больше максимального значения.
            return Integer.MAX_VALUE;
        }

        if (result < Integer.MIN_VALUE) { // Если результат меньше минимального значения.
            return Integer.MIN_VALUE;
        }

        return (int) result;            // Возвращаем результат, приводя его обратно к 32-битному int.
    }

    // abs - модуль числа.
    // В Java мы принимаем long, чтобы безопасно обработать Integer.MIN_VALUE без переполнения.
    // time: O(1), space: O(1)
    private static long abs(long x) {
        if (x < 0) {
            return -x;
        }
        return x;
    }
}
