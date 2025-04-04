package main

import (
	"fmt"
	"math"
)

/*
https://leetcode.com/problems/min-cost-climbing-stairs/description/

Вам дан целочисленный массив Cost, где Cost[i] — стоимость i-й ступеньки лестницы. Заплатив стоимость, вы сможете
 подняться на одну или две ступеньки.
Вы можете начать либо с шага с индексом 0, либо с шага с индексом 1.
Возвращает минимальную стоимость достижения вершины этажа.

Input: cost = [10,15,20]
Output: 15
Пояснение: Вы начнете с индекса 1.
- Заплатите 15 и поднимитесь на две ступеньки, чтобы достичь вершины.
Общая стоимость 15.

Input: cost = [1,100,1,1,1,100,1,1,100,1]
Output: 6
Объяснение: Вы начнете с индекса 0.
- Заплатите 1 и поднимитесь на две ступеньки, чтобы достичь индекса 2.
- Заплатите 1 и поднимитесь на две ступеньки, чтобы достичь индекса 4.
- Заплатите 1 и поднимитесь на две ступеньки, чтобы достичь индекса 6.
- Заплатите 1 и поднимитесь на одну ступень, чтобы достичь индекса 7.
- Заплатите 1 и поднимитесь на две ступеньки, чтобы достичь индекса 9.
- Заплатите 1 и поднимитесь на одну ступеньку, чтобы достичь вершины.
Общая стоимость 6.
*/

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20})) // 15 (шаг 1 → вершина)
}

// minCostClimbingStairs вычисляет минимальную стоимость достижения вершины лестницы
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return cost[0]
	}

	// Создаем копию массива, чтобы не изменять исходный
	dp := make([]int, n)
	copy(dp, cost)

	for i := 2; i < n; i++ {
		// Вычисляем минимальную стоимость достижения текущей ступеньки:
		// минимум из предыдущих двух ступенек + стоимость текущей
		dp[i] = int(math.Min(float64(dp[i-1]), float64(dp[i-2]))) + dp[i]
	}

	// Возвращаем минимум из последних двух ступенек,
	// так как можно прыгнуть на вершину с любой из них
	return int(math.Min(float64(dp[n-1]), float64(dp[n-2])))
}
