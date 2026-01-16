package binarySearch.KokoEatingBananas;

/* 875. Koko Eating Bananas
https://leetcode.com/problems/koko-eating-bananas/description/

Коко любит есть бананы. Есть n стопок бананов, в i-й стопке бананов находится [i]. Охранники ушли и вернутся через час.
Коко может решить, что ее скорость поедания бананов в час равна k. Каждый час она выбирает какую-то стопку бананов и
съедает k бананов из этой стопки. Если в куче меньше k бананов, она съедает их все и больше не будет есть бананов в
течение этого часа.
Коко любит есть медленно, но все равно хочет съесть все бананы до того, как вернутся охранники.
Найдите минимальное целое число k, при котором она сможет съесть все бананы за h часов.

Input: piles = [3,6,7,11], h = 8
Output: 4
*/

public class KokoEatingBananas {
    public static void main(String[] args) {
        int[] piles = {3, 6, 7, 11};
        System.out.println(minEatingSpeed(piles, 8)); // 4 
    }

    // minEatingSpeed - находит минимальную скорость, при которой Коко сможет съесть все бананы за заданное количество часов.
    private static int minEatingSpeed(int[] piles, int h) {
        int left = 1;                     // Минимальная скорость
        int right = maxInArray(piles);    // Максимальная скорость

        while (left < right) {
            int mid = left + (right - left) / 2;
            int hourSpent = getHourSpent(mid, piles); // getHourSpent - получить потраченные часы
            // Если потраченное время больше, чем заданное количество часов, увеличиваем минимальную скорость.
            if (h < hourSpent) {
                left = mid + 1;
            } else {
                right = mid;
            }
        }

        //  Возвращаем минимальную скорость, при которой Коко сможет съесть все бананы за заданное количество часов.
        return right;
    }

    // getHourSpent - метод для вычисления необходимого времени
     private static int getHourSpent(int speed, int[] piles) {
        int totalHours = 0; // Инициализируем переменную для хранения общего количества часов
        for (int pile : piles) {
            // Вычисляем количество часов для текущей стопки и добавляем к общему количеству часов
            totalHours += (pile + speed - 1) / speed;
        }

        return totalHours;
    }

    // maxInArray - метод для нахождения максимального значения в срезе
    private static int maxInArray(int[] nums) {
        int maxNum = nums[0]; // Инициализируем переменную для хранения максимального числа
        for (int num : nums) {
            if (num > maxNum) {
                maxNum = num;
            }
        }
        
        return maxNum;
    }
}
