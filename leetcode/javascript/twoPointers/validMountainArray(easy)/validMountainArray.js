/* 941. Valid Mountain Array
https://leetcode.com/problems/valid-mountain-array/

Учитывая массив целых чисел arr, верните true если и только если он является допустимым массивом гор.
Напомним, что arr является горным массивом тогда и только тогда, когда:

arr.length >= 3
Существует некоторое i с 0 < i < arr.length - 1 таким , что:
arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

Input: arr = [2,1]
Output: false

Input: arr = [3,5,5]
Output: false

Input: arr = [0,3,2,1]
Output: true
*/

/**
 * @param {number[]} arr
 * @return {boolean}
 */
var validMountainArray = function (arr) {
    // По условию горный массив должен содержать как минимум 3 элемента
    if (arr.length < 3) {
        return false;
    }

    let left = 0;                  // Левый указатель (старт с начала)
    let right = arr.length - 1;    // Правый указатель (старт с конца)

    // Двигаем левый указатель вправо, пока элементы строго возрастают
    while (left + 1 < arr.length && arr[left] < arr[left + 1]) {
        left++;
    }

    // Двигаем правый указатель влево, пока элементы строго убывают
    while (right - 1 >= 0 && arr[right] < arr[right - 1]) {
        right--;
    }

    // Массив является горным, если:
    // 1. Указатели встретились на одной вершине (left === right).
    // 2. Вершина не находится на самом краю массива (left > 0 и right < arr.length - 1),
    //    что гарантирует наличие И подъема, И спуска.
    return left === right && left > 0 && right < arr.length - 1;
};