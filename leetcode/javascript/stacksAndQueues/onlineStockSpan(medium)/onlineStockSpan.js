/* https://leetcode.com/problems/online-stock-span/description/

Разработайте алгоритм, который собирает ежедневные котировки цен на некоторые акции и возвращает диапазон цен этих акций
на текущий день.

Размах цены акции за один день — это максимальное количество последовательных дней (начиная с этого дня и в обратном
направлении), в течение которых цена акции была меньше или равна цене этого дня.

Например, если цена акции за последние четыре дня равна [7,2,1,2], а цена акции сегодня равна 2, то интервал
сегодняшнего дня равен 4, потому что начиная с сегодняшнего дня цена акции равна 2. запас был меньше или равен 2 в
течение 4 дней подряд.
Кроме того, если цена акции за последние четыре дня равна [7,34,1,2], а цена акции сегодня равна 8, то интервал
сегодняшнего дня равен 3, потому что начиная с сегодняшнего дня цена акции равна было меньше или равно 8 в течение
3 дней подряд.
Реализуйте класс StockSpanner:

StockSpanner() Инициализирует объект класса.
int next(int цена) Возвращает интервал цены акции, учитывая, что сегодняшняя цена равна цене

Input
["StockSpanner", "next", "next", "next", "next", "next", "next", "next"]
[[], [100], [80], [60], [70], [60], [75], [85]]
Output
[null, 1, 1, 1, 2, 1, 4, 6]
*/

var StockSpanner = function() {
    this.stack = [];
};

/**
 * @param {number} price
 * @return {number}
 */
StockSpanner.prototype.next = function(price) {
    let span= 1;

    while (this.stack.length && this.stack[this.stack.length - 1][1] <= price) {
        const top = this.stack.pop();
        span += top[0];
    }

    this.stack.push([span, price]);
    return span;
};

/**
 * Your StockSpanner object will be instantiated and called as such:
 * var obj = new StockSpanner()
 * var param_1 = obj.next(price)
 */