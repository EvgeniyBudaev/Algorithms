package stacksAndQueues.MinStack;

import java.util.ArrayDeque;
import java.util.Deque;

/* 155. Min Stack
https://leetcode.com/problems/min-stack/description/

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

public class MinStack {
    public static void main(String[] args) {
        MinStack obj = new MinStack();
        obj.push(-2);
        obj.push(0);
        obj.push(-3);
        System.out.println(obj.getMin()); // -3
        obj.pop();
        System.out.println(obj.top());    // 0
        System.out.println(obj.getMin()); // -2
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
    
    public void pop() {
        if (!elements.isEmpty()) {
            elements.pop();
        }
    }
    
    public int top() {
        if (elements.isEmpty()) {
            throw new RuntimeException("Stack is empty");
        }
        return elements.peek().value;
    }
    
    public int getMin() {
        if (elements.isEmpty()) {
            throw new RuntimeException("Stack is empty");
        }
        return elements.peek().min;
    }
}

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack obj = new MinStack();
 * obj.push(val);
 * obj.pop();
 * int param_3 = obj.top();
 * int param_4 = obj.getMin();
 */