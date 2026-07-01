/* 1346. Check If N and Its Double Exist
https://leetcode.com/problems/check-if-n-and-its-double-exist/description/

Учитывая массив целых чисел, проверьте, существуют ли два индекса i и j такие, что:
i != j
0 <= i, j < arr.length
arr[i] == 2 * arr[j]

Input: arr = [10,2,5,3]
Output: true
Explanation: For i = 0 and j = 2, arr[i] == 10 == 2 * 5 == 2 * arr[j]

Input: arr = [3,1,7,11]
Output: false
Пояснение: Не существует i и j, удовлетворяющих этим условиям.
*/

/**
 * @param {number[]} arr
 * @return {boolean}
 */
var checkIfExist = function (arr) {
    let left = 0, right = 1; // инициализируем два указателя

    while (left < arr.length - 1) { // запускаем цикл для перемещения левого указателя
        if (arr[left] === arr[right] * 2 || arr[right] === arr[left] * 2) { // проверяем условие на равенство
            return true; // если условие выполняется, возвращаем true
        } else if (right === arr.length - 1) { // если достигнут конец массива, перемещаем левый указатель вправо и обновляем правый
            left++; // двигаем левый указатель вправо
            right = left + 1;
        } else { // перемещаем правый указатель вправо
            right++;
        }
    }

    return false; // если не найдено ни одной пары, возвращаем false
};