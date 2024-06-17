/*
Учитывая массив неотрицательных целых чисел arr, вы изначально находитесь в начальном индексе массива.
Когда вы находитесь на индексе i, вы можете перейти к i + arr[i] или i - arr[i], проверьте, можете ли вы достичь любого
индекса со значением 0.
Обратите внимание, что вы не можете выйти за пределы массива в любой момент.
*/

/*
Input: arr = [4,2,3,0,3,1,2], start = 5
Output: true
Explanation:
All possible ways to reach at index 3 with value 0 are:
index 5 -> index 4 -> index 1 -> index 3
index 5 -> index 6 -> index 4 -> index 1 -> index 3
 */

/**
 * @param {number[]} arr
 * @param {number} start
 * @return {boolean}
 */
var canReach = function(arr, start) {
    // BFS
    const queue = [start];
    for (let len = 0, max = arr.length; len < queue.length; len++) {
        const idx = queue[len];
        if (arr[idx] === -1) continue;
        if (arr[idx] === 0) return true;
        idx + arr[idx] < max && queue.push(idx + arr[idx]);
        idx - arr[idx] >= 0 && queue.push(idx - arr[idx]);
        arr[idx] = -1;
    }
    return false;
};

console.log(canReach([4,2,3,0,3,1,2], 5)); // true