export function push(array,...elements) {
    // проверяем, является ли переданный объект массивом
    if (!Array.isArray(array)) {
        throw new Error('the first argument to the function must be an array');
    }

    // получаем текущую длину массива
    const length = elements.length;

    // добавляем новые элементы в массив
    for (let i = 0; i < length; i++) {
        array[array.length] = elements[i];
    }

    // возвращаем новый массив
    return array;
}