// shift - функция удаления первого элемента из массива

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
    for (let i = 1; i < array.length; i++) {
        // сдвигаем каждый элемент на одну позицию влево
        array[i - 1] = array[i];
    }

    // уменьшаем длину массива на 1
    length--;

    // Возвращаем удаленный первый элемент
    return firstElement;
}