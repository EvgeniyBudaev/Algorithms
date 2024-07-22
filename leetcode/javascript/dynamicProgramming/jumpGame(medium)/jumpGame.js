/* https://leetcode.com/problems/jump-game/description/

Вам дан целочисленный массив чисел. Изначально вы находитесь в первом индексе массива, и каждый элемент массива
представляет максимальную длину прыжка в этой позиции.
Возвращайте true, если вы можете достичь последнего индекса, или false в противном случае.

Input: nums = [2,3,1,1,4]
Output: true
Пояснение: Перейти на 1 шаг от индекса 0 к 1, затем на 3 шага до последнего индекса.

Input: nums = [3,2,1,0,4]
Output: false
Пояснение: Несмотря ни на что, вы всегда будете достигать индекса 3. Его максимальная длина перехода равна 0, что делает
невозможным достижение последнего индекса
*/

/**
 * @param {number[]} nums
 * @return {boolean}
 */
var canJump = function(nums) {
    if (nums.length <= 1) return true;
    let maximum = nums[0];

    for (let i = 0; i < nums.length; i++) {
        if (maximum <= i && nums[i] === 0) return false;
        if (i + nums[i] > maximum) maximum = i + nums[i];
        if (maximum >= nums.length - 1) return true;
    }

    return false;
};

console.log(canJump([2,3,1,1,4])); // true