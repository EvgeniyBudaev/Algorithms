package slidingWindow.SubarrayProductLessThanK;

/* 713. Subarray Product Less Than K
https://leetcode.com/problems/subarray-product-less-than-k/description/

Учитывая массив положительных целых чисел nums и целое число k, верните количество подмассивов, в которых произведение
всех элементов в подмассиве строго меньше k.

Input: nums = [10,5,2,6], k = 100
Output: 8
Пояснение: 8 подмассивов с произведением меньше 100:
[10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6]
Обратите внимание, что [10, 5, 2] не включены, поскольку произведение 100 не строго меньше k.

Input: nums = [1,2,3], k = 0
Output: 0
*/

public class SubarrayProductLessThanK {
    public static void main(String[] args) {
        int[] nums = {10, 5, 2, 6};
        System.out.println(numSubarrayProductLessThanK(nums, 100)); // 8
    }

    // numSubarrayProductLessThanK возвращает количество подмассивов, в которых произведение всех элементов в подмассиве строго меньше k.
    // time: O(n), space: O(1)
    private static int numSubarrayProductLessThanK(int[] nums, int k) {
        // Если k <= 1, то нет подмассивов, в которых произведение всех элементов в подмассиве строго меньше k.
        if (k <= 1) {
            return 0;
        }

        int left = 0, curr = 1, result = 0; // left - левая граница окна, curr - текущее произведение, result - количество подмассивов

        for (int right = 0; right < nums.length; right++) {
            curr *= nums[right]; // Увеличиваем текущее произведение на текущий элемент

            // Если произведение превышает или равно k, сдвигаем левую границу окна
            while (curr >= k) {
                curr /= nums[left];
                left++;
            }

            // Добавляем количество подмассивов, заканчивающихся на right
            result += right - left + 1;
        }

        return result; // Возвращаем количество подмассивов
    }
}
