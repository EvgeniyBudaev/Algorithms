/*
У вас есть несколько яблок, где arr[i] — вес i-го яблока. У вас также есть корзина, способная выдержать до 5000 единиц
веса.
Верните максимальное количество яблок, которое вы можете положить в корзину.
*/

/*
Input: arr = [100,200,150,1000]
Output: 4
Пояснение: Все 4 яблока можно нести в корзине, так как их сумма весов равна 1450.
 */

/*
Input: arr = [900,950,800,1000,700,800]
Output: 5
Пояснение: Сумма весов 6 яблок превышает 5000, поэтому мы выбираем любые 5 из них.
 */

/**
 * Finds the maximum number of apples that can fit within a weight limit.
 * Assumes each apple's weight is given in an array, with a weight limit of 5000.
 * @param appleWeights Array of individual apple weights.
 * @returns The maximum number of apples that can be included without exceeding the weight limit.
 */
function maxNumberOfApples(appleWeights) {
    const appleWeightsSorted = appleWeights.sort((a, b) => a - b);
    let totalWeight = 0;

    for (let i = 0; i < appleWeightsSorted.length; i++) {
        // Добавляем вес текущего яблока к общей сумме веса.
        totalWeight += appleWeightsSorted[i];

        // Если лимит веса превышен для текущего яблока, возвращаем количество яблок
        if (totalWeight > 5000) {
            return i;
        }
    }

    // Если общий вес никогда не превышает лимит, возвращаем общее количество яблок.
    return appleWeightsSorted.length;
}

console.log(maxNumberOfApples([900,950,800,1000,700,800])); // 5