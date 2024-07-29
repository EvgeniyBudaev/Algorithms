/* https://www.lintcode.com/problem/3662/description

В этом вопросе вам нужно создать класс HitCounter.

В этом классе имеются следующие функции:

HitCounter(): конструктор без аргументов
void hit(int timestamp): указывает, что касание происходит в указанное время.
int getHits(int timestamp): возвращает общее количество обращений за 300 секунд до указанного времени.
Где временная метка находится в секундах.

Input
["HitCounter", "hit", "hit", "hit", "getHits", "hit", "getHits", "getHits"]
[[], [1], [2], [3], [4], [300], [300], [301]]
Output
[null, null, null, null, 3, null, 4, 3]

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

let v = [];

/* Record a hit.
   @param timestamp - The current timestamp (in
                      seconds granularity).  */

function hit(timestamp)
{
  v.push(timestamp);
}

// Time Complexity : O(1)

/** Return the number of hits in the past 5 minutes.
 @param timestamp - The current timestamp (in
  seconds granularity). */
function getHits(timestamp)
{
  let i, j;
  for (i = 0; i < v.length; ++i) {
    if (v[i] > timestamp - 300) {
      break;
    }
  }
  return v.length - i;
}

// Time Complexity : O(n)
