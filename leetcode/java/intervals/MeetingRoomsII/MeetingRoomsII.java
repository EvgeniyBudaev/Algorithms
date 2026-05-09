package leetcode.java.intervals.MeetingRoomsII;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

/* Meeting Rooms II

Дан двумерный целочисленный массив A размером N x 2, обозначающий временные интервалы различных встреч.

Где:
A[i][0] = время начала i-й встречи.
A[i][1] = время окончания i-й встречи.

Найдите минимальное количество конференц-залов, необходимое для проведения всех встреч.

Примечание: если встреча заканчивается в момент времени t, другая встреча, начинающаяся в момент времени t,
может использовать тот же конференц-зал.

Example 1:
Input: intervals = [(0,30),(5,10),(15,20)]
Output: 2
Explanation:
room1: (0,30)
room2: (5,10),(15,20)
*/

public class MeetingRoomsII {
    public static void main(String[] args) {
        int[][] intervals = {{0, 30}, {5, 10}, {15, 20}};
        System.out.println(minMeetingRooms(intervals)); // 2
    }

    // Point - точка начала или конца встречи
    private static class Point {
        int time;  // время начала или конца встречи
        int delta; // дельта количества комнат: +1 если начало встречи, -1 если конец встречи

        Point(int time, int delta) {
            this.time = time;
            this.delta = delta;
        }
    }

    // minMeetingRooms - находит минимальное количество конференц-залов, необходимое для проведения всех встреч.
    // time: O(n * log n), space: O(n)
    public static int minMeetingRooms(int[][] intervals) {
        List<Point> points = new ArrayList<>(intervals.length * 2);

        // Создаем точки начала и конца событий
        for (int[] interval : intervals) {
            points.add(new Point(interval[0], 1));  // начало встречи (+1 комната)
            points.add(new Point(interval[1], -1));  // конец встречи (-1 комната)
        }
        // points [{0, 1} {30, -1} {5, 1} {10, -1} {15, 1} {20, -1}]

        // Сортируем точки:
        // 1. По времени
        // 2. Если время одинаковое, сначала идут окончания встреч (delta = -1)
        Collections.sort(points, (a, b) -> {
            if (a.time == b.time) {
                // если время совпадает, сначала обрабатываем окончания (delta = -1 < +1)
                return Integer.compare(a.delta, b.delta);
            }
            return Integer.compare(a.time, b.time);
        });
        // points [{0, 1} {5, 1} {10, -1} {15, 1} {20, -1} {30, -1}]

        int maxRooms = 0;      // максимальное количество комнат
        int currentRooms = 0;  // текущее количество комнат

        for (Point p : points) {
            currentRooms += p.delta; // обновляем текущее количество комнат
            // если количество комнат больше максимального, обновляем максимум
            if (currentRooms > maxRooms) {
                maxRooms = currentRooms;
            }
        }

        // возвращаем максимальное количество комнат
        return maxRooms;
    }
}
