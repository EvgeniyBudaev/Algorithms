export function filter(array, callback) {
    // проверяем, является ли переданный объект массивом
    if (!Array.isArray(array)) {
        throw new Error('the first argument to the function must be an array');
    }

    // Создаем новый массив для хранения результатов
    let result = [];

    // Применяем функцию-критерий ко всем элементам массива
    for (let i = 0; i < array.length; i++) {
        if (callback(array[i], i, array)) {
            result.push(array[i]);
        }
    }

    // Возвращаем новый массив
    return result;
}