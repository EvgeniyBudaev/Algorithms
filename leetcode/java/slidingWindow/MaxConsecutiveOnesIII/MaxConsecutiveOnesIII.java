package slidingWindow.MaxConsecutiveOnesIII;

/* 1004. Max Consecutive Ones III
https://leetcode.com/problems/max-consecutive-ones-iii/description/

Учитывая двоичный массив nums и целое число k, верните максимальное количество последовательных единиц в массиве,
если вы можете перевернуть не более k нулей.

Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
Output: 6
Explanation: [1,1,1,0,0,1,1,1,1,1,1]
Числа, выделенные жирным шрифтом, были перевернуты с 0 на 1. Самый длинный подмассив подчеркнут.

Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
Output: 10
Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
*/

public class MaxConsecutiveOnesIII {
    public static void main(String[] args) {
        int[] nums = {1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0};
        System.out.println(longestOnes(nums, 2)); // 6
    }

    // longestOnes возвращает максимальное количество последовательных единиц в массиве при условии, что
    // можно перевернуть не более k нулей.
    // time: O(n), space: O(1)
    private static int longestOnes(int[] nums, int k) {
        int left = 0, zeroCount = 0, result = 0;

        for (int right = 0; right < nums.length; right++) {
            // Если текущий элемент — ноль, увеличиваем счётчик нулей
            if (nums[right] == 0) {
                zeroCount++;
            }
            // Если количество нулей превышает допустимое значение k, сдвигаем левый указатель
            while (zeroCount > k) {
                if (nums[left] == 0) {
                    zeroCount--;
                }
                left++;
            }
            // Обновляем максимальную длину подмассива
            result = Math.max(result, right - left + 1);
        }

        return result; // Возвращаем максимальную длину подмассива
    }
}
