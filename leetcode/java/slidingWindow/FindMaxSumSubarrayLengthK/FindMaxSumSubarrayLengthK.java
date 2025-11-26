package slidingWindow.FindMaxSumSubarrayLengthK;

/* Fixed window size

Учитывая целочисленный массив nums и целое число k, найдите сумму подмассива с наибольшей суммой, длина которой равна k.
*/

public class FindMaxSumSubarrayLengthK {
    public static void main(String[] args) {
        int[] nums = {3, -1, 4, 12, -8, 5, 6};
        System.out.println(findMaxSumSubarrayLengthK(nums, 4)); // 18
    }

    // findMaxSumSubarrayLengthK находит сумму подмассива с наибольшей суммой, длина которой равна k.
    // time: O(n), space: O(1)
    private static int findMaxSumSubarrayLengthK(int[] nums, int k) {
        if (nums == null || k <= 0 || k > nums.length) {
            throw new IllegalArgumentException("Invalid input: k must be positive and not greater than array length");
        }

        int sum = 0; // Текущая сумма подмассива длины k

        // Вычисляем сумму первого окна длины k
        for (int i = 0; i < k; i++) {
            sum += nums[i];
        }

        int result = sum; // Инициализируем ответ суммой первого окна

        // Скользим окном по массиву
        for (int i = k; i < nums.length; i++) {
            sum += nums[i] - nums[i - k]; // Добавляем nums[i] и удаляем nums[i-k]
            result = Math.max(result, sum); // Обновляем result, если текущая сумма больше
        }

        return result; // Возвращаем наибольшую сумму подмассива
    }
}
