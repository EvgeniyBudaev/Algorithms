import {push} from "./push";

export function concat(array1, array2) {
    // проверяем, является ли переданный объект массивом
    if (!Array.isArray(array1) || !Array.isArray(array2)) {
        throw new Error('function arguments must be an arrays');
    }
    // создаем пустой массив для хранения результата
    let mergedArray = [];

    // добавляем элементы первого массива во временный массив
    for (let i = 0; i < array1.length; i++) {
        push(mergedArray, array1[i]);
    }

    // добавляем элементы второго массива во временный массив
    for (let j = 0; j < array2.length; j++) {
        push(mergedArray, array2[j]);
    }

    // возвращаем результат объединенного массива
    return mergedArray;
}