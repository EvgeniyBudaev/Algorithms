package leetcode.java.dynamicProgramming.RangeSumQueryImmutable;

// NumArray - класс для хранения массива чисел и эффективного вычисления суммы элементов в диапазоне.
public class NumArray {
    private int[] nums;      // Исходный массив чисел
    private int[] prefixSum; // Массив префиксных сумм для оптимизации

    // Временная сложность: O(n), где n - длина массива nums
    // Пространственная сложность: O(n) для хранения массива префиксных сумм
    public NumArray(int[] nums) {
        this.nums = nums;
        this.prefixSum = new int[nums.length + 1]; // Создаем массив префиксных сумм
        this.prefixSum[0] = 0;                      // Инициализируем первое значение в 0

        // Заполняем массив префиксных сумм
        for (int i = 0; i < nums.length; i++) {
            // Суммируем текущее число с предыдущей суммой префикса
            this.prefixSum[i + 1] = this.prefixSum[i] + nums[i];
        }
        // prefixSum [0, -2, -2, 1, -4, -2, -3]
    }

    // Временная сложность: O(1) - постоянное время, так как мы просто обращаемся к массиву префиксных сумм
    // Пространственная сложность: O(1) - постоянное пространство (дополнительная память метода)
    public int sumRange(int left, int right) {
        // Используем массив префиксных сумм для быстрого вычисления
        return this.prefixSum[right + 1] - this.prefixSum[left];
    }  
}
