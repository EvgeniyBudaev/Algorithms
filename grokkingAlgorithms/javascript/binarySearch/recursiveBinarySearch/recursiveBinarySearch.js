/**
 * Searches recursively number from the list
 * @param {Array} list
 * @param {number} item Search item
 * @param {number} low Lower limit of search in the list
 * @param {number} high Highest limit of search in the list
 * @return {(number | null)} Number if the value is found or NULL otherwise
 */
const binarySearch = (list, item, low = 0, high = list.length - 1) => {
    const mid = Math.floor((low + high) / 2);
    const guess = list[mid]; // guess - предполагаемое число

    if (low > high) return null;

    if (guess === item) {
        return mid;
    } else if (guess > item) {
        high = mid - 1;
        return binarySearch(list, item, low, high);
    } else {
        high = mid + 1;
        return binarySearch(list, item, low, high);
    }
}

const myList = [1, 3, 5, 7, 9];

console.log(binarySearch(myList, 3)); // 1
console.log(binarySearch(myList, -1)); // null