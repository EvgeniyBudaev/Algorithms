/* https://leetcode.com/problems/decode-ways/description/

Вы перехватили секретное сообщение, закодированное в виде строки цифр. Сообщение декодируется с помощью следующего
сопоставления:

«1» -> «А»

«2» -> «Б»

...

«25» -> «Й»

«26» -> «З»

Однако при декодировании сообщения вы понимаете, что существует много разных способов декодирования сообщения, поскольку
некоторые коды содержатся в других кодах («2» и «5» против «25»).

Например, «11106» можно декодировать как:

«AAJF» с группировкой (1, 1, 10, 6)
«KJF» с группировкой (11, 10, 6)
Группировка (1, 11, 06) недействительна, поскольку «06» не является допустимым кодом (действителен только «6»).
Примечание: могут существовать строки, которые невозможно декодировать.

Дана строка s, содержащая только цифры, верните количество способов ее декодирования. Если всю строку невозможно
декодировать каким-либо допустимым способом, верните 0.

Тестовые случаи генерируются таким образом, чтобы ответ умещался в 32-битное целое число.

Input: s = "12"
Output: 2
Explanation:
«12» можно расшифровать как «AB» (1 2) или «L» (12).

Input: s = "226"
Output: 3
Explanation:
"226" could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).

Input: s = "06"
Output: 0
Explanation:
"06" cannot be mapped to "F" because of the leading zero ("6" is different from "06"). In this case, the string is not a
valid encoding, so return 0.
*/

/**
 * @param {string} s
 * @return {number}
 */
var numDecodings = function(s) {
    if (!s || s[0] === '0') return 0;
    const n = s.length;
    const dp = new Array(n + 1).fill(0);
    dp[0] = 1;
    dp[1] = 1;

    for (let i = 2; i <= n; ++i) {
        const oneDigit = parseInt(s[i - 1]);
        const twoDigits = parseInt(s.substring(i - 2, i));
        if (oneDigit !== 0) dp[i] += dp[i - 1];
        if (10 <= twoDigits && twoDigits <= 26) dp[i] += dp[i - 2];
    }

    return dp[n];
};

console.log(numDecodings("12")); // 2