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

/**
 * @param {number} size
 */
var MovingAverage = function(size) {
    this.n = size;
    this.queue = [];
    this.average = 0.0;
};
/**
 * @param {number} val
 * @return {number}
 */
MovingAverage.prototype.next = function(val) {
    let removedVal;

    if (this.queue.length >= this.n) {
        removedVal = this.queue.shift();
        this.average = this.average - removedVal;
    }
    this.queue.push(val);
    this.average += val;

    return this.average / this.queue.length;
};

const movingAverage = new MovingAverage(3);

movingAverage.next(1); // return 1.0 = 1 / 1
movingAverage.next(10); // return 5.5 = (1 + 10) / 2
movingAverage.next(3); // return 4.66667 = (1 + 10 + 3) / 3
movingAverage.next(5); // return 6.0 = (10 + 3 + 5) / 3