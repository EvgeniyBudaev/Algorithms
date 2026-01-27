package binarySearch.SplitArrayLargestSum;

/* 410. Split Array Largest Sum

Учитывая целочисленный массив nums и целое число k, разбейте числа на k непустые подмассивы так, чтобы наибольшая сумма
любого подмассива была минимизирована.
Верните минимизированную наибольшую сумму разделения.
Подмассив — это непрерывная часть массива.

Input: nums = [7,2,5,10,8], k = 2
Output: 18
Объяснение: Существует четыре способа разбить числа на два подмассива.
Лучший способ — разделить его на [7,2,5] и [10,8], где наибольшая сумма среди двух подмассивов равна всего 18.

Input: nums = [1,2,3,4,5], k = 2
Output: 9
Объяснение: Существует четыре способа разбить числа на два подмассива.
Лучший способ — разбить его на [1,2,3] и [4,5], где наибольшая сумма среди двух подмассивов равна всего 9.
*/

public class SplitArrayLargestSum {
    public static void main(String[] args) {
        int[] nums = {7, 2, 5, 10, 8};
        System.out.println(splitArray(nums, 2)); // 18
    }

    private static int splitArray(int[] nums, int k) {
        int left = getMax(nums);
        int right = sum(nums);
        int ans = -1;

        while (left <= right) {
            int largestSum = (left + right) / 2;

            if (countSplitSubarray(largestSum, nums) <= k) {
                ans = largestSum;
                right = largestSum - 1;
            } else {
                left = largestSum + 1;
            }
        }

        return ans;
    }

    private static int countSplitSubarray(int largestSum, int[] nums) {
        int countSubarray = 1;
        int currentSum = 0;

        for (int num : nums) {
            currentSum += num;

            if (currentSum > largestSum) {
                countSubarray++;
                currentSum = num;
            }
        }

        return countSubarray;
    }

    private static int getMax(int[] nums) {
        int maxNum = Integer.MIN_VALUE;
        for (int num : nums) {
            if (num > maxNum) {
                maxNum = num;
            }
        }
        return maxNum;
    }

    private static int sum(int[] nums) {
        int sum = 0;
        for (int num : nums) {
            sum += num;
        }
        return sum;
    }
}
