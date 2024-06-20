// Сортировка выбором - O(n^2)
// Parameter:
//  1. random array

// 1. Находит наименьшее значение в массиве
/**
 * @param {number[]} array
 * @return {number}
 */
const findSmallestIndex = (array) => {
    let smallestElement = array[0]; // Сохраняет наименьшее значение
    let smallestIndex = 0; // Сохраняет индекс наименьшего значения

    for (let i = 1; i < array.length; i++) {
        if (array[i] < smallestElement) {
            smallestElement = array[i];
            smallestIndex = i;
        }
    }

    return smallestIndex;
};

// 2. Сортирует массив
/**
 * @param {number[]} array
 * @return {number[]}
 */
const selectionSort = (array) => {
    // Скопируйте значения из массива, поскольку он должен быть неизменяемым.
    // Без этого после вызова исходного массива selectSort станет пустым.
    const sortingArray = [...array];
    const sortedArray = [];

    for (let i = 0; i < sortingArray.length; i++) {
        // Находит наименьший элемент в заданном массиве
        const smallestIndex = findSmallestIndex(sortingArray);
        // Добавляет наименьший элемент в новый массив
        sortedArray.push(sortingArray.splice(smallestIndex, 1)[0]);
    }

    return sortedArray;
}

const array = [5, 3, 6, 2, 10];
console.log(selectionSort(array)); // [2, 3, 5, 6, 10]
console.log(array); // [5, 3, 6, 2, 10]