/**
 * Consider the factorial of the number
 * @param {number} x Number
 * @returns {number} Result
 */
const factorialRecursive = x => {
    if (x === 1) return 1;
    return  x  *  factorialRecursive(x - 1);
};

const factorialIterative = x => {
    let result = 1;
    for (let i = 1; i <= x; i++) {
        result *= i;
    }
    return result;
}

console.log(factorialRecursive(5));
console.log(factorialIterative(5));