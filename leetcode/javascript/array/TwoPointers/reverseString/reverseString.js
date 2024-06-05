/*
Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]
 */

/*
Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]
 */

/**
 * @param {character[]} s
 * @return {void} Do not return anything, modify s in-place instead.
 */
var reverseString = function(s) {
    let left = 0, right = s.length - 1;

    while (left < right) {
        let temp = s[left];
        s[left] = s[right];
        s[right] = temp;
        left++;
        right--;
    }
};

reverseString(["h","e","l","l","o"]);