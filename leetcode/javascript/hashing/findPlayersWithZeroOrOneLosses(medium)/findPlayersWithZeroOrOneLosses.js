/* https://leetcode.com/problems/find-players-with-zero-or-one-losses/description/

Вам дан целочисленный массив matchs, где match[i] = [winneri,loseri] указывает, что игрок-победитель победил
игрока-проигравшего в матче.
Вернуть ответ в виде списка размером 2, где:

ответ[0] — список всех игроков, которые не проиграли ни одного матча.
ответ[1] — список всех игроков, проигравших ровно один матч.
Значения в двух списках должны возвращаться в порядке возрастания.

Примечание:
Следует учитывать только игроков, сыгравших хотя бы один матч.
Тестовые случаи будут сгенерированы таким образом, чтобы никакие два совпадения не имели одинаковый результат.

Input: matches = [[1,3],[2,3],[3,6],[5,6],[5,7],[4,5],[4,8],[4,9],[10,4],[10,9]]
Output: [[1,2,10],[4,5,7,8]]
Explanation:
Players 1, 2, and 10 have not lost any matches.
Players 4, 5, 7, and 8 each have lost one match.
Players 3, 6, and 9 each have lost two matches.
Thus, answer[0] = [1,2,10] and answer[1] = [4,5,7,8].

Input: matches = [[2,3],[1,3],[5,4],[6,4]]
Output: [[1,2,5,6],[]]
Explanation:
Players 1, 2, 5, and 6 have not lost any matches.
Players 3 and 4 each have lost two matches.
Thus, answer[0] = [1,2,5,6] and answer[1] = [].
 */

/**
 * @param {number[][]} matches
 * @return {number[][]}
 */
var findWinners = function(matches) {
    const obj = {};

    for (let match of matches) {
        if(!obj[match[0]]) obj[match[0]] = 0;
        if(!obj[match[1]]) obj[match[1]] = 0;

        ++obj[match[1]];
    }

    const ans = [[],[]];

    for (let [key,value] of Object.entries(obj)) {
        if (value === 0) ans[0].push(key);
        if (value === 1) ans[1].push(key);
    }

    return ans;
};

const matches = [[1,3],[2,3],[3,6],[5,6],[5,7],[4,5],[4,8],[4,9],[10,4],[10,9]];
console.log(findWinners(matches));
