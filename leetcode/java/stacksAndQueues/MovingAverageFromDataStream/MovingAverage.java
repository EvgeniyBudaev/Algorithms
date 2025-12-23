package stacksAndQueues.MovingAverageFromDataStream;

import java.util.ArrayDeque;
import java.util.Deque;

public class MovingAverage {
    private final int size;
    private final Deque<Integer> queue;
    private double sum;

    public MovingAverage(int size) {
        this.size = size;
        this.queue = new ArrayDeque<>(size);
        this.sum = 0.0;
    }

    public double next(int val) {
        if (queue.size() >= size) {
            int removedVal = queue.pollFirst(); // Удаляем самый старый элемент
            sum -= removedVal;
        }
        queue.offerLast(val);
        sum += val;

        return sum / queue.size();
    }
}
