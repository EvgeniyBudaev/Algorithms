package array.MajorityElement;

/* 169. Majority Element
https://leetcode.com/problems/majority-element/description/

Учитывая массив nums размера n, верните элемент большинства.
Мажоритарным элементом является элемент, который появляется более ⌊n / 2⌋ раз. Вы можете предположить, что элемент
большинства всегда существует в массиве.

Input: nums = [3,2,3]
Output: 3

Input: nums = [2,2,1,1,1,2,2]
Output: 2

Input: nums = [3,3,4]
Output: 3
*/

public class MajorityElement {
    public static void main(String[] args) {
        int[] nums = {3, 2, 3};
        System.out.println(majorityElement(nums)); // 3
    }

    // majorityElement возвращает элемент большинства.
    // time: O(n), space: O(1)
    private static int majorityElement(int[] nums) {
        int count = 0;     // Количество элементов, которые мы считаем пока не нашли кандидата.
        int candidate = 0; // Кандидат на элемент большинства.

        for (int i = 0; i < nums.length; i++) {
            // Если количество элементов равно нулю, то текущий элемент становится кандидатом на элемент большинства.
            if (count == 0) {
                candidate = nums[i];
            }
            // Если текущий элемент совпадает с кандидатом, увеличиваем количество элементов, иначе уменьшаем количество.
            if (candidate == nums[i]) {
                count++;
                // Если количество элементов больше половины длины массива, то текущий элемент является элементом большинства.
            } else {
                count--;
            }
        }

        return candidate; // Возвращаем кандидата на элемент большинства.
    }
}
