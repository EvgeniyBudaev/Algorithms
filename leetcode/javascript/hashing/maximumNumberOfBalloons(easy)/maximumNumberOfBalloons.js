/* https://leetcode.com/problems/maximum-number-of-balloons/description/

Учитывая строковый текст, вы хотите использовать символы текста для формирования как можно большего количества
экземпляров слова «balloon».
Каждый символ в тексте можно использовать не более одного раза. Верните максимальное количество экземпляров,
которые можно сформировать.

Input: text = "nlaebolko"
Output: 1

Input: text = "loonbalxballpoon"
Output: 2

Input: text = "leetcode"
Output: 0
*/

/**
 * @param {string} text
 * @return {number}
 */
var maxNumberOfBalloons = function(text) {
    const map = { b: 0, a: 0, l: 0, o: 0, n: 0, };

    for(const char of text) {
        map[char]++;
    }

    return Math.floor(
        Math.min(map.b, map.a, map.l / 2, map.o / 2, map.n)
    );
};

console.log(maxNumberOfBalloons("nlaebolko")); // 1
console.log(maxNumberOfBalloons("loonbalxballpoon")); // 2
console.log(maxNumberOfBalloons("leetcode")); // 0