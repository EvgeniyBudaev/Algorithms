package slidingWindow.MaxConsecutiveOnesII;

/* 487. Max Consecutive Ones II
https://github.com/EvgeniyBudaev/doocs_leetcode/blob/main/solution/0400-0499/0487.Max%20Consecutive%20Ones%20II/README.md

Учитывая двоичный массив nums, если не более одного можно перевернуть 0, верните 1 максимальное количество
последовательных чисел в массиве.

Input: nums = [1,0,1,1,0]
Output: 4
Объяснение: переверните первый 0, чтобы получить самую длинную последовательную 1.
     После переворачивания максимальное количество последовательных единиц равно 4.

Input: nums = [1,0,1,1,0,1]
Output: 4
*/

public class MaxConsecutiveOnesII {
    public static void main(String[] args) {
        int[] nums = {1, 0, 1, 1, 0};
        System.out.println(findMaxConsecutiveOnes(nums)); // 4
    }

    // findMaxConsecutiveOnes возвращает максимальное количество последовательных 1 в массиве при условии, что
    // можно перевернуть 0 не более одного раза.
    // time: O(n), space: O(1)
    private static int findMaxConsecutiveOnes(int[] nums) {
        int left = 0, zeroCount = 0, result = 0, maxFlipOperations = 1;

        for (int right = 0; right < nums.length; right++) {
            // Если текущий элемент — ноль, увеличиваем счётчик нулей
            if (nums[right] == 0) {
                zeroCount++;
            }
            // Если количество нулей превышает разрешённое число переворотов, сдвигаем левый указатель
            while (zeroCount > maxFlipOperations) {
                if (nums[left] == 0) {
                    zeroCount--;
                }
                left++;
            }
            // Обновляем максимальную длину подмассива
            result = Math.max(result, right - left + 1);
        }

        return result; // Возвращаем максимальное количество последовательных 1
    }
}
