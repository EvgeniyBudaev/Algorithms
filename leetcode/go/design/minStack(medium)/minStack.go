package main

import "fmt"

/* 155. Min Stack
https://leetcode.com/problems/min-stack/description/

Разработайте стек, который поддерживает push, pop, top и извлечение минимального элемента за постоянное время.

Реализуйте класс MinStack:

MinStack() инициализирует объект стека.
void push(int val) помещает элемент val в стек.
void pop() удаляет элемент на вершине стека.
int top() получает верхний элемент стека.
int getMin() извлекает минимальный элемент в стеке.
Вам необходимо реализовать решение со сложностью времени O(1) для каждой функции.

Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]
*/

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	fmt.Println("Current min:", minStack.GetMin()) // -3
	minStack.Pop()
	fmt.Println("Top element:", minStack.Top())    // 0
	fmt.Println("Current min:", minStack.GetMin()) // -2
}

type MinStack struct {
	stack    []int // хранит значения стека
	minStack []int // хранит минимальные значения стека
}

// Constructor - конструктор стека
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

// Push - помещает элемент val в стек.
// time: O(1), space: O(1)
func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val) // добавляем значение в стек
	// если стек пустой или значение меньше минимального значения, то добавляем его в стек минимальных значений
	if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	}
}

// Pop - удаляет элемент на вершине стека.
// time: O(1), space: O(1)
func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}

	s := this.stack[len(this.stack)-1]          // сохраняем значение на вершине стека
	this.stack = this.stack[:len(this.stack)-1] // удаляем значение на вершине стека

	// если значение на вершине стека равно минимальному значению, то удаляем его из стека
	if len(this.minStack) > 0 && s == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

// Top - получает верхний элемент стека.
// time: O(1), space: O(1)
func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

// GetMin - получает минимальный элемент стека.
// time: O(1), space: O(1)
func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
