package main

import (
	"fmt"
	"sort"
)

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

func main() {
	trips := [][]int{{2, 1, 5}, {3, 3, 7}}
	fmt.Println(carPooling(trips, 4)) // false
	fmt.Println(carPooling(trips, 5)) // true
}

// carPooling возвращает true, если возможно посадить и высадить всех пассажиров для всех указанных поездок,
// или false в противном случае
// time complexity: O(N*log(N)), где N — количество поездок.
// space complexity: O(N), где N — количество поездок.
func carPooling(trips [][]int, capacity int) bool {
	// point — это структура, которая хранит координату и изменение пассажиров в данной точке.
	type point struct {
		coord int // координата
		delta int // изменение пассажиров
	}
	// points — это срез структур point, который содержит все координаты и изменения пассажиров.
	points := make([]point, 0, len(trips)*2)

	for _, trip := range trips {
		// добавляем изменения пассажиров в точках начала и конца поездок в срез
		points = append(points, point{coord: trip[1], delta: trip[0]})  // зашли
		points = append(points, point{coord: trip[2], delta: -trip[0]}) // вышли
	}
	// points [{1, 2} {5, -2} {3, 3} {7, -3}]

	// сортируем точки по координатам и изменениям пассажиров
	sort.Slice(points, func(i, j int) bool {
		// сначала сортируем по координатам, затем по изменениям пассажиров
		if points[i].coord == points[j].coord {
			return points[i].delta < points[j].delta
		}
		// если координаты совпадают, сортируем по изменениям пассажиров
		return points[i].coord < points[j].coord
	})
	// points [{1, 2} {3, 3} {5, -2} {7, -3}]

	currentPassengers := 0 // текущее количество пассажиров
	for _, p := range points {
		currentPassengers += p.delta // обновляем текущее количество пассажиров
		// если текущее количество пассажиров превышает вместимость, возвращаем false
		if currentPassengers > capacity {
			return false
		}
	}

	// если мы прошли все поездки и ни разу не превысили вместимость, возвращаем true
	return true
}
