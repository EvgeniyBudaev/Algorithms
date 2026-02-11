package leetcode.java.hashing.MissingNumber;

/* 268. Missing Number
https://leetcode.com/problems/missing-number/description/

Учитывая массив nums, содержащий n различных чисел в диапазоне [0, n], верните единственное число в диапазоне,
которое отсутствует в массиве.

Input: nums = [3,0,1]
Output: 2
Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the
range since it does not appear in nums.

Input: nums = [0,1]
Output: 2
Explanation: n = 2 since there are 2 numbers, so all numbers are in the range [0,2]. 2 is the missing number in the
range since it does not appear in nums.

Input: nums = [9,6,4,2,3,5,7,0,1]
Output: 8
Explanation: n = 9 since there are 9 numbers, so all numbers are in the range [0,9]. 8 is the missing number in the
range since it does not appear in nums.
*/

public class MissingNumber {
    public static void main(String[] args) {
        int[] nums = {3, 0, 1};
        System.out.println(missingNumber(nums)); // 2
    }

    // missingNumber находит пропущенное число в последовательности [0, n]
    // time: O(n), space: O(1)
    private static int missingNumber(int[] nums) {
        int n = nums.length;                // Длина массива
        int expectedSum = n * (n + 1) / 2;  // Сумма чисел от 0 до n
        int actualSum = 0;                  // Сумма чисел в массиве
        
        // Суммируем числа в массиве
        for (int num : nums) {
            actualSum += num; // Суммируем число с текущей суммой
        }
        
        // Возвращаем разницу между ожидаемой и фактической суммой
        return expectedSum - actualSum;
    }
}
