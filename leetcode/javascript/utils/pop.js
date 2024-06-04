export function pop(array) {
    // проверяем, является ли переданный объект массивом
    if (!Array.isArray(array)) {
        throw new Error('function argument must be an array');
    }

    // проверяем, не пуст ли массив
    if (array.length === 0) {
        return undefined;
    }

    // извлекаем последний элемент массива
    const lastElement = array[array.length - 1];

    // удаляем последний элемент из массива
    array.length -= 1;

    // возвращаем извлеченный элемент
    return lastElement;
}