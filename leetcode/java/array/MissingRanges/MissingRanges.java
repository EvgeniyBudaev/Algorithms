package array.MissingRanges;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/* 163. Missing Ranges
https://leetcode.ca/2016-05-11-163-Missing-Ranges/

Вам дан инклюзивный диапазон [нижний, верхний] и отсортированный уникальный целочисленный массив nums, где все элементы
находятся в пределах инклюзивного диапазона.

Число x считается отсутствующим, если x находится в диапазоне [нижний, верхний] и x не находится в nums.

Верните кратчайший отсортированный список диапазонов, который точно охватывает все отсутствующие числа. То есть ни один
элемент nums не включен ни в один из диапазонов, и каждое отсутствующее число покрывается одним из диапазонов.

Input: nums = [0,1,3,50,75], lower = 0, upper = 99
Output: [[2,2],[4,49],[51,74],[76,99]]
Explanation: The ranges are:
[2,2]
[4,49]
[51,74]
[76,99]
*/

public class MissingRanges {
    public static void main(String[] args) {
        int[] nums = {0, 1, 3, 50, 75};
        System.out.println(findMissingRanges(nums, 0, 99)); // [[2,2], [4,49], [51,74], [76,99]]
    }

    // findMissingRanges возвращает кратчайший отсортированный список диапазонов, который точно охватывает все отсутствующие числа.
    // Time: O(n), Space: O(n)
    private static List<List<Integer>> findMissingRanges(int[] nums, int lower, int upper) {
        int n = nums.length; // количество чисел в массиве
        List<List<Integer>> result = new ArrayList<>();

        if (n == 0) { // если массив пустой
            result.add(Arrays.asList(lower, upper)); // вернуть диапазон от нижнего до верхнего
            return result;
        }

        if (nums[0] > lower) { // если первый элемент массива больше нижнего
            result.add(Arrays.asList(lower, nums[0] - 1)); // добавить диапазон от нижнего до первого элемента массива минус 1
        }

        for (int i = 0; i < n - 1; i++) { // проход по массиву от первого до предпоследнего элемента
            int a = nums[i];
            int b = nums[i + 1];
            if (b - a > 1) { // если разница между текущим и следующим элементами больше 1
                result.add(Arrays.asList(a + 1, b - 1)); // добавить диапазон от текущего элемента массива плюс 1 до следующего минус 1
            }
        }

        if (nums[n - 1] < upper) { // если последний элемент массива меньше верхнего
            result.add(Arrays.asList(nums[n - 1] + 1, upper)); // добавить диапазон от последнего элемента массива плюс 1 до верхнего
        }

        return result; // вернуть список диапазонов
    }
}
