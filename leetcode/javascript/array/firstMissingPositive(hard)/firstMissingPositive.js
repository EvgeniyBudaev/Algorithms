/* 41. First Missing Positive
https://leetcode.com/problems/first-missing-positive/description/

Дан несортированный целочисленный массив nums. Верните наименьшее положительное целое число, которого нет в nums.
Вы должны реализовать алгоритм, который работает за время O(n) и использует вспомогательное пространство O(1).

Input: nums = [1,2,0]
Output: 3
Объяснение: Все числа в диапазоне [1,2] входят в массив

Input: nums = [3,4,-1,1]
Output: 2
Пояснение: 1 есть в массиве, но 2 отсутствует.

Input: nums = [7,8,9,11,12]
Output: 1
Объяснение: Пропущено наименьшее положительное целое число 1.
*/

/**
 * @param {number[]} nums
 * @return {number}
 * @time O(n)
 * @space O(1)
 * @description Cyclic Sort (циклическая сортировка)
 */
var firstMissingPositive = function(nums) {
  const n = nums.length;
  
  // Шаг 1: Расставляем числа на свои места
  for (let i = 0; i < n; i++) {
    // Пока текущее число находится в диапазоне [1, n] 
    // и оно не стоит на своем правильном месте (nums[i] должно быть по индексу nums[i] - 1)
    // и чтобы избежать бесконечного цикла при дубликатах (nums[nums[i] - 1] !== nums[i])
    while (
      nums[i] > 0 && 
      nums[i] <= n && 
      nums[nums[i] - 1] !== nums[i]
    ) {
      // Меняем местами nums[i] и число на его "законной" позиции
      const correctIdx = nums[i] - 1;
      const temp = nums[correctIdx];
      nums[correctIdx] = nums[i];
      nums[i] = temp;
    }
  }
  
  // Шаг 2: Ищем первое число, которое не стоит на своем месте
  for (let i = 0; i < n; i++) {
    if (nums[i] !== i + 1) {
      return i + 1;
    }
  }
  
  // Если все числа от 1 до n присутствуют, значит пропущено n + 1
  return n + 1;
};

console.log(firstMissingPositive([1,2,0])); // 3