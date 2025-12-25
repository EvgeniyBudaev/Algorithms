package stacksAndQueues.OnlineStockSpan;

import java.util.ArrayDeque;
import java.util.Deque;

public class StockSpanner {
    // Используем стек пар [span, price]
    private Deque<int[]> stack;

    public StockSpanner() {
        stack = new ArrayDeque<>();
    }
    
    public int next(int price) {
        int span = 1;

        // Удаляем из стека все элементы с ценой <= текущей, суммируя их span
        while (!stack.isEmpty() && stack.peek()[1] <= price) {
            span += stack.pop()[0];
        }

        // Добавляем текущую пару [span, price] в стек
        stack.push(new int[]{span, price});

        return span;
    }
}
