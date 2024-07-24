/* https://leetcode.com/problems/task-scheduler/description/

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

/**
 * @param {string[]} tasks
 * @param {number} n
 * @return {number}
 */
var leastInterval = function(tasks, n) {
    const taskMap = new Map();

    for (const task of tasks) {
        if (taskMap.has(task)) taskMap.set(task, taskMap.get(task) + 1);
        else taskMap.set(task, 1);
    }

    const taskCounts = Array.from(taskMap.values()).sort((a, b) => b - a);
    const maxCount = taskCounts[0];
    let idleSlots = (maxCount - 1) * n;

    for (let i = 1; i < taskCounts.length; i++) {
        idleSlots -= Math.min(taskCounts[i], maxCount - 1);
    }

    idleSlots = Math.max(0, idleSlots);

    return tasks.length + idleSlots;
};

const tasks = ["A","A","A","B","B","B"];
const n = 2;
console.log(leastInterval(tasks, n)); // 8