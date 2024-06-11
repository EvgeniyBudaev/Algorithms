/*
Учитывая целочисленный массив arr, подсчитайте, сколько элементов x существует так, что x + 1 также находится в arr.

Если в arr есть дубликаты, посчитайте их отдельно.
 */

/*
Input: arr = [1,2,3]
Output: 2
Explanation: 1 and 2 are counted cause 2 and 3 are in arr.
 */

/*
Input: arr = [1,1,3,3,5,5,7,7]
Output: 0
Explanation: No numbers are counted, cause there's no 2, 4, 6, or 8 in arr.
 */

/*
Input: arr = [1,3,2,3,5,0]
Output: 3
Explanation: 0, 1 and 2 are counted cause 1, 2 and 3 are in arr.
 */

/*
Input: arr = [1,1,2,2]
Output: 2
Explanation: Two 1s are counted cause 2 is in arr.
 */

/*
Input: arr = [1,1,2]
Output: 2
Explanation: Both 1s are counted because 2 is in the array.
 */

/**
 * @param {number[]} arr
 * @return {number}
 */
var countElements = function (arr) {
    let map = {}, total = 0;

    // Добавляем элементы в map
    for(let i of nums) map[i] ? map[i]++ : map[i] = 1;

    // Удаление повторяющихся элементов
    let newNums = [... new Set(nums)];

    // Если длина массива после удаления повторяющихся чисел меньше трёх, возвращаем 0;
    if(newNums.length < 3) return 0;

    // Сортируем массив newNums и удаляем первый и последний элементы.
    // для всех остальных элементов проверяем их количество на карте и добавляем в переменную total
    newNums.sort((a,b) => a-b).slice(1, newNums.length-1).forEach(num => total += map[num]);

    // Возвращаем итоговую переменную
    return total;
};

console.log(countElements([1,2,3]));
