package greedy.HowManyApplesCanYouPutIntoTheBasket;

import java.util.Arrays;

/* 1196. How Many Apples Can You Put into the Basket

У вас есть несколько яблок, где arr[i] — вес i-го яблока.
У вас также есть корзина, способная выдержать до 5000 единиц веса.
Верните максимальное количество яблок, которое вы можете положить в корзину.

Input: arr = [100,200,150,1000]
Output: 4
Пояснение: Все 4 яблока можно нести в корзине, так как их сумма весов равна 1450.

Input: arr = [900,950,800,1000,700,800]
Output: 5
Пояснение: Сумма весов 6 яблок превышает 5000, поэтому мы выбираем любые 5 из них.
*/

public class HowManyApplesCanYouPutIntoTheBasket {
    public static void main(String[] args) {
        int[] arr = {900, 950, 800, 1000, 700, 800};
        System.out.println(maxNumberOfApples(arr)); // 5
    }

    // maxNumberOfApples возвращает максимальное количество яблок, которое можно поместить в корзину без превышения веса.
    // time: O(n), space: O(1)
    private static int maxNumberOfApples(int[] appleWeights) {
        final int LIMIT = 5000; // Максимальный вес корзины
        int totalWeight = 0;    // Общий вес яблок

        Arrays.sort(appleWeights); // Сортируем яблоки по весу (от легких к тяжелым)

        for (int i = 0; i < appleWeights.length; i++) { // Перебираем все яблоки
            totalWeight += appleWeights[i]; // Добавляем вес текущего яблока к общему весу
            if (totalWeight > LIMIT) { // Если превысили лимит в 5000, возвращаем текущее количество
                return i;
            }
        }

        return appleWeights.length; // Если все яблоки поместились, возвращаем общее количество
    }
}
