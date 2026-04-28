package leetcode.java.heaps.KthLargestElementInAnArray;

import java.util.PriorityQueue;

/* 215. Kth Largest Element in an Array
https://leetcode.com/problems/kth-largest-element-in-an-array/description/

Дан целочисленный массив nums и целое число k, вернуть k-й наибольший элемент в массиве.
Обратите внимание, что это k-й наибольший элемент в отсортированном порядке, а не k-й отдельный элемент.
Можете ли вы решить эту задачу без сортировки?

Input: nums = [3,2,1,5,6,4], k = 2
Output: 5
*/

public class KthLargestElementInAnArray {
    public static void main(String[] args) {
        int[] nums = {3, 2, 1, 5, 6, 4};
        int k = 2;
        System.out.println(findKthLargest(nums, k)); // 5
    }

    /**
     * findKthLargest - находит k-й наибольший элемент в массиве.
     * time: O(n*log(k)), space: O(k)
     * <p>
     * Алгоритм использует мин-кучу размера k:
     * - Поддерживаем в куче только k наибольших элементов
     * - Наименьший из этих k элементов будет корнем кучи
     * - После обработки всех элементов корень = k-й наибольший
     */
    private static int findKthLargest(int[] nums, int k) {
        // PriorityQueue в Java по умолчанию реализует мин-кучу
        PriorityQueue<Integer> minHeap = new PriorityQueue<>();

        for (int num : nums) {
            minHeap.offer(num);           // добавляем элемент в кучу, O(log k)
            if (minHeap.size() > k) {     // если размер превысил k
                minHeap.poll();           // удаляем наименьший элемент, O(log k)
            }
        }

        // Корень кучи — это k-й наибольший элемент
        return minHeap.poll();
    }
}
