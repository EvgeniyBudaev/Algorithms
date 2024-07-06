/* https://leetcode.com/problems/check-if-the-sentence-is-pangram/description/
Панграмма – это предложение, в котором каждая буква английского алфавита встречается хотя бы один раз.
Учитывая строковое предложение, содержащее только строчные буквы английского языка, верните true, если предложение
является панграммой, или false в противном случае.

Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
Output: true

Input: sentence = "leetcode"
Output: false
*/

/**
 * @param {string} sentence
 * @return {boolean}
 */
var checkIfPangram = function(sentence) {
    if (sentence.length < 26) return false;
    return new Set([...sentence]).size === 26;
};

console.log(checkIfPangram("thequickbrownfoxjumpsoverthelazydog"));