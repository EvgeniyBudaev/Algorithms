/* 986. Interval List Intersections
https://leetcode.com/problems/interval-list-intersections/description/

Вам даны два списка закрытых интервалов, firstList и SecondList,
где firstList[i] = [starti, endi] и SecondList[j] = [startj, endj]. Каждый список интервалов попарно непересекающийся и
отсортирован.
Верните пересечение этих двух списков интервалов.
Замкнутый интервал [a, b] (с a <= b) обозначает набор действительных чисел x с a <= x <= b.
Пересечение двух закрытых интервалов представляет собой набор действительных чисел, которые либо пусты, либо
представлены в виде замкнутого интервала. Например, пересечение [1, 3] и [2, 4] — это [2, 3].

Input: firstList = [[0,2],[5,10],[13,23],[24,25]], secondList = [[1,5],[8,12],[15,24],[25,26]]
Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
*/

/**
 * @param {number[][]} firstList
 * @param {number[][]} secondList
 * @return {number[][]}
 */
var intervalIntersection = function (firstList, secondList) {
    const result = []; // Результирующий список интервалов
    let left = 0;
    let right = 0;

    while (left < firstList.length && right < secondList.length) { // Пока не закончатся интервалы в обоих списках
        const first = firstList[left];          // Текущий интервал первого списка
        const second = secondList[right];       // Текущий интервал второго списка
        const start = Math.max(first[0], second[0]); // Максимальное начало
        const end = Math.min(first[1], second[1]);   // Минимальное окончание

        if (start <= end) {                 // Если начало меньше окончания, то добавляем интервал в результат
            result.push([start, end]);
        }

        // Если окончание первого интервала меньше второго, то двигаем указатель первого интервала
        if (first[1] < second[1]) {
            left++;
            // Если окончание первого интервала больше второго, то двигаем указатель второго интервала
        } else if (first[1] > second[1]) {
            right++;
        } else {
            left++;
            right++;
        }
    }

    return result; // Список интервалов
};