/* https://leetcode.com/problems/min-stack/description/

Спроектируйте стек, который поддерживает push, pop, top и получение минимального элемента за постоянное время.

Реализуйте класс MinStack:

MinStack() инициализирует объект стека.
void push(int val) помещает элемент val в стек.
void pop() удаляет элемент из вершины стека.
int top() получает верхний элемент стека.
int getMin() извлекает минимальный элемент в стеке.
Вы должны реализовать решение с временной сложностью O(1) для каждой функции.

Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2
*/

var MinStack = function() {
    this.elements = [];
};

/**
 * @param {number} val
 * @return {void}
 */
MinStack.prototype.push = function(val) {
    this.elements.push({
        value: val,
        min: this.elements.length === 0 ? val : Math.min(val, this.getMin()),
    });
};

/**
 * @return {void}
 */
MinStack.prototype.pop = function() {
    this.elements.pop();
};

/**
 * @return {number}
 */
MinStack.prototype.top = function() {
    return this.elements[this.elements.length - 1].value;
};

/**
 * @return {number}
 */
MinStack.prototype.getMin = function() {
    return this.elements[this.elements.length - 1].min;
};

/**
 * Your MinStack object will be instantiated and called as such:
 * var obj = new MinStack()
 * obj.push(val)
 * obj.pop()
 * var param_3 = obj.top()
 * var param_4 = obj.getMin()
 */