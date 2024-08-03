/* https://leetcode.com/problems/insert-delete-getrandom-o1/description/

Реализуйте класс RandomizedSet:

RandomizedSet() Инициализирует объект RandomizedSet.
bool Insert(int val) Вставляет элемент val в набор, если он отсутствует. Возвращает true, если элемент отсутствует, в
противном случае — false.
bool Remove(int val) Удаляет элемент val из набора, если он присутствует. Возвращает true, если элемент присутствовал, и
false в противном случае.
int getRandom() Возвращает случайный элемент из текущего набора элементов (гарантируется, что при вызове этого метода
существует хотя бы один элемент). Каждый элемент должен иметь одинаковую вероятность возврата.
Вы должны реализовать функции класса так, чтобы каждая функция работала со средней временной сложностью O(1).

Input
["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
[[], [1], [2], [2], [], [1], [2], []]
Output: [null, true, false, true, 2, true, false, 2]

Объяснение
RandomizedSet randomizedSet = новый RandomizedSet();
randomizedSet.insert(1); // Вставляет 1 в набор. Возвращает true, поскольку 1 была вставлена успешно.
randomizedSet.remove(2); // Возвращает false, поскольку числа 2 не существует в наборе.
randomizedSet.insert(2); // Вставляет 2 в набор, возвращает true. Набор теперь содержит [1,2].
randomizedSet.getRandom(); // getRandom() должен возвращать либо 1, либо 2 случайным образом.
randomizedSet.remove(1); // Удаляет 1 из набора, возвращает true. Набор теперь содержит [2].
randomizedSet.insert(2); // 2 уже было в наборе, поэтому возвращаем false.
randomizedSet.getRandom(); // Поскольку 2 — единственное число в наборе, getRandom() всегда будет возвращать 2.
*/

var RandomizedSet = function() {
    this.container = new Set();
};

/**
 * @param {number} val
 * @return {boolean}
 */
RandomizedSet.prototype.insert = function(val) {
    if (this.container.has(val)) return false;
    else {
        this.container.add(val);
        return true;
    }
};

/**
 * @param {number} val
 * @return {boolean}
 */
RandomizedSet.prototype.remove = function(val) {
    return this.container.delete(val);
};

/**
 * @return {number}
 */
RandomizedSet.prototype.getRandom = function() {
    const index = Math.floor(Math.random() * this.container.size);
    const arr = [...this.container];
    return arr[index];
};

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * var obj = new RandomizedSet()
 * var param_1 = obj.insert(val)
 * var param_2 = obj.remove(val)
 * var param_3 = obj.getRandom()
 */

const randomizedSet = new RandomizedSet();
randomizedSet.insert(1);
randomizedSet.remove(2);
randomizedSet.insert(2);
randomizedSet.getRandom();
randomizedSet.remove(1);
randomizedSet.insert(2);
randomizedSet.getRandom();