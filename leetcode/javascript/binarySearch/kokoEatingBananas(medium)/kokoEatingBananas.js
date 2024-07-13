/* https://leetcode.com/problems/koko-eating-bananas/description/

Коко любит есть бананы. Есть n стопок бананов, в i-й стопке бананов находится [i]. Охранники ушли и вернутся через час.
Коко может решить, что ее скорость поедания бананов в час равна k. Каждый час она выбирает какую-то стопку бананов и
съедает k бананов из этой стопки. Если в куче меньше k бананов, она съедает их все и больше не будет есть бананов в
течение этого часа.
Коко любит есть медленно, но все равно хочет съесть все бананы до того, как вернутся охранники.
Найдите минимальное целое число k, при котором она сможет съесть все бананы за h часов.

Input: piles = [3,6,7,11], h = 8
Output: 4
*/

/**
 * @param {number[]} piles
 * @param {number} h
 * @return {number}
 */
var minEatingSpeed = function(piles, h) {
    let left = 1, right = Math.max(...piles);

    while (left < right) {
        const mid = Math.floor((left + right) / 2); // 6 -> 3 -> 5 -> 4
        const hourSpent = getHourSpent(mid, piles); // 6 -> 10 -> 8 -> 8
        if (h < hourSpent) left = mid + 1;
        else right = mid;
    }

    return right;
};


// getHourSpent - получить потраченные часы
const getHourSpent = (mid, piles, hourSpent = 0) => {
    for (const pile of piles) {
        hourSpent += Math.ceil(pile / mid);
    }
    return hourSpent;
};

console.log(minEatingSpeed([3,6,7,11], 8)); // 4