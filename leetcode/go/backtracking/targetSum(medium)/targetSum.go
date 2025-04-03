package main

import (
	"fmt"
	"strconv"
)

/* https://leetcode.com/problems/target-sum/description/

Вам дан целочисленный массив nums и целочисленная цель.
Вы хотите построить выражение из чисел, добавив один из символов «+» и «-» перед каждым целым числом в числах, а затем
объединить все целые числа.
Например, если nums = [2, 1], вы можете добавить «+» перед 2 и «-» перед 1 и объединить их, чтобы построить выражение
«+2-1».
Верните количество различных выражений, которые вы можете создать и которые достигают целевого значения.

Input: nums = [1,1,1,1,1], target = 3
Output: 5
Пояснение: Есть 5 способов назначить символы, чтобы сумма чисел соответствовала целевому значению 3.
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3

Input: nums = [1], target = 1
Output: 1
*/

func main() {
	nums := []int{1, 1, 1, 1, 1}
	fmt.Println(findTargetSumWays(nums, 3)) // 5
}

// findTargetSumWays возвращает количество способов получить целевую сумму
func findTargetSumWays(nums []int, target int) int {
	memo := make(map[string]int)
	return dp(nums, target, 0, 0, memo)
}

// dp - рекурсивная функция с мемоизацией для подсчета способов
func dp(nums []int, target, index, sum int, memo map[string]int) int {
	// Создаем ключ для мемоизации
	key := strconv.Itoa(index) + "_" + strconv.Itoa(sum)

	// Проверяем, есть ли уже результат в кэше
	if val, ok := memo[key]; ok {
		return val
	}

	// Базовый случай: дошли до конца массива
	if index == len(nums) {
		if sum == target {
			return 1
		}
		return 0
	}

	// Рекурсивно считаем варианты с "+" и "-" текущего числа
	positive := dp(nums, target, index+1, sum+nums[index], memo)
	negative := dp(nums, target, index+1, sum-nums[index], memo)

	// Суммируем оба варианта
	total := positive + negative

	// Сохраняем результат в кэш
	memo[key] = total

	return total
}
