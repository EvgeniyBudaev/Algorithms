package leetcode.java.heaps.MinHeap;

import java.util.ArrayList;
import java.util.List;

/**
 * MinHeap - реализация мин-кучи (min-heap)
 * Корень кучи всегда содержит минимальный элемент
 */
public class MinHeap {
    private List<Integer> arr; // Список для хранения элементов кучи

    /**
     * Конструктор - создаёт пустую кучу
     */
    public MinHeap() {
        this.arr = new ArrayList<>();
    }

    /**
     * Push - добавить элемент в кучу.
     * time: O(log(n)), space: O(1)
     */
    public void push(int x) {
        arr.add(x);                        // Добавляем элемент в конец списка
        shiftUp(arr.size() - 1);           // Просеиваем элемент вверх
    }

    /**
     * popTop - удалить корневой элемент из кучи.
     * time: O(log(n)), space: O(1)
     */
    public void popTop() {
        if (arr.isEmpty()) {               // Если куча пуста, ничего не делаем
            return;
        }
        // Меняем местами корневой элемент и последний
        swap(0, arr.size() - 1);
        arr.remove(arr.size() - 1);        // Удаляем последний элемент
        shiftDown(0);                      // Просеиваем корневой элемент вниз
    }

    /**
     * top - вернуть корневой элемент из кучи.
     * time: O(1), space: O(1)
     */
    public int top() {
        if (arr.isEmpty()) {               // Если куча пуста, выбрасываем исключение
            throw new IllegalStateException("Heap is empty");
        }
        return arr.get(0);                 // Возвращаем корневой элемент
    }

    /**
     * empty - проверить, пуста ли куча.
     * time: O(1), space: O(1)
     */
    public boolean empty() {
        return arr.isEmpty();              // Если список пуст, то куча пуста
    }

    /**
     * shiftDown - просеивание вниз.
     * time: O(log(n)), space: O(1)
     */
    private void shiftDown(int i) {
        int leftChildIdx = i * 2 + 1;      // Индекс левого потомка
        int rightChildIdx = i * 2 + 2;     // Индекс правого потомка

        if (leftChildIdx >= arr.size()) {  // Если левого потомка нет, просеивание завершено
            return;
        }

        int nextIdx = leftChildIdx;        // Индекс следующего элемента для просеивания
        int nextMinVal = arr.get(leftChildIdx);

        // Если правый потомок существует и меньше левого, выбираем его
        if (rightChildIdx < arr.size() && arr.get(rightChildIdx) < nextMinVal) {
            nextIdx = rightChildIdx;
            nextMinVal = arr.get(rightChildIdx);
        }

        int currentVal = arr.get(i);

        // Если потомок меньше текущего, меняем местами и продолжаем просеивание
        if (nextMinVal < currentVal) {
            swap(i, nextIdx);
            shiftDown(nextIdx);
        }
    }

    /**
     * shiftUp - просеивание вверх.
     * time: O(log(n)), space: O(1)
     */
    private void shiftUp(int i) {
        if (i == 0) {                      // Если элемент - корень, просеивание завершено
            return;
        }
        int parentIdx = (i - 1) / 2;       // Индекс родителя

        // Если родитель больше текущего, меняем местами и продолжаем
        if (arr.get(parentIdx) > arr.get(i)) {
            swap(parentIdx, i);
            shiftUp(parentIdx);
        }
    }

    /**
     * Вспомогательный метод для обмена элементов местами
     */
    private void swap(int i, int j) {
        int temp = arr.get(i);
        arr.set(i, arr.get(j));
        arr.set(j, temp);
    }

    /**
     * Точка входа в программу
     */
    public static void main(String[] args) {
        MinHeap h = new MinHeap();
        h.push(3);
        h.push(1);
        h.push(4);
        h.push(1);
        h.push(5);

        System.out.println("Top: " + h.top());           // 1
        h.popTop();
        System.out.println("Top after pop: " + h.top()); // 1
        h.popTop();
        System.out.println("Top after pop: " + h.top()); // 3
    }
}
