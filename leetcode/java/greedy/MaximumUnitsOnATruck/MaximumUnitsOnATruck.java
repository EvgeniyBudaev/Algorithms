package greedy.MaximumUnitsOnATruck;

import java.util.Arrays;

/* 1710. Maximum Units on a Truck
https://leetcode.com/problems/maximum-units-on-a-truck/description/

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

public class MaximumUnitsOnATruck {
    public static void main(String[] args) {
        int[][] boxTypes = {{1, 3}, {2, 2}, {3, 1}};
        int truckSize = 4;
        System.out.println(maximumUnits(boxTypes, truckSize)); // 8
    }

    // maximumUnits возвращает максимальное количество коробок, которое можно разместить в грузовике.
    // time: O(n), space: O(1)
    private static int maximumUnits(int[][] boxTypes, int truckSize) {
        // Сортируем коробки по количеству единиц в убывающем порядке
        Arrays.sort(boxTypes, (a, b) -> Integer.compare(b[1], a[1]));
        // После сортировки: [[1, 3], [2, 2], [3, 1]]

        int countOfUnits = 0; // Количество единиц, которое можно разместить в грузовике

        // Итерируемся по отсортированным коробкам
        for (int[] box : boxTypes) {
            int boxesAvailable = box[0];      // Количество коробок данного типа
            int unitsPerBox = box[1];         // Единиц товара в одной коробке

            // Берём столько коробок, сколько позволяет оставшееся место в грузовике
            int boxesToTake = Math.min(boxesAvailable, truckSize);
            countOfUnits += boxesToTake * unitsPerBox;
            truckSize -= boxesToTake;

            if (truckSize == 0) {
                break;
            }
        }

        return countOfUnits;
    }
}
