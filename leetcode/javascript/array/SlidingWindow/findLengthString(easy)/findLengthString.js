/*
Вам дана двоичная строка s (строка, содержащая только «0» и «1»). Вы можете выбрать до одного «0» и изменить его на «1».
Какова длина самой длинной подстроки, содержащей только «1»?

Например, если задано s = «1101100111», ответ будет равен 5. Если вы выполните переворот по индексу 2,
строка станет 1111100111.
 */

/*
Input: nums = "1101100111"
Output: 5
 */

/**
 * @param {string} s
 * @return {number}
 */
var findLengthString = function(s) {
    // curr — текущее количество нулей в окне
    let left = 0, curr = 0, ans = 0;
    for (let right = 0; right < s.length; right++) {
        if (s[right] === "0") {
            curr++;
        }

        while (curr > 1) {
            if (s[left] === "0") {
                curr -= 1;
            }
            left++;
        }

        ans = Math.max(ans, right - left + 1);
    }

    return ans;
}

console.log(findLengthString("1101100111")); // 5