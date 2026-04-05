package leetcode.java.dynamicProgramming.PartitionToKEqualSumSubsets;

/* 698. Partition to K Equal Sum Subsets
https://leetcode.com/problems/partition-to-k-equal-sum-subsets/description/

Учитывая целочисленный массив nums и целое число k, верните true, если можно разделить этот массив на k непустые
подмножества, все суммы которых равны.

Input: nums = [4,3,2,3,5,2,1], k = 4
Output: true
Explanation: It is possible to divide it into 4 subsets (5), (1, 4), (2,3), (2,3) with equal

Input: nums = [1,2,3,4], k = 3
Output: false
*/

public class PartitionToKEqualSumSubsets {
    public static void main(String[] args) {
        int[] nums = {4, 3, 2, 3, 5, 2, 1};
        System.out.println(canPartitionKSubsets(nums, 4)); // true ([5], [4,1], [3,2], [3,2])
    }

    // canPartitionKSubsets проверяет, можно ли разбить массив на K подмножеств с равными суммами
    // time: O(k * 2^n), где n - количество элементов, k - количество подмножеств
    // space: O(n) для массива visited и стека рекурсии
    private static boolean canPartitionKSubsets(int[] nums, int k) {
        // Вычисляем общую сумму элементов
        int total = 0;
        for (int num : nums) {
            total += num;
        }

        // Если сумма не делится на k - разбиение невозможно
        if (total % k != 0) {
            return false;
        }

        int target = total / k;
        boolean[] visited = new boolean[nums.length];

        return canPartition(nums, visited, 0, k, 0, target);
    }

    // canPartition рекурсивно пытается разбить массив на k подмножеств
    // time: O(2^n) в худшем случае для каждого вызова
    // space: O(n) для стека рекурсии
    private static boolean canPartition(int[] nums, boolean[] visited,
                                        int start, int numberOfSubsets,
                                        int currentSum, int target) {
        // Базовый случай: все подмножества найдены
        if (numberOfSubsets == 1) {
            return true;
        }

        // Нашли подмножество, ищем следующее
        if (currentSum == target) {
            return canPartition(nums, visited, 0, numberOfSubsets - 1, 0, target);
        }

        // Перебираем элементы, начиная с индекса start
        for (int i = start; i < nums.length; i++) {
            if (!visited[i] && currentSum + nums[i] <= target) {
                visited[i] = true;
                // Пробуем добавить текущий элемент в подмножество
                if (canPartition(nums, visited, i + 1, numberOfSubsets, currentSum + nums[i], target)) {
                    return true;
                }
                // Откатываем изменения (backtracking)
                visited[i] = false;
            }
        }

        return false;
    }
}
