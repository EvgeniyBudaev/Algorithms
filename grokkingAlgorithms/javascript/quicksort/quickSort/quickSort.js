// O(n log n)
/**
 * Quick array sorting
 * @param {Array} array Source array
 * @returns {Array} Sorted array
 */
const quickSort = array => {
    if (array.length < 2) return array; // Базовый случай
    const pivot = array[0]; // Опорный элемент
    const keysAreLessPivot = array.slice(1).filter(key => key <= pivot);
    const keysAreMorePivot = array.slice(1).filter(key => key > pivot);
    return [
        ...quickSort(keysAreLessPivot),
        pivot,
        ...quickSort(keysAreMorePivot)
    ];
};

console.log(quickSort([10, 5, 2, 3])); // [2, 3, 5, 10]