export function includes(array, searchString) {
    for (let i = 0; i < array.length; i++) {
        if (array[i] === searchString) {
            return true;
        }
    }

    return false;
}