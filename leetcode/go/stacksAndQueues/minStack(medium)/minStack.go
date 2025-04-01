package main

import "math"

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

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	println(obj.GetMin()) // -3
	obj.Pop()
	println(obj.Top())    // 0
	println(obj.GetMin()) // -2
}

type Element struct {
	value int
	min   int
}

type MinStack struct {
	elements []Element
}

func Constructor() MinStack {
	return MinStack{
		elements: make([]Element, 0),
	}
}

func (this *MinStack) Push(val int) {
	minValue := val
	if len(this.elements) > 0 {
		currentMin := this.GetMin()
		if currentMin < minValue {
			minValue = currentMin
		}
	}
	this.elements = append(this.elements, Element{value: val, min: minValue})
}

func (this *MinStack) Pop() {
	if len(this.elements) == 0 {
		return
	}
	this.elements = this.elements[:len(this.elements)-1]
}

func (this *MinStack) Top() int {
	if len(this.elements) == 0 {
		return math.MinInt32
	}
	return this.elements[len(this.elements)-1].value
}

func (this *MinStack) GetMin() int {
	if len(this.elements) == 0 {
		return math.MinInt32
	}
	return this.elements[len(this.elements)-1].min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
