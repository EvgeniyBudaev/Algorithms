/* https://leetcode.com/problems/daily-temperatures/description/

Учитывая, что массив целых чисел представляет собой ежедневную температуру, верните ответ массива, такой,
что ответ [i] — это количество дней, которое вам нужно подождать после i-го дня, чтобы получить более высокую
температуру. Если нет будущего дня, для которого это возможно, вместо этого сохраните ответ [i] == 0.

Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]

Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]

Input: temperatures = [30,60,90]
Output: [1,1,0]
*/

/**
 * @param {number[]} temperatures
 * @return {number[]}
 */
var dailyTemperatures = function(temperatures) {
    const stack = [];
    const result = new Array(temperatures.length).fill(0);

    for (let i = 0; i < temperatures.length; i++) {
        while (stack.length > 0 && temperatures[i] > temperatures[stack[stack.length - 1]]) {
            const idx = stack.pop(); // 0 -> 1 -> 4 -> 3 -> 5 -> 2
            result[idx] = i - idx;
        }
        stack.push(i);
    }

    return result;
};

console.log(dailyTemperatures([73,74,75,71,69,72,76,73])); // [1,1,4,2,1,1,0,0]