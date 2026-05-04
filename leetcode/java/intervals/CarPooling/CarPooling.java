package leetcode.java.intervals.CarPooling;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

/* 1094. Car Pooling
https://leetcode.com/problems/car-pooling/description/

Есть автомобиль с вместимостью пустых мест. Автомобиль едет только на восток
(т. е. он не может развернуться и ехать на запад).

Вам дана целая вместимость и массив trips, где trips[i] = [numPassengersi, fromi, toi] указывает,
что в i-й поездке находится numPassengersi пассажиров, а места посадки и высадки — fromi и toi соответственно.
Места указаны как количество километров к востоку от начального местоположения автомобиля.

Верните true, если возможно посадить и высадить всех пассажиров для всех указанных поездок, или false в противном случае.

Example 1:
Input: trips = [[2,1,5],[3,3,7]], capacity = 4
Output: false

Example 2:
Input: trips = [[2,1,5],[3,3,7]], capacity = 5
Output: true
*/

public class CarPooling {
    // Point — это класс, который хранит координату и изменение пассажиров в данной точке.
    private static class Point {
        int coord; // координата
        int delta; // изменение пассажиров

        Point(int coord, int delta) {
            this.coord = coord;
            this.delta = delta;
        }
    }

    public static void main(String[] args) {
        int[][] trips = {{2, 1, 5}, {3, 3, 7}};
        System.out.println(carPooling(trips, 4)); // false
        System.out.println(carPooling(trips, 5)); // true
    }

    // carPooling возвращает true, если возможно посадить и высадить всех пассажиров для всех указанных поездок,
    // или false в противном случае
    // time: O(N*log(N)), где N — количество поездок.
    // space: O(N), где N — количество поездок.
    private static boolean carPooling(int[][] trips, int capacity) {
        // points — это список структур Point, который содержит все координаты и изменения пассажиров.
        List<Point> points = new ArrayList<>(trips.length * 2);

        for (int[] trip : trips) {
            // добавляем изменения пассажиров в точках начала и конца поездок в список
            points.add(new Point(trip[1], trip[0]));   // зашли (+пассажиры)
            points.add(new Point(trip[2], -trip[0]));   // вышли (-пассажиры)
        }
        // points [{1, 2} {5, -2} {3, 3} {7, -3}]

        // сортируем точки по координатам и изменениям пассажиров
        Collections.sort(points, (a, b) -> {
            // сначала сортируем по координатам, затем по изменениям пассажиров
            if (a.coord == b.coord) {
                // при совпадении координат: сначала высадка (отрицательные delta), потом посадка
                return Integer.compare(a.delta, b.delta);
            }
            return Integer.compare(a.coord, b.coord);
        });
        // points [{1, 2} {3, 3} {5, -2} {7, -3}]

        int currentPassengers = 0; // текущее количество пассажиров
        for (Point p : points) {
            currentPassengers += p.delta; // обновляем текущее количество пассажиров
            // если текущее количество пассажиров превышает вместимость, возвращаем false
            if (currentPassengers > capacity) {
                return false;
            }
        }

        // если мы прошли все поездки и ни разу не превысили вместимость, возвращаем true
        return true;
    }
}
