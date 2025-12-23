package stacksAndQueues.MovingAverageFromDataStream;

/* 346. Moving Average from Data Stream

Учитывая поток целых чисел и размер окна, вычислите скользящее среднее всех целых чисел в скользящем окне.

* For example,
*
* MovingAverage m = new MovingAverage(3);
* m.next(1) = 1
* m.next(10) = (1 + 10) / 2
* m.next(3) = (1 + 10 + 3) / 3
* m.next(5) = (10 + 3 + 5) / 3

Input: ["MovingAverage", "next", "next", "next", "next"]
[[3], [1], [10], [3], [5]]
Output: [null, 1.0, 5.5, 4.66667, 6.0]
*/

public class MovingAverageFromDataStream {
    public static void main(String[] args) {
        MovingAverage movingAverage = new MovingAverage(3);

        System.out.println(movingAverage.next(1));  // 1.0
        System.out.println(movingAverage.next(10)); // 5.5
        System.out.println(movingAverage.next(3));  // 4.666666666666667
        System.out.println(movingAverage.next(5));  // 6.0
    }
}
