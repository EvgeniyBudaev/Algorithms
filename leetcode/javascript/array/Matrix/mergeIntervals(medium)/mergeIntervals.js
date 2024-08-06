/* https://leetcode.com/problems/merge-intervals/description/

Учитывая массив интервалов, где интервалы[i] = [starti, endi], объедините все перекрывающиеся интервалы и верните массив
непересекающихся интервалов, охватывающих все интервалы во входных данных.

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Пояснение: Поскольку интервалы [1,3] и [2,6] перекрываются, объедините их в [1,6].

Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Пояснение: Интервалы [1,4] и [4,5] считаются перекрывающимися
*/

/**
 * @param {number[][]} intervals
 * @return {number[][]}
 */
var merge = function(intervals) {
    if(!intervals.length) return [];
    intervals.sort((a, b) => a[0] - b[0]); // [[1,3],[2,6],[8,10],[15,18]]
    const merged = [];
    let prev = intervals[0];

    for (let i = 1; i < intervals.length; i++) {
        let interval = intervals[i];
        if (interval[0] <= prev[1]) {
            prev[1] = Math.max(prev[1], interval[1]);
        } else {
            merged.push(prev);
            prev = interval;
        }
    }

    merged.push(prev);
    return merged;
};

const intervals = [[1,3],[2,6],[8,10],[15,18]];
console.log(merge(intervals)); // [[1,6],[8,10],[15,18]]