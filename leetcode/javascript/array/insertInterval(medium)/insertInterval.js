/* https://leetcode.com/problems/insert-interval/description/

Вам дан массив непересекающихся интервалов, где интервалы[i] = [starti, endi] представляют начало и конец i-го
интервала, а интервалы сортируются в порядке возрастания по starti. Вам также дан интервал newInterval = [start, end],
который представляет начало и конец другого интервала.

Вставьте newInterval в интервалы таким образом, чтобы интервалы по-прежнему сортировались в порядке возрастания по
начальным значениям, а интервалы по-прежнему не имели перекрывающихся интервалов (при необходимости объединяйте
перекрывающиеся интервалы).
Интервалы возврата после вставки.
Обратите внимание, что вам не нужно изменять интервалы на месте. Вы можете создать новый массив и вернуть его.

Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]

Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Пояснение: Потому что новый интервал [4,8] перекрывается с [3,5],[6,7],[8,10].
*/

/**
 * @param {number[][]} intervals
 * @param {number[]} newInterval
 * @return {number[][]}
 */
var insert = function(intervals, newInterval) {
    let [start, end] = newInterval;
    let left = [];
    let right = [];
    for (const interval of intervals) {
        const [first, last] = interval;
        if (last < start) left.push(interval);
        else if (first > end) right.push(interval);
        else {
            start = Math.min(start, first);
            end = Math.max(end, last);
        }
    }
    return [...left, [start, end], ...right];
};

const intervals = [[1,3],[6,9]];
const newInterval = [2,5];
console.log(insert(intervals, newInterval)); // [[1,5],[6,9]]