package main

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

func main() {
	obj := Constructor()
	println(obj.Next(100)) // 1
	println(obj.Next(80))  // 1
	println(obj.Next(60))  // 1
	println(obj.Next(70))  // 2
	println(obj.Next(60))  // 1
	println(obj.Next(75))  // 4
	println(obj.Next(85))  // 6
}

type StockSpanner struct {
	stack [][2]int // Каждый элемент стека - массив из двух чисел [span, price]
}

func Constructor() StockSpanner {
	return StockSpanner{
		stack: make([][2]int, 0),
	}
}

func (this *StockSpanner) Next(price int) int {
	span := 1

	// Удаляем из стека все элементы с ценой <= текущей, суммируя их span
	for len(this.stack) > 0 && this.stack[len(this.stack)-1][1] <= price {
		top := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1] // Удаляем верхний элемент
		span += top[0]                              // Увеличиваем span на значение из удаленного элемента
	}

	// Добавляем текущий элемент в стек
	this.stack = append(this.stack, [2]int{span, price})
	return span
}
