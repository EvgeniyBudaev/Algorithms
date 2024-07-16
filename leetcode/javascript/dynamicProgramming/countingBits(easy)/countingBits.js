/* https://leetcode.com/problems/counting-bits/description/

Учитывая целое число n, верните массив ans длины n + 1 такой, что для каждого i (0 <= i <= n) ans[i] — это количество
единиц в двоичном представлении i.

Input: n = 2
Output: [0,1,1]
Explanation:
0 --> 0
1 --> 1
2 --> 10
*/

/**
 * @param {number} n
 * @return {number[]}
 */
var countBits = function(n) {
    const dp = new Array(n + 1).fill(0);
    let offset = 1;

    for (let i = 1; i < n + 1; i++) {
        if (offset * 2 === i) offset = i;
        dp[i] = 1 + dp[i - offset];
    }

    return dp;
};

console.log(countBits(2)); // [0,1,1]