package leetcode.java.design.ShuffleAnArray;

import java.util.Random;

/* 384. Shuffle an Array
https://leetcode.com/problems/shuffle-an-array/description/

Имея целочисленный массив nums, разработайте алгоритм для случайного перемешивания массива. Все перестановки массива должны быть одинаково вероятны в результате перемешивания.

Реализуйте класс Solution:

Solution(int[] nums) Инициализирует объект целочисленным массивом nums.
int[] reset() Сбрасывает массив в исходную конфигурацию и возвращает его.
int[] shuffle() Возвращает случайное перемешивание массива.

Input
["Solution", "shuffle", "reset", "shuffle"]
[[[1, 2, 3]], [], [], []]
Output
[null, [3, 1, 2], [1, 2, 3], [1, 3, 2]]
*/

public class ShuffleAnArray {
    private int[] orig; // исходный массив
    private int[] copy; // копия исходного массива
    private Random rand;

    // Constructor - инициализирует объект целочисленным массивом nums.
    // time: O(n), space: O(n)
    public ShuffleAnArray(int[] nums) {
        this.orig = nums;           // исходный массив
        this.copy = nums.clone();   // копия исходного массива
        this.rand = new Random();   // инициализация генератора случайных чисел
    }
    
    // reset - сбрасывает массив в исходную конфигурацию и возвращает его.
    // time: O(n), space: O(1)
    public int[] reset() {
        return this.orig; // возвращаем исходный массив
    }
    
    // shuffle - возвращает случайное перемешивание массива.
    // time: O(n), space: O(n)
    public int[] shuffle() {
        // перемешиваем копию исходного массива (алгоритм Фишера-Йетса)
        for (int i = this.copy.length - 1; i > 0; i--) {
            int j = rand.nextInt(i + 1); // генерируем случайный индекс от 0 до i
            // меняем элементы местами
            int temp = this.copy[i];
            this.copy[i] = this.copy[j];
            this.copy[j] = temp;
        }

        return this.copy; // возвращаем перемешанный массив
    }
}
