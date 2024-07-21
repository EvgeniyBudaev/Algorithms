/* https://leetcode.com/problems/house-robber-ii/description/

Вы профессиональный грабитель, планирующий грабить дома на улице. В каждом доме хранится определенная сумма денег.
Все дома в этом месте расположены по кругу. Это означает, что первый дом является соседом последнего.
При этом в соседних домах подключена система безопасности, и она автоматически обратится в полицию, если в одну ночь
были взломаны два соседних дома.
Учитывая целочисленный массив чисел, представляющий сумму денег в каждом доме, верните максимальную сумму денег,
которую вы можете ограбить сегодня вечером, не предупредив полицию.

Input: nums = [2,3,2]
Output: 3
Пояснение: Вы не можете ограбить дом 1 (деньги = 2), а затем ограбить дом 3 (деньги = 2), потому что это соседние дома

Input: nums = [1,2,3,1]
Output: 4
Пояснение: Ограбить дом 1 (деньги = 1), а затем ограбить дом 3 (деньги = 3).
Общая сумма, которую вы можете ограбить = 1 + 3 = 4.

Input: nums = [1,2,3]
Output: 3
*/

/**
 * @param {number[]} nums
 * @return {number}
 */
var rob = function(nums) {
    const len = nums.length;
    if (len < 4) return Math.max(...nums);

    function robberHelper(start, end) {
        let prev = 0, beforePrev = 0;
        for (let i = start; i < end; i++) {
            let tmp = prev;
            prev = Math.max(nums[i] + beforePrev, prev);
            beforePrev = tmp;
        }
        return prev;
    }
    return Math.max(robberHelper(0, len - 1), robberHelper(1, len));
};

console.log(rob([2,3,2])); // 3