export function slice(array, start = 0, end = array.length) {
    // проверяем, является ли переданный объект массивом
    if (!Array.isArray(array)) {
        throw new Error('the first argument to the function must be an array');
    }

    // Создаем новый массив для хранения результатов
    let result = [];

    // Определяем границы извлечения
    let startIndex = Math.max(start, 0);
    let endIndex = Math.min(end, array.length);

    // Извлекаем элементы из исходного массива
    for (let i = startIndex; i < endIndex; i++) {
        result.push(array[i]);
    }

    // Возвращаем новый массив
    return result;
}