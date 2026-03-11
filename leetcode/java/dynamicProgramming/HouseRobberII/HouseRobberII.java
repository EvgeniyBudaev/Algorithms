package leetcode.java.dynamicProgramming.HouseRobberII;

/* 213. House Robber II
https://leetcode.com/problems/house-robber-ii/description/

Вы профессиональный грабитель, планирующий грабить дома на улице. В каждом доме хранится определенная сумма денег.
Все дома в этом месте расположены по кругу. Это означает, что первый дом является соседом последнего.
При этом в соседних домах подключена система безопасности, и она автоматически обратится в полицию, если в одну ночь
были взломаны два соседних дома.
Учитывая целочисленный массив чисел, представляющий сумму денег в каждом доме, верните максимальную сумму денег,
которую вы можете ограбить сегодня вечером, не предупредив полицию.

Input: nums = [2,3,2]
Output: 3
Пояснение: Вы не можете ограбить дом 1 (деньги = 2), а затем ограбить дом 3 (деньги = 2), потому что это соседние дома

Input: nums = [1,2,3,1]
Output: 4
Пояснение: Ограбить дом 1 (деньги = 1), а затем ограбить дом 3 (деньги = 3).
Общая сумма, которую вы можете ограбить = 1 + 3 = 4.

Input: nums = [1,2,3]
Output: 3
*/

public class HouseRobberII {
    public static void main(String[] args) {
        System.out.println(rob(new int[]{2, 3, 2})); // 3
    }

    // rob вычисляет максимальную сумму для ограбления домов на кольцевой улице
    // time: O(n), space: O(1)
    private static int rob(int[] nums) {
        int n = nums.length;
        // Обрабатываем особые случаи для массивов длиной менее 4
        if (n == 0) {
            return 0;
        }
        if (n < 4) {
            return maxInSlice(nums);
        }

        // Возвращаем максимум из двух вариантов:
        // 1. Грабим первый дом, но не последний
        // 2. Грабим последний дом, но не первый
        return Math.max(
            robberHelper(nums, 0, n - 1),
            robberHelper(nums, 1, n)
        );
    }

    // robberHelper возвращает максимальную сумму на отрезке
    private static int robberHelper(int[] nums, int start, int end) {
        int prev = 0;
        int beforePrev = 0;
        for (int i = start; i < end; i++) {
            // Сохраняем предыдущее значение
            int tmp = prev;
            // Вычисляем новое значение как максимум:
            // - ограбление текущего дома + beforePrev
            // - или пропуск текущего дома (prev)
            prev = Math.max(nums[i] + beforePrev, prev);
            beforePrev = tmp;
        }
        return prev;
    }

    // maxInSlice возвращает максимальное значение в массиве
    private static int maxInSlice(int[] nums) {
        int maxNum = nums[0];
        for (int num : nums) {
            if (num > maxNum) {
                maxNum = num;
            }
        }
        return maxNum;
    }
}
