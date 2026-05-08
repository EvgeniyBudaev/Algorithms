package leetcode.java.intervals.MeetingRooms;

import java.util.Arrays;

/* 252. Meeting Rooms

Учитывая массив временных интервалов встреч, где intervals[i] = [starti, endi], определите, может ли человек
присутствовать на всех собраниях.

Input: intervals = [[0,30],[5,10],[15,20]]
Output: false

Input: intervals = [[7,10],[2,4]]
Output: true
*/

public class MeetingRooms {
    public static void main(String[] args) {
        int[][] intervals = {{7, 10}, {2, 4}};
        System.out.println(canAttendMeetings(intervals)); // true 
    }

    // canAttendMeetings определяет, может ли человек присутствовать на всех собраниях.
    // time: O(n log n) из-за сортировки, space: O(1) если не учитывать память сортировки
    public static boolean canAttendMeetings(int[][] intervals) {
        // Сортируем интервалы по начальному времени
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[0], b[0]));
        // [[2,4],[7,10]]

        // Проверяем, есть ли пересечения интервалов
        for (int i = 1; i < intervals.length; i++) {
            // Если конец предыдущего интервала больше начала текущего интервала, 
            // значит они пересекаются
            if (intervals[i][0] < intervals[i - 1][1]) {
                return false;
            }
        }

        return true; // Все интервалы не пересекаются, человек может присутствовать на всех собраниях
    }
}
