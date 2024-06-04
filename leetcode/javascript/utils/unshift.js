export function unshift(array,...elements) {
    // Проверяем, является ли переданный объект массивом
    if (!Array.isArray(array)) {
        throw new Error('the first argument to the function must be an array');
    }

    // Добавляем каждый элемент в начало массива
    let i= array.length + elements.length -1;
    for( i ; i >= elements.length; i--) {
        array[i] = array[i - elements.length ];
    }

    for(i; i >= 0; i--) {
        array[i] = elements[i];
    }

    // возвращаем новую длину массива
    return array.length;
}