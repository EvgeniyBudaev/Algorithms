export function sort(array) {
    // проверяем, является ли переданный объект массивом
    if (!Array.isArray(array)) {
        throw new Error('function argument must be an array');
    }

    // базовый случай, массивы с 0 или 1 элементом уже «отсортированы»
    if (array.length < 2) return array;

    // рекурсивный случай
    let pivot = array[0];

    // подмассив всех элементов меньше опорного
    let less = array.slice(1).filter(function(el) {
        return el <= pivot;
    });

    // подмассив всех элементов, больших, чем стержень
    let greater = array.slice(1).filter(function(el) {
        return el > pivot;
    });

    // возвращаем отсортированный массив
    return sort(less).concat([pivot], sort(greater));
}