package leetcode.java.design.MinStack;

import java.util.ArrayDeque;
import java.util.Deque;

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

public class MinStack {
    public static void main(String[] args) {
        MinStack minStack = new MinStack();
        minStack.push(-2);
        minStack.push(0);
        minStack.push(-3);

        System.out.println("Current min: " + minStack.getMin()); // -3
        minStack.pop();
        System.out.println("Top element: " + minStack.top());    // 0
        System.out.println("Current min: " + minStack.getMin()); // -2
    }

    private static class Element {
        int value;
        int min;

        Element(int value, int min) {
            this.value = value;
            this.min = min;
        }
    }

    private final Deque<Element> elements;

    public MinStack() {
        elements = new ArrayDeque<>();
    }
   
    // push - помещает элемент val в стек.
    // time: O(1), space: O(1)
    public void push(int val) {
        int minValue = val;
        if (!elements.isEmpty()) {
            int currentMin = getMin();
            if (currentMin < minValue) {
                minValue = currentMin;
            }
        }
        elements.push(new Element(val, minValue));
    }
    
    // pop - удаляет элемент на вершине стека.
    // time: O(1), space: O(1)
    public void pop() {
        if (!elements.isEmpty()) {
            elements.pop();
        }
    }
    
    // top - получает верхний элемент стека.
    // time: O(1), space: O(1)
    public int top() {
        if (elements.isEmpty()) {
            throw new RuntimeException("Stack is empty");
        }
        return elements.peek().value;
    }
    
    // getMin - получает минимальный элемент стека.
    // time: O(1), space: O(1)
    public int getMin() {
        if (elements.isEmpty()) {
            throw new RuntimeException("Stack is empty");
        }
        return elements.peek().min;
    }
}
