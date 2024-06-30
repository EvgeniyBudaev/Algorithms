// https://www.youtube.com/watch?v=btS8Qf-wM2M&ab_channel=%D0%95%D0%BB%D0%B5%D0%BD%D0%B0%D0%9B%D0%B8%D1%82%D0%B2%D0%B8%D0%BD%D0%BE%D0%B2%D0%B0%E2%80%94%D0%98%D1%81%D0%BA%D1%83%D1%81%D1%81%D1%82%D0%B2%D0%BE%D0%92%D0%B5%D0%B1-%D1%80%D0%B0%D0%B7%D1%80%D0%B0%D0%B1%D0%BE%D1%82%D0%BA%D0%B8%F0%9F%9B%B8
// O(n log n)

/**
 * Quick array sorting
 * @param {Array} array Source array
 * @returns {Array} Sorted array
 */
const quickSort = array => {
    if (array.length < 2) return array;
    let pivotIndex = Math.floor(array.length / 2);
    let pivot = array[pivotIndex];
    let keysAreLessPivot = [];
    let keysAreMorePivot= [];

    for (let i = 0; i < array.length; i++) {
        if (i === pivotIndex) continue;
        if (array[i] <= pivot) {
            keysAreLessPivot.push(array[i]);
        } else {
            keysAreMorePivot.push(array[i]);
        }
    }

    return [
        ...quickSort(keysAreLessPivot),
        pivot,
        ...quickSort(keysAreMorePivot)
    ];
};

console.log(quickSort([10, 5, 2, 3])); // [2, 3, 5, 10]