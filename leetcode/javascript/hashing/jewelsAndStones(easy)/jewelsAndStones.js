/* https://leetcode.com/problems/jewels-and-stones/description/
Вам даны строки jewels, представляющие типы камней, которые являются jewels, и stones,
представляющие те камни, которые у вас есть. Каждый символ в stones — это тип камня, который у вас есть.
Вы хотите знать, сколько stones у вас тоже являются jewels.
Буквы чувствительны к регистру, поэтому «а» считается другим типом камня, чем «А».

Input: jewels = "aA", stones = "aAAbbbb"
Output: 3

Input: jewels = "z", stones = "ZZ"
Output: 0
*/

/**
 * @param {string} jewels
 * @param {string} stones
 * @return {number}
 */
var numJewelsInStones = function(jewels, stones) {
    let count = 0;

    for (let i = 0; i < stones.length; i++) {
        if (jewels.includes(stones[i])) count++;
    }

    return count;
};

// var numJewelsInStones = function(jewels, stones) {
//     const map = new Map();
//     let count = 0;
//
//     for (let i = 0; i < jewels.length; i++) {
//         if (map.has(jewels[i])) {
//             map.set(jewels[i], map.get(jewels[i]) + 1);
//         } else {
//             map.set(jewels[i], 1);
//         }
//     }
//
//     for (let i = 0; i < stones.length; i++) {
//         if (map.has(stones[i])) count++;
//     }
//
//     return count;
// };

console.log(numJewelsInStones("aA", "aAAbbbb")); // 3