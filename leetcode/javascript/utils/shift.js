export function shift(array) {
    // проверяем, является ли переданный объект массивом
    if (!Array.isArray(array)) {
        throw new Error('function argument must be an array');
    }

    // Проверяем, не пуст ли массив
    if (array.length === 0) {
        return undefined;
    }

    // Удаляем первый элемент из массива
    const firstElement = array[0];
    for (let i = 0; i < array.length - 1; i++) {
        array[i] = array[i + 1];
    }
    --array.length;

    // Возвращаем удаленный первый элемент
    return firstElement;
}