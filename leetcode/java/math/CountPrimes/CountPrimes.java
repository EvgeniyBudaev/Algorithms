package leetcode.java.math.CountPrimes;

/* 204. Count Primes
https://leetcode.com/problems/count-primes/description/

Для заданного целого числа n вернуть количество простых чисел, которые строго меньше n.

Input: n = 10
Output: 4
Пояснение: Существует 4 простых числа, меньших 10: 2, 3, 5, 7.
*/

public class CountPrimes {
    public static void main(String[] args) {
        System.out.println(countPrimes(10)); // 4
    }

    // countPrimes - реализация решета Эратосфена.
    // time: O(n*log(log(n))), space: O(n)
    private static int countPrimes(int n) {
        // Если число меньше или равно 2, то нет простых чисел.
        if (n <= 2) {
            return 0;
        }

        // Массив, где каждый элемент обозначает, является ли число составным или нет.
        // В Java boolean массив по умолчанию инициализируется значениями false.
        boolean[] isComposite = new boolean[n];
        int count = 0; // Счетчик простых чисел

        // Перебираем все числа от 2 до n-1
        for (int i = 2; i < n; i++) {
            // Если число не является составным, то увеличиваем счетчик простых чисел 
            // и отмечаем все его кратные числа как составные.
            if (!isComposite[i]) {
                count++; // Увеличиваем счетчик простых чисел

                // Отмечаем все кратные числа как составные.
                for (int j = i * 2; j < n; j += i) {
                    isComposite[j] = true; // Отмечаем число как составное.
                }
            }
        }

        // Возвращаем количество простых чисел, меньших n.
        return count;
    }
}
