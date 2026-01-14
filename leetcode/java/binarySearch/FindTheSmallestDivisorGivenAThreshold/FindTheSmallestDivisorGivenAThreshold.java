package binarySearch.FindTheSmallestDivisorGivenAThreshold;

/* 1283. Find the Smallest Divisor Given a Threshold
https://leetcode.com/problems/find-the-smallest-divisor-given-a-threshold/description/

Учитывая массив целых чисел nums и целочисленный threshold, мы выберем положительный целочисленный divisor, разделим на
него весь массив и просуммируем результат деления. Найдите наименьший divisor, чтобы упомянутый выше результат был
меньше или равен threshold.
Каждый результат деления округляется до ближайшего целого числа, большего или равного этому элементу.
(Например: 7/3 = 3 и 10/2 = 5).
Тестовые случаи генерируются для того, чтобы был ответ.

Input: nums = [1,2,5,9], threshold = 6
Output: 5
Пояснение: Мы можем получить сумму 17 (1+2+5+9), если делитель равен 1.
Если делитель равен 4, мы можем получить сумму 7 (1+1+2+3), а если делитель равен 5, сумма будет 5 (1+1+1+2).

Input: nums = [44,22,33,11,1], threshold = 5
Output: 44
*/

public class FindTheSmallestDivisorGivenAThreshold {
    public static void main(String[] args) {
        int[] nums = {1, 2, 5, 9};
        System.out.println(smallestDivisor(nums, 6)); // 5
    }

    // smallestDivisor находит наименьший делитель, который при делении всех элементов массива на него дает сумму, не превышающую порога.
    private static int smallestDivisor(int[] nums, int threshold) {
        int left = 1;
        int right = maxInArray(nums);

        while (left <= right) {
            int mid = left + (right - left) / 2;
            // Если сумма делителей не превышает порога
            if (sumDivisions(nums, mid) <= threshold) {
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }

        return left;
    }

    // sumDivisions вычисляет сумму округленных вверх частных
    private static int sumDivisions(int[] nums, int divisor) {
        int total = 0;
        for (int num : nums) {
            // Эквивалент ceil(num / divisor) без использования Math.ceil
            total += (num + divisor - 1) / divisor;
        }
        return total;
    }

    // maxInArray находит максимальное значение в срезе
    private static int maxInArray(int[] nums) {
        int maxNum = nums[0];
        for (int num : nums) {
            if (num > maxNum) {
                maxNum = num;
            }
        }
        return maxNum;
    }
}
