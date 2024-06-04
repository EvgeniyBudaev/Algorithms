import {push} from "./push.js";

export function splice(array, start, deleteCount,...items) {
    // Проверяем, является ли переданный объект массивом
    if (!Array.isArray(array)) {
        throw new Error('the first argument to the function must be an array');
    }

    // Проверяем, не превышает ли начальный индекс длину массива
    if (start < 0 || start > array.length) {
        throw new Error('the starting index is outside the array\'s bounds.');
    }

    // Удаляем элементы
    const removedItems = [];
    if (deleteCount!== undefined && deleteCount > 0) {
        for (let i = 0; i < deleteCount; i++) {
            push(removedItems, array[start + i]);
        }
        for (let i = start + deleteCount; i < array.length; i++) {
            array[i - deleteCount] = array[i];
        }
        array.length -= deleteCount;
    }

    // Добавляем новые элементы
    items.forEach(item => {
        array.push(item);
    });

    // Возвращаем удаленные элементы
    return removedItems;
}