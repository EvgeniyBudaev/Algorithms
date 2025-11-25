package slidingWindow.FindLengthSubarray;

/*
Дан массив целых положительных чисел nums и целое число k. Найдите длину самого длинного подмассива, сумма которого
меньше или равна k.

Input: nums = [3, 1, 2, 7, 4, 2, 1, 1, 5], k = 8
Output: 4
*/

public class FindLengthSubarray {
    public static void main(String[] args) {
        int[] nums = {3, 1, 2, 7, 4, 2, 1, 1, 5};
        System.out.println(findLengthSubarray(nums, 8)); // 4
    }

    // findLengthSubarray находит длину самого длинного подмассива, сумма которого меньше или равна k.
    // time: O(n), space: O(1)
    private static int findLengthSubarray(int[] nums, int k) {
        int left = 0, sum = 0, result = 0; // Инициализируем левый указатель, сумму и результат

        for (int right = 0; right < nums.length; right++) { // Перебираем все элементы массива
            sum += nums[right]; // Добавляем текущий элемент в сумму

            // Пока сумма превышает k — сдвигаем левую границу
            while (sum > k) {
                sum -= nums[left]; // Убираем элемент из суммы
                left++; // Сдвигаем левый указатель
            }

            // Обновляем максимальную длину подмассива
            result = Math.max(result, right - left + 1);
        }

        return result; // Возвращаем длину самого длинного подмассива
    }
}
