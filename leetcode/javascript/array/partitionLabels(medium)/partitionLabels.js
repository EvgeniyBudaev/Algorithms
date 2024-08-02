/*

Вам дана строка s. Мы хотим разделить строку на как можно большее количество частей, чтобы каждая буква появлялась не
более чем в одной части. Обратите внимание, что разбиение выполнено так, что после объединения всех частей по порядку
результирующая строка должна быть s. Верните список целых чисел, представляющих размер этих частей.

Input: s = "ababcbacadefegdehijhklij"
Output: [9,7,8]
Объяснение:
Раздел — «ababcbaca», «defegde», «hijhklij».
Это раздел, в котором каждая буква встречается не более чем в одной части.
Раздел типа «ababcbacadefegde», «hijhklij» неверен, поскольку он разбивает s на меньшее количество частей.

Input: s = "eccbbbbdec"
Output: [10]
*/

/**
 * @param {string} s
 * @return {number[]}
 */
var partitionLabels = function(s) {
    const lastIdx = {};
    for (let i = 0; i < s.length; i++) {
        lastIdx[s[i]] = i;
    }
    let curLast = 0, res = [], acc = 0;
    for (let i = 0; i < s.length; i++) {
        curLast = Math.max(curLast, lastIdx[s[i]]);
        if (i === curLast) {
            res.push(i + 1 - acc);
            acc = i + 1;
        }
    }
    return res;
};

console.log(partitionLabels("ababcbacadefegdehijhklij")); // [9,7,8]