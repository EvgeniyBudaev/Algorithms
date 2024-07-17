/*
Учитывая массив временных интервалов встреч, где intervals[i] = [starti, endi], определите, может ли человек
присутствовать на всех собраниях.

Input: intervals = [[0,30],[5,10],[15,20]]
Output: false

Input: intervals = [[7,10],[2,4]]
Output: true
*/

var canAttendMeetings = function(intervals) {
    intervals.sort((a, b) => a[0] - b[0]); // [[2,4],[7,10]]
    for (let i = 1; i < intervals.length; i++) {
        if (intervals[i][0] < intervals[i - 1][1]) {
            return false;
        }
    }
    return true;
};

console.log(canAttendMeetings([[7,10],[2,4]])); // true