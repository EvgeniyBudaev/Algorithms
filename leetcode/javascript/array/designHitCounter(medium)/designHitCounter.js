/* https://www.lintcode.com/problem/3662/description

В этом вопросе вам нужно создать класс HitCounter.
В этом классе имеются следующие функции:

HitCounter(): конструктор без аргументов
void hit(int timestamp): указывает, что касание происходит в указанное время.
int getHits(int timestamp): возвращает общее количество обращений за 300 секунд до указанного времени.
Где временная метка находится в секундах.

Input:
["HitCounter", "hit", "hit", "hit", "getHits", "hit", "getHits", "getHits"]
[[], [1], [2], [3], [4], [300], [300], [301]]
Output: [null, null, null, null, 3, null, 4, 3]

Explanation
HitCounter hitCounter = new HitCounter();
hitCounter.hit(1);       // hit at timestamp 1.
hitCounter.hit(2);       // hit at timestamp 2.
hitCounter.hit(3);       // hit at timestamp 3.
hitCounter.getHits(4);   // get hits at timestamp 4, return 3.
hitCounter.hit(300);     // hit at timestamp 300.
hitCounter.getHits(300); // get hits at timestamp 300, return 4.
hitCounter.getHits(301); // get hits at timestamp 301, return 3.
*/

class HitCounter {
  ts = [];

  constructor() {}

  hit(timestamp) {
    this.ts.push(timestamp);
  }

  getHits(timestamp) {
    const search = (x) => {
      let [l, r] = [0, this.ts.length];
      while (l < r) {
        const mid = (l + r) >> 1;
        if (this.ts[mid] >= x) {
          r = mid;
        } else {
          l = mid + 1;
        }
      }
      return l;
    };
    return this.ts.length - search(timestamp - 300 + 1);
  }
}

/**
 * Your HitCounter object will be instantiated and called as such:
 * var obj = new HitCounter()
 * obj.hit(timestamp)
 * var param_2 = obj.getHits(timestamp)
 */

const hitCounter = new HitCounter();
hitCounter.hit(1);       // hit at timestamp 1.
hitCounter.hit(2);       // hit at timestamp 2.
hitCounter.hit(3);       // hit at timestamp 3.
hitCounter.getHits(4);   // get hits at timestamp 4, return 3.
hitCounter.hit(300);     // hit at timestamp 300.
hitCounter.getHits(300); // get hits at timestamp 300, return 4.
hitCounter.getHits(301); // get hits at timestamp 301, return 3.
