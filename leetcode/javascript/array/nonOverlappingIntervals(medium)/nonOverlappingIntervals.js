/* https://leetcode.com/problems/non-overlapping-intervals/description/

Учитывая массив интервалов интервалов, где интервалы[i] = [starti, endi], верните минимальное количество интервалов,
которое вам нужно удалить, чтобы остальные интервалы не перекрывались.

Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
Output: 1
Пояснение: [1,3] можно удалить, а остальные интервалы не перекрываться.

Input: intervals = [[1,2],[1,2],[1,2]]
Output: 2
Пояснение: Вам нужно удалить два [1,2], чтобы остальные интервалы не перекрывались.

Input: intervals = [[1,2],[2,3]]
Output: 0
Объяснение: Вам не нужно удалять какие-либо интервалы, поскольку они уже не перекрываются.
*/

/**
 * @param {number[][]} intervals
 * @return {number}
 */
var eraseOverlapIntervals = function(intervals) {
    let res = 0;
    intervals.sort((a, b) => a[1] - b[1]);
    let prev_end = intervals[0][1];

    for (let i = 1; i < intervals.length; i++) {
        if (prev_end > intervals[i][0]) res++;
        else prev_end = intervals[i][1];
    }

    return res;
};

const intervals = [[1,2],[2,3],[3,4],[1,3]];
console.log(eraseOverlapIntervals(intervals)); // 1