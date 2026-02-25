package greedy.TaskScheduler;

import java.util.ArrayList;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/* 621. Task Scheduler
https://leetcode.com/problems/task-scheduler/description/

Вам дан массив задач ЦП, каждая из которых обозначена буквами от A до Z, и время охлаждения n. Каждый цикл или интервал
позволяет выполнить одну задачу. Задачи можно выполнять в любом порядке, но есть ограничение: одинаковые задачи должны
быть разделены как минимум n интервалами из-за времени остывания.
Верните минимальное количество интервалов, необходимое для выполнения всех задач.

Input: tasks = ["A","A","A","B","B","B"], n = 2
Output: 8

Пояснение: Возможная последовательность: A -> B -> ожидание -> A -> B -> ожидание -> A -> B.
После завершения задачи А вы должны подождать два цикла, прежде чем снова выполнить А. То же самое относится и к
задаче Б. В 3-м интервале ни А, ни Б сделать невозможно, поэтому вы бездельничаете. К 4-му циклу можно снова делать А,
так как прошло 2 интервала.
*/

public class TaskScheduler {
    public static void main(String[] args) {
        char[] tasks = {'A', 'A', 'A', 'B', 'B', 'B'};
        System.out.println(leastInterval(tasks, 2)); // 8 
    }

    // leastInterval вычисляет минимальное количество интервалов, необходимых для выполнения всех задач.
    // time: O(n), space: O(n)
    private static int leastInterval(char[] tasks, int n) {
        Map<Character, Integer> freqMap = new HashMap<>(); // Создаем карту для подсчета частот задач
        for (char task : tasks) {
            freqMap.put(task, freqMap.getOrDefault(task, 0) + 1);
        }

        // Создаем список для хранения частот и сортируем по убыванию
        List<Integer> taskCounts = new ArrayList<>(freqMap.values());
        taskCounts.sort(Collections.reverseOrder());

        int maxCount = taskCounts.get(0); // Вычисляем максимальное количество повторений
        int idleSlots = (maxCount - 1) * n;     // Вычисляем количество "холостых" слотов

        // Заполняем холостые слоты другими задачами
        for (int i = 1; i < taskCounts.size(); i++) {
            // Каждая другая задача может заполнить максимум (maxCount - 1) слотов
            idleSlots -= Math.min(taskCounts.get(i), maxCount - 1);
        }

        // Холостые слоты не могут быть отрицательными
        idleSlots = Math.max(0, idleSlots);
        
        return tasks.length + idleSlots; // Общее время = все задачи + холостые слоты
    }
}
