/*
Учитывая массив целых чисел A, верните наибольшее целое число, которое встречается только один раз.
Если ни одно целое число не встречается один раз, верните -1.
 */

/*
Input: [5,7,3,9,4,9,8,3,1]
Output: 8
Explanation:
The maximum integer in the array is 9 but it is repeated. The number 8 occurs only once, so it's the answer.
 */

/*
Input: [9,9,8,8]
Output: -1
Explanation:
There is no number that occurs only once.
 */

var largestUniqueNumber = function (nums) {
    const map = new Map();

    // Заполняем мапу
    for (let i = 0; i < nums.length; i++) {
        if (map.has(nums[i])) {
            map.set(nums[i], map.get(nums[i]) + 1);
        } else {
            map.set(nums[i], 1);
        }
    }

    // Находим наибольшее число, которое встречается только один раз
    let max = -1;

    for (let [key, value] of map) {
        if (value === 1 && key > max) {
            max = key;
        }
    }

    return max;
}

console.log(largestUniqueNumber([5,7,3,9,4,9,8,3,1])); // 8
console.log(largestUniqueNumber([9,9,8,8])); // -1