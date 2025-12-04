package slidingWindow.MaxConsecutiveOnes;

/* 485. Max Consecutive Ones
https://leetcode.com/problems/max-consecutive-ones/description/

Для двоичного массива nums вернуть максимальное количество последовательных 1 в массиве.

Example 1
Input: nums = [1,1,0,1,1,1]
Output: 3

Example 2
Input: nums = [1,0,1,1,0,1]
Output: 2
*/

public class MaxConsecutiveOnes {
    public static void main(String[] args) {
        int[] nums = {1, 1, 0, 1, 1, 1};
        System.out.println(findMaxConsecutiveOnes(nums)); // 3
    }

    // findMaxConsecutiveOnes возвращает максимальное количество последовательных 1 в массиве.
    // time: O(n), space: O(1)
    private static int findMaxConsecutiveOnes(int[] nums) {
        int left = 0, zeroCount = 0, result = 0;
        
        for (int right = 0; right < nums.length; right++) {
            // Если ноль, увеличиваем счетчик подсчета нулей
            if (nums[right] == 0) {
                zeroCount++;
            }
            // Если количество нулей больше 0, двигаем левый указатель
            while (zeroCount > 0) {
                if (nums[left] == 0) {
                    zeroCount--;
                }
                left++;
            }
            // Обновляем максимальную длину последовательных единиц
            result = Math.max(result, right - left + 1);
        }
        
        return result; // Возвращаем максимальную длину последовательных единиц
    }
}
