package main

import (
	"fmt"
	"sort"
)

/* 1710. Maximum Units on a Truck

Вам поручено погрузить некоторое количество коробок в один грузовик. Вам дан двумерный массив boxTypes,
где boxTypes[i] = [numberOfBoxesi, NumberOfUnitsPerBoxi]:

NumberOfBoxesi — количество ящиков типа i.
NumberOfUnitsPerBoxi — количество единиц в каждом ящике типа i.
Вам также дается целое число TruckSize, которое представляет собой максимальное количество коробок, которые можно
разместить в грузовике. Вы можете выбрать любые ящики для установки в грузовик, если их количество не превышает
TruckSize.
Верните максимальное общее количество единиц, которое можно разместить на грузовике.

Input: boxTypes = [[1,3],[2,2],[3,1]], truckSize = 4
Output: 8
Пояснение: Есть:
- 1 коробка первого типа, содержащая 3 единицы.
- 2 коробки второго типа по 2 единицы в каждой.
- 3 коробки третьего типа по 1 единице в каждой.
Вы можете взять все коробки первого и второго типа и одну коробку третьего типа.
Общее количество единиц будет = (1 * 3) + (2 * 2) + (1 * 1) = 8.
*/

func main() {
	boxTypes := [][]int{{1, 3}, {2, 2}, {3, 1}}
	fmt.Println(maximumUnits(boxTypes, 4)) // 8
}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	// Сортируем коробки по количеству единиц в убывающем порядке
	sort.Slice(boxTypes, func(i, j int) bool { // [[1 3] [2 2] [3 1]]
		return boxTypes[i][1] > boxTypes[j][1]
	})
	fmt.Println(boxTypes)

	countOfUnits := 0

	// Итерируемся по отсортированным коробкам
	for _, box := range boxTypes {
		// Определяем сколько коробок можно взять (не больше чем есть или вмещает грузовик)
		allowedBoxes := box[0]
		if allowedBoxes > truckSize {
			allowedBoxes = truckSize
		}

		// Уменьшаем оставшийся объем грузовика
		truckSize -= allowedBoxes
		// Увеличиваем общее количество единиц
		countOfUnits += allowedBoxes * box[1]

		// Если грузовик заполнен, выходим из цикла
		if truckSize == 0 {
			break
		}
	}

	return countOfUnits
}
